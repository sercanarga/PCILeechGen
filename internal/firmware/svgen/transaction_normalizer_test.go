package svgen

import (
	"strings"
	"testing"
)

func compactHDL(source string) string {
	return strings.Join(strings.Fields(source), "")
}

func requireHDLFragments(t *testing.T, source string, fragments map[string]string) {
	t.Helper()
	compact := compactHDL(source)
	for name, fragment := range fragments {
		t.Run(name, func(t *testing.T) {
			if !strings.Contains(compact, compactHDL(fragment)) {
				t.Errorf("generated HDL does not contain %s mechanism\nmissing: %s", name, fragment)
			}
		})
	}
}

func TestGenerateTransactionNormalizerInterfaceAndDecisions(t *testing.T) {
	sv, err := GenerateTransactionNormalizerSV(testConfig())
	if err != nil {
		t.Fatalf("GenerateTransactionNormalizerSV() error = %v", err)
	}

	requireHDLFragments(t, sv, map[string]string{
		"module":                "module pcileech_tlp_normalizer",
		"RCB parameter":         "parameter integer READ_COMPLETION_BOUNDARY = 64",
		"MPS parameter":         "parameter integer MAX_PAYLOAD_BYTES = 128",
		"supported decision":    "output reg request_supported",
		"unsupported decision":  "output reg unsupported_request",
		"UR decision":           "output reg ur_required",
		"non-posted decision":   "output reg non_posted_request",
		"malformed decision":    "output reg malformed_request",
		"posted write":          "output reg posted_write",
		"memory read":           "output reg memory_read",
		"memory write":          "output reg memory_write",
		"3DW and 4DW flag":      "output reg header_4dw",
		"normalized address":    "output reg [63:0] address",
		"normalized length":     "output reg [10:0] length_dw",
		"first BE":              "output reg [3:0] first_be",
		"last BE":               "output reg [3:0] last_be",
		"requester ID":          "output reg [15:0] requester_id",
		"tag":                   "output reg [7:0] tag",
		"traffic class":         "output reg [2:0] traffic_class",
		"attributes":            "output reg [2:0] attributes",
		"BIR":                   "output reg [2:0] bir",
		"enabled byte count":    "output reg [12:0] enabled_byte_count",
		"completion byte count": "output reg [12:0] first_completion_byte_count",
		"planned first CplD":    "output reg [10:0] first_completion_dw",
		"lower address":         "output reg [6:0] first_lower_address",
	})

	if strings.Contains(sv, "{{") || strings.Contains(sv, "}}") {
		t.Error("generated normalizer contains unrendered template delimiters")
	}
}

