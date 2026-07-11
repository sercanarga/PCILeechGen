#
# ac701_ft601 / xc7a200tfbg676-2
# Device: 144D:A809 rev 00
#

set origin_dir "."
set _xil_proj_name_ "ac701_ft601"

create_project ${_xil_proj_name_} ./${_xil_proj_name_} -part xc7a200tfbg676-2
set proj_dir [get_property directory [current_project]]

# Project properties
set obj [current_project]
set_property -name "default_lib" -value "xil_defaultlib" -objects $obj
set_property -name "enable_vhdl_2008" -value "1" -objects $obj
set_property -name "part" -value "xc7a200tfbg676-2" -objects $obj
set_property -name "simulator_language" -value "Mixed" -objects $obj
set_property -name "xpm_libraries" -value "XPM_CDC XPM_MEMORY" -objects $obj

# Source files (board sources prepared in local src/ to ensure generated
# controller / donor-specific modules take precedence over stock versions)
if {[string equal [get_filesets -quiet sources_1] ""]} {
  create_fileset -srcset sources_1
}

set obj [get_filesets sources_1]
set sv_files [glob -nocomplain "${origin_dir}/src/*.sv"]
set svh_files [glob -nocomplain "${origin_dir}/src/*.svh"]
set v_files ""
set all_src_files [concat $sv_files $svh_files $v_files]
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
foreach f [get_files -of_objects [get_filesets sources_1] -filter {NAME =~ "*.v"}] {
  set_property -name "file_type" -value "SystemVerilog" -objects $f
}


# Remove any stock bar controller that might have been imported from board sources.
# We want the generated donor-specific version (with NVMe/MSIX logic etc.).
set stock_bar_ctrl [get_files -of_objects [get_filesets sources_1] "*pcileech_tlps128_bar_controller.sv"]
if { $stock_bar_ctrl != "" } {
  remove_files $stock_bar_ctrl
  puts "Removed stock bar controller (will use generated donor-specific version)"
}


# Generated SV modules
set gen_sv_files [glob -nocomplain "${origin_dir}/*.sv"]
if {[llength $gen_sv_files] > 0} {
  set imported_gen [import_files -fileset sources_1 $gen_sv_files]
  foreach f $imported_gen {
    set file_obj [get_files -of_objects [get_filesets sources_1] [list "*[file tail $f]"]]
    if { $file_obj != "" } {
      set_property -name "file_type" -value "SystemVerilog" -objects $file_obj
    }
  }
  puts "Imported [llength $gen_sv_files] generated SV module(s)"
}

# HEX init files
set hex_files [glob -nocomplain "${origin_dir}/*.hex"]
if {[llength $hex_files] > 0} {
  import_files -fileset sources_1 $hex_files
  foreach f [get_files -of_objects [get_filesets sources_1] -filter {NAME =~ "*.hex"}] {
    set_property -name "file_type" -value "Memory Initialization Files" -objects $f
  }
  puts "Imported [llength $hex_files] HEX init file(s)"
}

# Import IP cores from board library
set ip_files [glob -nocomplain "/Users/can/PCILeechGen/lib/pcileech-fpga/ac701_ft601/ip/*.xci"]
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


# Upgrade IPs to current Vivado version (must happen before configuring properties)
set all_ips [get_ips -quiet *]
if {[llength $all_ips] > 0} {
  upgrade_ip $all_ips
}

# Point BRAM IPs to the generated COE files (library defaults are all-zero)
set bar_bram [get_ips -quiet bram_bar_zero4k]
# bram_bar_zero4k is a fixed 4KB (1024-dword) IP; skip donor COE when BAR0 > 4KB (content via SV) or stock-bar.
if { $bar_bram != "" && 0 } {
  set coe_path [file normalize "${origin_dir}/pcileech_bar_zero4k.coe"]
  if {[file exists $coe_path]} {
    set_property -dict [list CONFIG.Coe_File $coe_path] $bar_bram
    puts "BAR BRAM COE updated: $coe_path"
  }
}

set cfg_bram [get_ips -quiet bram_pcie_cfgspace]
if { $cfg_bram != "" } {
  set coe_path [file normalize "${origin_dir}/pcileech_cfgspace.coe"]
  if {[file exists $coe_path]} {
    set_property -dict [list CONFIG.Coe_File $coe_path] $cfg_bram
    puts "Config space BRAM COE updated: $coe_path"
  }
}

