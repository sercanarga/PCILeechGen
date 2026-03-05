package tclgen

import (
	"bytes"
	"fmt"
	"path/filepath"
	"text/template"

	"github.com/sercanarga/pcileechgen/internal/board"
	"github.com/sercanarga/pcileechgen/internal/donor"
	"github.com/sercanarga/pcileechgen/internal/firmware"
	"github.com/sercanarga/pcileechgen/internal/pci"
)

// projectTCLData holds template data for Vivado project generation.
type projectTCLData struct {
	BoardName string
	FPGAPart  string
	SrcPath   string
	IPPath    string
	TopModule string

	DeviceID       string
	VendorID       string
	RevisionID     string
	SubsysVendorID string
	SubsysDeviceID string
	ClassCodeBase  string
	ClassCodeSub   string
	ClassCodeIntf  string

	LinkSpeed     string
	LinkWidth     string
	TrgtLinkSpeed string

	Bar0Enabled bool
	Bar0Size    string
	Bar0Scale   string
	Bar064bit   bool
}

// buildTCLData holds template data for Vivado build script.
type buildTCLData struct {
	BoardName string
	FPGAPart  string
	Jobs      int
	Timeout   int
}

var projectTCLTmpl = template.Must(template.New("project").Parse(`#
# {{.BoardName}} / {{.FPGAPart}}
# Device: {{.VendorID}}:{{.DeviceID}} rev {{.RevisionID}}
#

set origin_dir "."
set _xil_proj_name_ "{{.BoardName}}"

create_project ${_xil_proj_name_} ./${_xil_proj_name_} -part {{.FPGAPart}}
set proj_dir [get_property directory [current_project]]

# Project properties
set obj [current_project]
set_property -name "default_lib" -value "xil_defaultlib" -objects $obj
set_property -name "enable_vhdl_2008" -value "1" -objects $obj
set_property -name "part" -value "{{.FPGAPart}}" -objects $obj
set_property -name "simulator_language" -value "Mixed" -objects $obj
set_property -name "xpm_libraries" -value "XPM_CDC XPM_MEMORY" -objects $obj

# Source files
if {[string equal [get_filesets -quiet sources_1] ""]} {
  create_fileset -srcset sources_1
}

set obj [get_filesets sources_1]
set sv_files [glob -nocomplain "{{.SrcPath}}/*.sv"]
set svh_files [glob -nocomplain "{{.SrcPath}}/*.svh"]
set all_src_files [concat $sv_files $svh_files]
if {[llength $all_src_files] > 0} {
  set imported_files [import_files -fileset sources_1 $all_src_files]
}

# Set file types
foreach f [get_files -of_objects [get_filesets sources_1] -filter {NAME =~ "*.sv"}] {
  set_property -name "file_type" -value "SystemVerilog" -objects $f
}
foreach f [get_files -of_objects [get_filesets sources_1] -filter {NAME =~ "*.svh"}] {
  set_property -name "file_type" -value "Verilog Header" -objects $f
}

# Generated COE files
set coe_files [list \
  [file normalize "${origin_dir}/pcileech_cfgspace.coe"] \
  [file normalize "${origin_dir}/pcileech_cfgspace_writemask.coe"] \
  [file normalize "${origin_dir}/pcileech_bar_zero4k.coe"] \
]
import_files -fileset sources_1 $coe_files

# Import IP cores from board library
set ip_files [glob -nocomplain "{{.IPPath}}/*.xci"]
if {[llength $ip_files] > 0} {
  set imported_ip [import_files -fileset sources_1 $ip_files]
  foreach ip $imported_ip {
    set ip_obj [get_files -of_objects [get_filesets sources_1] [list "*[file tail $ip]"]]
    if { $ip_obj != "" } {
      set_property -name "generate_files_for_reference" -value "0" -objects $ip_obj
      set_property -name "registered_with_manager" -value "1" -objects $ip_obj
      if { ![get_property "is_locked" $ip_obj] } {
        set_property -name "synth_checkpoint_mode" -value "Singular" -objects $ip_obj
      }
    }
  }
}

set ip_coe_files [glob -nocomplain "{{.IPPath}}/*.coe"]
if {[llength $ip_coe_files] > 0} {
  import_files -fileset sources_1 $ip_coe_files
}

# Upgrade IPs to current Vivado version (must happen before configuring properties)
set all_ips [get_ips -quiet *]
if {[llength $all_ips] > 0} {
  upgrade_ip $all_ips
}

# Patch PCIe IP core with donor identity
set pcie_ip [get_ips -quiet pcie_7x_0]
if { $pcie_ip != "" } {
  puts "Configuring PCIe IP core with donor device identity..."

  # IDs
  set_property -dict [list \
    CONFIG.Device_ID            {{.DeviceID}} \
    CONFIG.Vendor_Id            {{.VendorID}} \
    CONFIG.Revision_ID          {{.RevisionID}} \
    CONFIG.Subsystem_Vendor_ID  {{.SubsysVendorID}} \
    CONFIG.Subsystem_ID         {{.SubsysDeviceID}} \
    CONFIG.Class_Code_Base      {{.ClassCodeBase}} \
    CONFIG.Class_Code_Sub       {{.ClassCodeSub}} \
    CONFIG.Class_Code_Interface {{.ClassCodeIntf}} \
  ] $pcie_ip

  # Link config (clamped to board lanes)
  set_property -dict [list \
    CONFIG.Maximum_Link_Width   {{.LinkWidth}} \
    CONFIG.Link_Speed           {{.LinkSpeed}} \
    CONFIG.Trgt_Link_Speed      {{.TrgtLinkSpeed}} \
  ] $pcie_ip
{{if .Bar0Enabled}}
  # BAR0
  set_property -dict [list \
    CONFIG.Bar0_Enabled         true \
    CONFIG.Bar0_Type            Memory \
    CONFIG.Bar0_Scale           {{.Bar0Scale}} \
    CONFIG.Bar0_Size            {{.Bar0Size}} \
    CONFIG.Bar0_64bit           {{if .Bar064bit}}true{{else}}false{{end}} \
  ] $pcie_ip
{{end}}
  puts "PCIe IP core configured: {{.VendorID}}:{{.DeviceID}} Link={{.LinkWidth}} {{.LinkSpeed}}"
} else {
  puts "WARNING: PCIe IP core pcie_7x_0 not found, skipping donor identity configuration"
}

# Top module
set_property -name "top" -value "{{.TopModule}}" -objects [get_filesets sources_1]
set_property -name "top_auto_set" -value "0" -objects [get_filesets sources_1]

# Constraints
if {[string equal [get_filesets -quiet constrs_1] ""]} {
  create_fileset -constrset constrs_1
}
set xdc_files [glob -nocomplain "{{.SrcPath}}/*.xdc"]
if {[llength $xdc_files] > 0} {
  set imported_xdc [import_files -fileset constrs_1 $xdc_files]
  foreach f $imported_xdc {
    set file_obj [get_files -of_objects [get_filesets constrs_1] [list "*[file tail $f]"]]
    if { $file_obj != "" } {
      set_property -name "file_type" -value "XDC" -objects $file_obj
    }
  }
}
set_property -name "target_part" -value "{{.FPGAPart}}" -objects [get_filesets constrs_1]

# Simulation fileset
if {[string equal [get_filesets -quiet sim_1] ""]} {
  create_fileset -simset sim_1
}
set_property -name "top" -value "{{.TopModule}}" -objects [get_filesets sim_1]

# Synthesis run
if {[string equal [get_runs -quiet synth_1] ""]} {
  create_run -name synth_1 -part {{.FPGAPart}} -flow {Vivado Synthesis 2022} -strategy "Vivado Synthesis Defaults" -report_strategy {No Reports} -constrset constrs_1
}
current_run -synthesis [get_runs synth_1]

# Implementation run
if {[string equal [get_runs -quiet impl_1] ""]} {
  create_run -name impl_1 -part {{.FPGAPart}} -flow {Vivado Implementation 2022} -strategy "Vivado Implementation Defaults" -report_strategy {No Reports} -constrset constrs_1 -parent_run synth_1
}
current_run -implementation [get_runs impl_1]
set_property STEPS.PLACE_DESIGN.ARGS.DIRECTIVE ExtraPostPlacementOpt [get_runs impl_1]

puts "Project ${_xil_proj_name_} created successfully."
`))

