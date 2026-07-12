import cocotb
from test_helpers import reset, read_bar, write_bar

CSR = 0x000000
GP_CTL = 0x000024
HW_REV = 0x000028
UCODE_DRV_GP1 = 0x000054
RF_ID = 0x00009C

@cocotb.test()
async def test_wifi_csr_read(dut):
    await reset(dut)
    val = await read_bar(dut, CSR, tag=1)
    assert val is not None, "CSR no completion"

@cocotb.test()
async def test_wifi_hw_rev_read(dut):
    await reset(dut)
    val = await read_bar(dut, HW_REV, tag=2)
    assert val is not None, "HW_REV no completion"

@cocotb.test()
async def test_wifi_rf_id_read(dut):
    await reset(dut)
    val = await read_bar(dut, RF_ID, tag=3)
    assert val is not None, "RF_ID no completion"

@cocotb.test()
async def test_wifi_gp_ctl_write_readback(dut):
    await reset(dut)
    await write_bar(dut, GP_CTL, 0x00000100)
    val = await read_bar(dut, GP_CTL, tag=4)
    assert val is not None and val == 0x100, f"GP_CTL readback: {val:#x}"

@cocotb.test()
async def test_wifi_ucode_write_readback(dut):
    await reset(dut)
    await write_bar(dut, UCODE_DRV_GP1, 0x00000007)
    val = await read_bar(dut, UCODE_DRV_GP1, tag=5)
    assert val is not None and val == 0x07, f"UCODE_DRV_GP1 readback: {val:#x}"

@cocotb.test()
async def test_wifi_ro_write_ignored(dut):
    await reset(dut)
    await write_bar(dut, CSR, 0xDEADBEEF)
    val = await read_bar(dut, CSR, tag=6)
    assert val is not None and val == 0, f"RO CSR changed: {val:#x}"

@cocotb.test()
async def test_wifi_full_driver_init(dut):
    await reset(dut)
    assert await read_bar(dut, CSR, tag=10) is not None
    assert await read_bar(dut, HW_REV, tag=11) is not None
    assert await read_bar(dut, RF_ID, tag=12) is not None
    await write_bar(dut, GP_CTL, 0x100)
    assert await read_bar(dut, GP_CTL, tag=13) == 0x100
    await write_bar(dut, UCODE_DRV_GP1, 0x07)
    assert await read_bar(dut, UCODE_DRV_GP1, tag=14) == 0x07