set writemask_drom [get_ips -quiet drom_pcie_cfgspace_writemask]
if { $writemask_drom != "" } {
  set coe_path [file normalize "${origin_dir}/pcileech_cfgspace_writemask.coe"]
  if {[file exists $coe_path]} {
    set_property -dict [list CONFIG.coefficient_file $coe_path] $writemask_drom
    puts "Writemask DROM COE updated: $coe_path"
  }
}

# Patch PCIe IP core with donor identity
set pcie_ip [get_ips -quiet pcie_7x_0]
if { $pcie_ip != "" } {
  puts "Configuring PCIe IP core with donor device identity..."

  # IDs
  set_property -dict [list \
    CONFIG.Device_ID            A809 \
    CONFIG.Vendor_Id            144D \
    CONFIG.Revision_ID          00 \
    CONFIG.Subsystem_Vendor_ID  0000 \
    CONFIG.Subsystem_ID         0000 \
    CONFIG.Class_Code_Base      01 \
    CONFIG.Class_Code_Sub       08 \
    CONFIG.Class_Code_Interface 02 \
  ] $pcie_ip

  # Link config (clamped to board lanes)
  set_property -dict [list \
    CONFIG.Maximum_Link_Width   X4 \
    CONFIG.Link_Speed           5.0_GT/s \
    CONFIG.Trgt_Link_Speed      4'h2 \
  ] $pcie_ip

  set_property CONFIG.Bar0_Enabled true $pcie_ip
  set_property -dict [list \
    CONFIG.Bar0_Type            Memory \
    CONFIG.Bar0_Scale           Kilobytes \
    CONFIG.Bar0_Size            16 \
    CONFIG.Bar0_64bit           false \
    CONFIG.Bar0_Prefetchable    false \
  ] $pcie_ip

  set_property CONFIG.Bar1_Enabled false $pcie_ip

  set_property CONFIG.Bar2_Enabled false $pcie_ip

  set_property CONFIG.Bar3_Enabled false $pcie_ip

  set_property CONFIG.Bar4_Enabled false $pcie_ip

  set_property CONFIG.Bar5_Enabled false $pcie_ip

  # Enable or disable DSN based on donor device
  set_property -dict [list \
    CONFIG.DSN_Enabled false \
  ] $pcie_ip

  # MSI vector capability from donor device
  set_property -dict [list \
    CONFIG.Multiple_Message_Capable 1_vector \
  ] $pcie_ip

  puts "PCIe IP configured: 144D:A809 Link=X4 5.0_GT/s DSN=disabled MSI=1_vector"
} else {
  puts "WARNING: PCIe IP core pcie_7x_0 not found, skipping donor identity configuration"
}

# Top module
set_property -name "top" -value "pcileech_ac701_ft601_top" -objects [get_filesets sources_1]
set_property -name "top_auto_set" -value "0" -objects [get_filesets sources_1]

# Constraints (from locally prepared board sources)
if {[string equal [get_filesets -quiet constrs_1] ""]} {
  create_fileset -constrset constrs_1
}
set xdc_files [glob -nocomplain "${origin_dir}/src/*.xdc"]
if {[llength $xdc_files] > 0} {
  set imported_xdc [import_files -fileset constrs_1 $xdc_files]
  foreach f $imported_xdc {
    set file_obj [get_files -of_objects [get_filesets constrs_1] [list "*[file tail $f]"]]
    if { $file_obj != "" } {
      set_property -name "file_type" -value "XDC" -objects $file_obj
    }
  }
}
set_property -name "target_part" -value "xc7a200tfbg676-2" -objects [get_filesets constrs_1]

# Simulation fileset
if {[string equal [get_filesets -quiet sim_1] ""]} {
  create_fileset -simset sim_1
}
set_property -name "top" -value "pcileech_ac701_ft601_top" -objects [get_filesets sim_1]

# Synthesis run
if {[string equal [get_runs -quiet synth_1] ""]} {
  create_run -name synth_1 -part xc7a200tfbg676-2 -flow {Vivado Synthesis 2022} -strategy "Vivado Synthesis Defaults" -report_strategy {No Reports} -constrset constrs_1
}
current_run -synthesis [get_runs synth_1]

# Implementation run
if {[string equal [get_runs -quiet impl_1] ""]} {
  create_run -name impl_1 -part xc7a200tfbg676-2 -flow {Vivado Implementation 2022} -strategy "Vivado Implementation Defaults" -report_strategy {No Reports} -constrset constrs_1 -parent_run synth_1
}
current_run -implementation [get_runs impl_1]
set_property STEPS.PLACE_DESIGN.ARGS.DIRECTIVE ExtraPostPlacementOpt [get_runs impl_1]

puts "Project ${_xil_proj_name_} created successfully."
