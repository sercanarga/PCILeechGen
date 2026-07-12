import cocotb
from test_helpers import reset, read_bar, write_bar

PMC_BOOT_0 = 0x000000
PMC_ENABLE = 0x000200
PBUS_PCI_NV_1 = 0x001804
PTIMER_TIME_0 = 0x009400

@cocotb.test()
async def test_gpu_pmc_boot_read(dut):
    await reset(dut)
    val = await read_bar(dut, PMC_BOOT_0, tag=1)
    assert val is not None, "PMC_BOOT_0 no completion"

@cocotb.test()
async def test_gpu_pmc_enable_read_default(dut):
    await reset(dut)
    val = await read_bar(dut, PMC_ENABLE, tag=2)
    assert val is not None, "PMC_ENABLE no completion"

@cocotb.test()
async def test_gpu_pmc_enable_write_readback(dut):
    await reset(dut)
    await write_bar(dut, PMC_ENABLE, 0x00000133)
    val = await read_bar(dut, PMC_ENABLE, tag=3)
    assert val is not None and val == 0x00000133, f"PMC_ENABLE readback: {val:#x}"

@cocotb.test()
async def test_gpu_ro_write_ignored(dut):
    await reset(dut)
    await write_bar(dut, PMC_BOOT_0, 0xDEADBEEF)
    val = await read_bar(dut, PMC_BOOT_0, tag=4)
    assert val is not None and val == 0, f"RO register changed: {val:#x}"

@cocotb.test()
async def test_gpu_pbus_read(dut):
    await reset(dut)
    val = await read_bar(dut, PBUS_PCI_NV_1, tag=5)
    assert val is not None, "PBUS_PCI_NV_1 no completion"

@cocotb.test()
async def test_gpu_unregistered_offset_zero(dut):
    await reset(dut)
    val = await read_bar(dut, 0x100, tag=6)
    assert val is not None and val == 0, f"unregistered offset nonzero: {val:#x}"

@cocotb.test()
async def test_gpu_full_driver_init(dut):
    await reset(dut)
    boot = await read_bar(dut, PMC_BOOT_0, tag=10)
    assert boot is not None
    await write_bar(dut, PMC_ENABLE, 0x133)
    en = await read_bar(dut, PMC_ENABLE, tag=11)
    assert en == 0x133, f"PMC_ENABLE write failed: {en:#x}"
    pbus = await read_bar(dut, PBUS_PCI_NV_1, tag=12)
    assert pbus is not None