func TestGenerateTransactionNormalizerHDLFormulaVectors(t *testing.T) {
	sv, err := GenerateTransactionNormalizerSV(testConfig())
	if err != nil {
		t.Fatalf("GenerateTransactionNormalizerSV() error = %v", err)
	}

	requireHDLFragments(t, sv, map[string]string{
		"wire length 0 means 1024 DW":  "length_dw = (tlp_data[9:0] == 10'h000) ? 11'd1024 : {1'b0, tlp_data[9:0]}",
		"3DW and 4DW decode":           "header_4dw = fmt[0]",
		"3DW and 4DW MRd":              "8'b000_00000, 8'b001_00000",
		"3DW and 4DW MWr":              "8'b010_00000, 8'b011_00000",
		"3DW and 4DW MRdLk":            "8'b000_00001, 8'b001_00001: request_kind = KIND_MRDLK",
		"I/O read classification":      "8'b000_00010: request_kind = KIND_IORD",
		"I/O write classification":     "8'b010_00010: request_kind = KIND_IOWR",
		"Atomic classification":        "8'b010_01100, 8'b011_01100",
		"3DW address":                  "{32'h00000000, tlp_data[95:66], 2'b00}",
		"4DW address":                  "{tlp_data[95:64], tlp_data[127:98], 2'b00}",
		"first BE decode":              "first_be = tlp_data[35:32]",
		"last BE decode":               "last_be = tlp_data[39:36]",
		"requester preservation":       "requester_id = tlp_data[63:48]",
		"tag preservation":             "tag = tlp_data[47:40]",
		"TC preservation":              "traffic_class = tlp_data[22:20]",
		"attribute preservation":       "attributes = {tlp_data[18], tlp_data[13:12]}",
		"all 1DW first BE patterns":    "popcount4(first_be)",
		"zero-length 1DW BE legality":  "if (length_dw == 11'd1) legal_be = (last_be == 4'h0)",
		"QW-aligned 2DW BE exception":  "else if ((length_dw == 11'd2) && (address[2:0] == 3'b000)) legal_be = (first_be != 4'h0) && (last_be != 4'h0)",
		"zero-length CplD byte count":  "((length_dw == 11'd1) && (first_be == 4'h0) && memory_read) ? 13'd4",
		"actual zero-byte count":       "enabled_byte_count = (length_dw == 11'd1) ? popcount4(first_be)",
		"completion byte count source": "? 13'd4 : enabled_byte_count",
		"legal multi-DW first BEs":     "(value == 4'h8) || (value == 4'hC) || (value == 4'hE) || (value == 4'hF)",
		"legal multi-DW last BEs":      "(value == 4'h1) || (value == 4'h3) || (value == 4'h7) || (value == 4'hF)",
		"multi-DW interior bytes":      "({(length_dw - 11'd2), 2'b00})",
		"first enabled byte":           "first_lower_address = address[6:0] + first_enabled_offset(first_be)",
		"64 or 128 byte RCB split":     "rcb_bytes_left = READ_COMPLETION_BOUNDARY - (address & (READ_COMPLETION_BOUNDARY - 1))",
		"MPS split":                    "mps_dw = MAX_PAYLOAD_BYTES / 4",
		"RCB minimum":                  "if (first_completion_dw > rcb_dw_left) first_completion_dw = rcb_dw_left",
		"MPS minimum":                  "if (first_completion_dw > mps_dw) first_completion_dw = mps_dw",
		"posted MWr has no completion": "posted_write = memory_write && !malformed_request",
		"unsupported needs UR":         "unsupported_request = 1'b1; ur_required = 1'b1",
		"I/O Atomic non-posted":        "non_posted_request = (request_kind == KIND_IORD) || (request_kind == KIND_IOWR) || (request_kind == KIND_ATOMIC)",
		"MRdLk non-posted":             "(request_kind == KIND_MRDLK)",
		"malformed BE decision":        "malformed_request = !legal_bar || !legal_be || crosses_4k",
	})
}

func TestGenerateURCompleterPreservesMetadataWithoutData(t *testing.T) {
	sv, err := GenerateURCompleterSV(testConfig())
	if err != nil {
		t.Fatalf("GenerateURCompleterSV() error = %v", err)
	}

	requireHDLFragments(t, sv, map[string]string{
		"module":                     "module pcileech_tlp_ur_completer",
		"request valid":              "input request_valid",
		"request ready":              "output request_ready",
		"requester input":            "input [15:0] requester_id",
		"tag input":                  "input [7:0] tag",
		"TC input":                   "input [2:0] traffic_class",
		"attributes input":           "input [2:0] attributes",
		"full tag-space depth":       "parameter integer REQUEST_FIFO_DEPTH = 256",
		"full and pop admission":     "request_can_accept = !request_fifo_full || request_pop",
		"request handshake":          "request_push = request_valid && request_can_accept",
		"no-data Cpl format":         "completion_data[31:29] = 3'b000",
		"completion type":            "completion_data[28:24] = 5'b01010",
		"TC preservation":            "completion_data[22:20] = tc_fifo[read_ptr]",
		"attribute preservation":     "completion_data[18] = attr_fifo[read_ptr][2]",
		"UR completion status":       "3'b001, 1'b0, 12'h000",
		"requester tag preservation": "{requester_fifo[read_ptr], tag_fifo[read_ptr], 8'h00}",
		"no-data Cpl width":          "assign tlps_out.tkeepdw = 4'b0111",
		"queued stream availability": "assign tlps_out.has_data = (request_count != 0)",
	})
}