var buildTCLTmpl = template.Must(template.New("build").Parse(`#
# PCILeechGen - Vivado Build Script
# Board: {{.BoardName}}
#

open_project {{.BoardName}}/{{.BoardName}}.xpr

# Run synthesis
puts "Starting synthesis..."
launch_runs synth_1 -jobs {{.Jobs}}
wait_on_run synth_1 -timeout {{.Timeout}}

if {[get_property STATUS [get_runs synth_1]] != "synth_design Complete!"} {
  puts "ERROR: Synthesis failed!"
  exit 1
}
puts "Synthesis completed successfully."

# Run implementation
puts "Starting implementation..."
launch_runs impl_1 -to_step write_bitstream -jobs {{.Jobs}}
wait_on_run impl_1 -timeout {{.Timeout}}

if {[get_property STATUS [get_runs impl_1]] != "write_bitstream Complete!"} {
  puts "ERROR: Implementation failed!"
  exit 1
}
puts "Implementation completed successfully."

# Generate .bin from .bit
set bit_file [glob {{.BoardName}}/{{.BoardName}}.runs/impl_1/*.bit]
set bin_file [file rootname $bit_file].bin
write_cfgmem -format bin -interface SPIx4 -size 16 -loadbit "up 0x0 $bit_file" -file $bin_file -force

puts "Build complete! Output: $bin_file"
exit 0
`))

