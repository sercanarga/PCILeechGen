import cocotb
from test_helpers import reset, read_bar, write_bar

CAP = 0x00
GHC = 0x04
VS = 0x10
PI = 0x0C
P0CMD = 0x118
P0SSTS = 0x128
ABAR = 5

@cocotb.test()
async def test_ahci_cap_read(dut):
    await reset(dut)
    val = await read_bar(dut, CAP, tag=1, bar=ABAR)
    assert val is not None, "CAP no completion"

@cocotb.test()
async def test_ahci_ghc_bram_write_readback(dut):
    await reset(dut)
    await write_bar(dut, GHC, 0x80000000, bar=ABAR)
    val = await read_bar(dut, GHC, tag=2, bar=ABAR)
    assert val is not None and val == 0x80000000, f"GHC readback: {val:#x}"

@cocotb.test()
async def test_ahci_version_read(dut):
    await reset(dut)
    val = await read_bar(dut, VS, tag=3, bar=ABAR)
    assert val is not None, "VS no completion"

@cocotb.test()
async def test_ahci_pi_read(dut):
    await reset(dut)
    val = await read_bar(dut, PI, tag=4, bar=ABAR)
    assert val is not None, "PI no completion"

@cocotb.test()
async def test_ahci_port0_read(dut):
    await reset(dut)
    pcmd = await read_bar(dut, P0CMD, tag=5, bar=ABAR)
    pssts = await read_bar(dut, P0SSTS, tag=6, bar=ABAR)
    assert pcmd is not None, "P0CMD no completion"
    assert pssts is not None, "P0SSTS no completion"

@cocotb.test()
async def test_ahci_bram_full_init(dut):
    await reset(dut)
    assert await read_bar(dut, CAP, tag=10, bar=ABAR) is not None
    await write_bar(dut, GHC, 0x80000000, bar=ABAR)
    assert await read_bar(dut, GHC, tag=11, bar=ABAR) == 0x80000000
    await write_bar(dut, P0CMD, 0x00000011, bar=ABAR)
    assert await read_bar(dut, P0CMD, tag=12, bar=ABAR) == 0x00000011
