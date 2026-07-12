import cocotb
from cocotb.triggers import RisingEdge
from test_helpers import reset, read_bar, write_bar

USBCMD = 0x20
USBSTS = 0x24
PORTSC1 = 0x420
CONFIG = 0x58
CRCR_LO = 0x38
DCBAAP_LO = 0x50

@cocotb.test()
async def test_xhci_capability_reads(dut):
    await reset(dut)
    for off, name in [(0x00, "CAPLENGTH"), (0x04, "HCSPARAMS1"), (0x10, "HCCPARAMS1"),
                      (0x14, "DBOFF"), (0x18, "RTSOFF")]:
        val = await read_bar(dut, off, tag=10)
        assert val is not None, f"no completion for {name}"

@cocotb.test()
async def test_xhci_halt_then_run(dut):
    await reset(dut)
    cmd = await read_bar(dut, USBCMD, tag=20)
    assert cmd is not None and (cmd & 0x01), f"USBCMD reset: {cmd:#x}"
    await write_bar(dut, USBCMD, 0x00)
    for _ in range(50):
        await RisingEdge(dut.clk)
        sts = await read_bar(dut, USBSTS, tag=21)
        if sts is not None and (sts & 0x01):
            break
    else:
        assert False, "USBSTS.HCH not set after halt"
    await write_bar(dut, USBCMD, 0x01)
    for _ in range(50):
        await RisingEdge(dut.clk)
        sts = await read_bar(dut, USBSTS, tag=22)
        if sts is not None and not (sts & 0x01):
            break
    else:
        assert False, "USBSTS.HCH not cleared after run"

@cocotb.test()
async def test_xhci_portsc_read(dut):
    await reset(dut)
    p1 = await read_bar(dut, PORTSC1, tag=30)
    assert p1 is not None, "PORTSC1 no completion"

@cocotb.test()
async def test_xhci_config_write_readback(dut):
    await reset(dut)
    await write_bar(dut, CONFIG, 0x24)
    val = await read_bar(dut, CONFIG, tag=40)
    assert val is not None and (val & 0xFF) == 0x24, f"CONFIG readback: {val:#x}"

@cocotb.test()
async def test_xhci_full_driver_init(dut):
    await reset(dut)
    assert await read_bar(dut, 0x00, tag=50) is not None
    await write_bar(dut, USBCMD, 0x00)
    for _ in range(50):
        await RisingEdge(dut.clk)
        if await read_bar(dut, USBSTS, tag=51) is not None: break
    await write_bar(dut, CRCR_LO, 0x0000A000)
    await write_bar(dut, DCBAAP_LO, 0x0000B000)
    await write_bar(dut, CONFIG, 0x20)
    await write_bar(dut, USBCMD, 0x05)
    for _ in range(50):
        await RisingEdge(dut.clk)
        sts = await read_bar(dut, USBSTS, tag=52)
        if sts is not None and not (sts & 0x01): break
    p1 = await read_bar(dut, PORTSC1, tag=53)
    assert p1 is not None, "PORTSC1 no completion after init"