// GenerateProjectTCL generates the Vivado project creation TCL script.
func GenerateProjectTCL(ctx *donor.DeviceContext, b *board.Board, libDir string) string {
	ids := firmware.ExtractDeviceIDs(ctx.ConfigSpace, ctx.ExtCapabilities)

	// Clamp link width to board physical lanes
	linkWidth := clampLinkWidth(ids.LinkWidth, b.PCIeLanes)
	linkSpeed := ids.LinkSpeed
	// Clamp link speed to board's physical capability
	if linkSpeed == 0 || linkSpeed > b.MaxLinkSpeedOrDefault() {
		linkSpeed = b.MaxLinkSpeedOrDefault()
	}

	// BAR0 configuration
	bar0Enabled := false
	bar0Scale := "Kilobytes"
	bar0Size := "4"
	bar064bit := false
	if len(ctx.BARs) > 0 {
		bar0 := ctx.BARs[0]
		if bar0.Size > 0 {
			bar0Enabled = true
			// clamp to FPGA BRAM
			bar0SizeClamped := bar0.Size
			bramSize := uint64(b.BRAMSizeOrDefault())
			if bar0SizeClamped > bramSize {
				bar0SizeClamped = bramSize
			}
			bar0Scale, bar0Size = barSizeToTCL(bar0SizeClamped)
			bar064bit = bar0.Type == pci.BARTypeMem64
		}
	}

	// Resolve to absolute paths so TCL works from any working directory
	srcAbs, _ := filepath.Abs(b.SrcPath(libDir))
	ipAbs, _ := filepath.Abs(b.IPPath(libDir))

	data := projectTCLData{
		BoardName:      b.Name,
		FPGAPart:       b.FPGAPart,
		SrcPath:        srcAbs,
		IPPath:         ipAbs,
		TopModule:      b.TopModule,
		DeviceID:       fmt.Sprintf("%04X", ctx.Device.DeviceID),
		VendorID:       fmt.Sprintf("%04X", ctx.Device.VendorID),
		RevisionID:     fmt.Sprintf("%02X", ctx.Device.RevisionID),
		SubsysVendorID: fmt.Sprintf("%04X", ctx.Device.SubsysVendorID),
		SubsysDeviceID: fmt.Sprintf("%04X", ctx.Device.SubsysDeviceID),
		ClassCodeBase:  fmt.Sprintf("%02X", (ctx.Device.ClassCode>>16)&0xFF),
		ClassCodeSub:   fmt.Sprintf("%02X", (ctx.Device.ClassCode>>8)&0xFF),
		ClassCodeIntf:  fmt.Sprintf("%02X", ctx.Device.ClassCode&0xFF),
		LinkSpeed:      linkSpeedToTCL(linkSpeed),
		LinkWidth:      linkWidthToTCL(linkWidth),
		TrgtLinkSpeed:  linkSpeedToTrgt(linkSpeed),
		Bar0Enabled:    bar0Enabled,
		Bar0Size:       bar0Size,
		Bar0Scale:      bar0Scale,
		Bar064bit:      bar064bit,
	}

	var buf bytes.Buffer
	if err := projectTCLTmpl.Execute(&buf, data); err != nil {
		panic(fmt.Sprintf("project TCL template error: %v", err))
	}
	return buf.String()
}

// GenerateBuildTCL generates the Vivado build/synthesis TCL script.
func GenerateBuildTCL(b *board.Board, jobs int, timeout int) string {
	if jobs <= 0 {
		jobs = 4
	}
	if timeout <= 0 {
		timeout = 3600
	}

	data := buildTCLData{
		BoardName: b.Name,
		FPGAPart:  b.FPGAPart,
		Jobs:      jobs,
		Timeout:   timeout,
	}

	var buf bytes.Buffer
	if err := buildTCLTmpl.Execute(&buf, data); err != nil {
		panic(fmt.Sprintf("build TCL template error: %v", err))
	}
	return buf.String()
}
