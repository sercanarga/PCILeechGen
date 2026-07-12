import cocotb
from cocotb.triggers import RisingEdge
from test_helpers import reset, read_bar, write_bar

GCAP = 0x00
GCTL = 0x08
WAKEEN = 0x0C
INTCTL = 0x20
CORBLBASE = 0x40
CORBWP = 0x48
CORBCTL = 0x4C
RIRBLBASE = 0x50
RIRBWP = 0x58
RIRBCTL = 0x5C
RIRBINTSTS = 0x60
RIRBRESP_LO = 0x70

GCAP_RESET = 0x01006401

@cocotb.test()
async def test_hda_gcap_read(dut):
    await reset(dut)
    val = await read_bar(dut, GCAP, tag=1)
    assert val is not None, "GCAP no completion"
    assert val == GCAP_RESET, f"GCAP={val:#x}, want {GCAP_RESET:#x}"

@cocotb.test()
async def test_hda_gctl_crst_handshake(dut):
    await reset(dut)
    g = await read_bar(dut, GCTL, tag=2)
    assert g is not None and (g & 1), f"CRST not set after reset: {g:#x}"
    await write_bar(dut, GCTL, 0x00000000)
    for _ in range(20): await RisingEdge(dut.clk)
    ss = await read_bar(dut, WAKEEN, tag=3)
    assert (ss >> 16) & 0xFFFF == 0, f"STATESTS not cleared on CRST=0: {ss:#x}"
    await write_bar(dut, GCTL, 0x00000001)
    for _ in range(20): await RisingEdge(dut.clk)
    ss = await read_bar(dut, WAKEEN, tag=4)
    assert (ss >> 16) & 1 == 1, f"codec-0 not discovered after CRST=1: {ss:#x}"

@cocotb.test()
async def test_hda_corb_setup(dut):
    await reset(dut)
    await write_bar(dut, GCTL, 0x00000001)
    for _ in range(20): await RisingEdge(dut.clk)
    await write_bar(dut, CORBLBASE, 0x1000)
    await write_bar(dut, CORBWP, 0x80000000)
    for _ in range(20): await RisingEdge(dut.clk)
    await write_bar(dut, CORBCTL, 0x00000003)
    for _ in range(20): await RisingEdge(dut.clk)
    val = await read_bar(dut, CORBCTL, tag=5)
    assert val is not None and val & 0x3 == 0x3, f"CORBRUN not enabled: {val:#x}"

@cocotb.test()
async def test_hda_rirb_setup(dut):
    await reset(dut)
    await write_bar(dut, GCTL, 0x00000001)
    for _ in range(20): await RisingEdge(dut.clk)
    await write_bar(dut, RIRBLBASE, 0x2000)
    await write_bar(dut, RIRBWP, 0x00018000)
    for _ in range(20): await RisingEdge(dut.clk)
    await write_bar(dut, RIRBCTL, 0x00000007)
    for _ in range(20): await RisingEdge(dut.clk)
    val = await read_bar(dut, RIRBCTL, tag=6)
    assert val is not None and val & 0x7 == 0x7, f"RIRBDMAEN not enabled: {val:#x}"

@cocotb.test()
async def test_hda_interrupt_enable(dut):
    await reset(dut)
    await write_bar(dut, INTCTL, 0xC0000000)
    for _ in range(20): await RisingEdge(dut.clk)
    val = await read_bar(dut, INTCTL, tag=7)
    assert val is not None and (val >> 30) & 0x3 == 0x3, f"GIE/CIE not set: {val:#x}"

@cocotb.test()
async def test_hda_corbwp_advances_rirb(dut):
    await reset(dut)
    await write_bar(dut, GCTL, 0x00000001)
    for _ in range(20): await RisingEdge(dut.clk)
    await write_bar(dut, CORBLBASE, 0x1000)
    await write_bar(dut, RIRBLBASE, 0x2000)
    await write_bar(dut, CORBCTL, 0x00000003)
    await write_bar(dut, RIRBCTL, 0x00000007)
    for _ in range(20): await RisingEdge(dut.clk)
    await write_bar(dut, CORBWP, 0x00000001)
    for _ in range(2000): await RisingEdge(dut.clk)
    wp = await read_bar(dut, RIRBWP, tag=8)
    intfl = await read_bar(dut, RIRBINTSTS, tag=9)
    resp = await read_bar(dut, RIRBRESP_LO, tag=10)
    assert wp is not None and (wp & 0xFF) >= 1, f"RIRBWP not advanced: {wp:#x}"
    assert intfl is not None and intfl & 1, f"INTFL not set: {intfl:#x}"

@cocotb.test()
async def test_hda_full_driver_init(dut):
    await reset(dut)
    assert await read_bar(dut, GCAP, tag=20) == GCAP_RESET
    await write_bar(dut, GCTL, 0x00000000)
    for _ in range(20): await RisingEdge(dut.clk)
    await write_bar(dut, GCTL, 0x00000001)
    for _ in range(20): await RisingEdge(dut.clk)
    await write_bar(dut, CORBLBASE, 0x1000)
    await write_bar(dut, CORBCTL, 0x00000003)
    await write_bar(dut, RIRBLBASE, 0x2000)
    await write_bar(dut, RIRBCTL, 0x00000007)
    await write_bar(dut, INTCTL, 0xC0000000)
    for _ in range(20): await RisingEdge(dut.clk)
    await write_bar(dut, CORBWP, 0x00000001)
    for _ in range(2000): await RisingEdge(dut.clk)
    intfl = await read_bar(dut, RIRBINTSTS, tag=21)
    assert intfl is not None and intfl & 1, "RIRB INTFL not set after full init"
