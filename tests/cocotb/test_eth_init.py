import cocotb
from test_helpers import reset, read_bar, write_bar

CTRL = 0x0000
STATUS = 0x0008
RCTL = 0x0100
TCTL = 0x0400
RDBAL = 0x0200

@cocotb.test()
async def test_eth_ctrl_read(dut):
    await reset(dut)
    val = await read_bar(dut, CTRL, tag=1)
    assert val is not None, "CTRL no completion"

@cocotb.test()
async def test_eth_status_read(dut):
    await reset(dut)
    val = await read_bar(dut, STATUS, tag=2)
    assert val is not None, "STATUS no completion"

@cocotb.test()
async def test_eth_bram_write_readback(dut):
    await reset(dut)
    await write_bar(dut, RCTL, 0x0000003E)
    await write_bar(dut, TCTL, 0x000000FF)
    rctl = await read_bar(dut, RCTL, tag=3)
    tctl = await read_bar(dut, TCTL, tag=4)
    assert rctl == 0x3E, f"RCTL readback: {rctl:#x}"
    assert tctl == 0xFF, f"TCTL readback: {tctl:#x}"

@cocotb.test()
async def test_eth_descriptor_base_write(dut):
    await reset(dut)
    await write_bar(dut, RDBAL, 0xDEADBEEF)
    val = await read_bar(dut, RDBAL, tag=5)
    assert val == 0xDEADBEEF, f"RDBAL readback: {val:#x}"

@cocotb.test()
async def test_eth_full_driver_init(dut):
    await reset(dut)
    assert await read_bar(dut, CTRL, tag=10) is not None
    assert await read_bar(dut, STATUS, tag=11) is not None
    await write_bar(dut, RCTL, 0x3E)
    await write_bar(dut, TCTL, 0xFF)
    assert await read_bar(dut, RCTL, tag=12) == 0x3E
    assert await read_bar(dut, TCTL, tag=13) == 0xFF