func TestGenerateBarReadEngineUsesNormalizedVectors(t *testing.T) {
	sv, err := GenerateBarReadEngineSV(testConfig())
	if err != nil {
		t.Fatalf("GenerateBarReadEngineSV() error = %v", err)
	}

	requireHDLFragments(t, sv, map[string]string{
		"normalized address input":     "input [63:0] norm_address",
		"normalized length input":      "input [10:0] norm_length_dw",
		"first BE input":               "input [3:0] norm_first_be",
		"last BE input":                "input [3:0] norm_last_be",
		"requester input":              "input [15:0] norm_requester_id",
		"tag input":                    "input [7:0] norm_tag",
		"TC input":                     "input [2:0] norm_traffic_class",
		"attributes input":             "input [2:0] norm_attributes",
		"enabled bytes input":          "input [12:0] norm_enabled_byte_count",
		"completion bytes input":       "input [12:0] norm_first_completion_byte_count",
		"completion bytes FIFO":        "req_bc_fifo[req_wr_ptr] <= norm_first_completion_byte_count",
		"first completion input":       "input [10:0] norm_first_completion_dw",
		"full tag-space depth":         "localparam integer REQUEST_FIFO_DEPTH = 256",
		"tag descriptor storage":       "reg [7:0] req_tag_fifo [0:REQUEST_FIFO_DEPTH-1]",
		"full descriptor count":        "reg [8:0] req_count",
		"full and pop admission":       "wire request_can_accept = !request_fifo_full || request_pop",
		"lossless request push":        "wire request_push = tlps_in_valid && request_can_accept",
		"simultaneous pop push hold":   "case ({request_push, request_pop})",
		"simultaneous count stable":    "default: req_count <= req_count",
		"partial first decrement":      "value = value - (13'd4 - popcount4(first_enable))",
		"partial last decrement":       "value = value - (13'd4 - popcount4(last_enable))",
		"remaining byte count update":  "remaining_byte_count <= remaining_byte_count - completed_enabled_bytes(",
		"next RCB/MPS split":           "packet_length_dw <= next_completion_dw(",
		"requester context":            "request_id",
		"tag context":                  "request_tag",
		"TC completion header":         "tdata[22:20] <= rd_rsp_tc",
		"attribute completion header":  "tdata[18] <= rd_rsp_attr[2]",
		"byte count completion header": "rd_rsp_byte_count",
		"lower address header":         "rd_rsp_lower_addr",
		"tuser BAR bits driven":        "assign tlps_out.tuser[8:2] = 7'h00",
		"tuser last bit driven":        "assign tlps_out.tuser[1] = tlps_out.tlast",
		"stream has data contract":     "assign tlps_out.has_data = tlps_out.tvalid",
	})

	for _, legacy := range []string{
		"IfAXIS128.sink_lite tlps_in",
		"tlps_in.tdata",
		"tlps_in.tuser",
		"tlps_in.tlast",
	} {
		if strings.Contains(compactHDL(sv), compactHDL(legacy)) {
			t.Errorf("generated read engine still consumes legacy raw request expression %q", legacy)
		}
	}
}

func TestGenerateTransactionLimitsAreSharedAcrossRTL(t *testing.T) {
	cfg := testConfig()
	cfg.ReadCompletionBoundaryBytes = 128
	cfg.MaxPayloadBytes = 256

	normalizer, err := GenerateTransactionNormalizerSV(cfg)
	if err != nil {
		t.Fatalf("GenerateTransactionNormalizerSV() error = %v", err)
	}
	readEngine, err := GenerateBarReadEngineSV(cfg)
	if err != nil {
		t.Fatalf("GenerateBarReadEngineSV() error = %v", err)
	}
	controller, err := GenerateBarControllerSV(cfg)
	if err != nil {
		t.Fatalf("GenerateBarControllerSV() error = %v", err)
	}

	for name, source := range map[string]string{
		"normalizer":  normalizer,
		"read_engine": readEngine,
	} {
		compact := compactHDL(source)
		if !strings.Contains(compact, "parameterintegerREAD_COMPLETION_BOUNDARY=128") {
			t.Errorf("%s RCB parameter does not resolve to 128 bytes", name)
		}
		if !strings.Contains(compact, "parameterintegerMAX_PAYLOAD_BYTES=256") {
			t.Errorf("%s MPS parameter does not resolve to 256 bytes", name)
		}
	}

	compactController := compactHDL(controller)
	if got := strings.Count(compactController, ".READ_COMPLETION_BOUNDARY(128)"); got != 2 {
		t.Errorf("controller RCB override count = %d, want 2", got)
	}
	if got := strings.Count(compactController, ".MAX_PAYLOAD_BYTES(256)"); got != 2 {
		t.Errorf("controller MPS override count = %d, want 2", got)
	}
}

func TestGeneratedBarControllerIntegratesTransactionNormalizer(t *testing.T) {
	sv, err := GenerateBarControllerSV(testConfig())
	if err != nil {
		t.Fatalf("GenerateBarControllerSV() error = %v", err)
	}

	requireHDLFragments(t, sv, map[string]string{
		"normalizer instance":         ") i_tlp_normalizer(",
		"UR stream":                   "IfAXIS128 tlps_ur()",
		"UR completer instance":       "pcileech_tlp_ur_completer i_tlp_ur_completer",
		"UR non-posted gate":          "norm_request_present && norm_ur_required && norm_non_posted_request",
		"UR requester connection":     ".requester_id ( norm_requester_id )",
		"UR tag connection":           ".tag ( norm_tag )",
		"UR TC connection":            ".traffic_class ( norm_traffic_class )",
		"UR attribute connection":     ".attributes ( norm_attributes )",
		"UR output stream":            ".tlps_out ( tlps_ur.source )",
		"UR completion mux":           "cpl_select_ur ? tlps_ur.tdata : tlps_rdeng.tdata",
		"normalized read gate":        "norm_memory_read && norm_request_supported",
		"posted write first beat":     "in_is_first && in_is_wr_ready && norm_posted_write",
		"posted write continuation":   "norm_request_supported || in_is_wr_last",
		"zero-byte write suppression": "norm_enabled_byte_count != 13'd0",
		"read address connection":     ".norm_address ( norm_address )",
		"read length connection":      ".norm_length_dw ( norm_length_dw )",
		"first BE connection":         ".norm_first_be ( norm_first_be )",
		"last BE connection":          ".norm_last_be ( norm_last_be )",
		"requester connection":        ".norm_requester_id ( norm_requester_id )",
		"tag connection":              ".norm_tag ( norm_tag )",
		"TC connection":               ".norm_traffic_class ( norm_traffic_class )",
		"attribute connection":        ".norm_attributes ( norm_attributes )",
		"enabled bytes connection":    ".norm_enabled_byte_count ( norm_enabled_byte_count )",
		"completion bytes connection": ".norm_first_completion_byte_count ( norm_first_completion_byte_count )",
		"completion plan connection":  ".norm_first_completion_dw ( norm_first_completion_dw )",
	})

	for _, legacy := range []string{
		"tlps_in.tdata[31:25] == 7'b0000000",
		"tlps_in.tdata[31:25] == 7'b0010000",
		"tlps_in.tdata[31:24] == 8'b00000010",
	} {
		if strings.Contains(compactHDL(sv), compactHDL(legacy)) {
			t.Errorf("generated controller still uses legacy hardcoded transaction formula %q", legacy)
		}
	}
}
