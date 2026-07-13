import cocotb
from cocotb.triggers import RisingEdge
from test_helpers import reset, read_bar, write_bar

CTRL = 0x0000
STATUS = 0x0008
RCTL = 0x0100
TCTL = 0x0400
RDBAL = 0x2800
RDBAL_DMA = RDBAL
RDLEN_DMA = 0x2808
RDH_DMA = 0x2810
RDT_DMA = 0x2818
TDBAL_DMA = 0x3800
TDLEN_DMA = 0x3808
TDH_DMA = 0x3810
TDT_DMA = 0x3818
IMS_DMA = 0x00D0
ICR_DMA = 0x00C0
ICS_DMA = 0x00C8
IMC_DMA = 0x00D8


async def host_poke(dut, word_index, value):
    dut.host_poke_addr.value = word_index & 0xFFFF
    dut.host_poke_data.value = value & 0xFFFFFFFF
    dut.host_poke_valid.value = 1
    await RisingEdge(dut.clk)
    dut.host_poke_valid.value = 0


async def host_peek(dut, word_index):
    dut.host_peek_addr.value = word_index & 0xFFFF
    await RisingEdge(dut.clk)
    await RisingEdge(dut.clk)
    return dut.host_peek_data.value.integer

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
    assert val == 0xDEADBE80, f"RDBAL alignment mask: {val:#x}"


@cocotb.test()
async def test_eth_interrupt_cause_and_mask_semantics(dut):
    await reset(dut)
    await write_bar(dut, IMS_DMA, 0x80)
    await write_bar(dut, IMS_DMA, 0x01)
    assert await read_bar(dut, IMS_DMA, tag=6) == 0x81
    await write_bar(dut, IMC_DMA, 0x01)
    assert await read_bar(dut, IMS_DMA, tag=7) == 0x80
    await write_bar(dut, ICS_DMA, 0x80)
    assert await read_bar(dut, ICR_DMA, tag=8) == 0x80
    assert await read_bar(dut, ICR_DMA, tag=9) == 0


@cocotb.test()
async def test_eth_full_driver_init(dut):
    await reset(dut)
    assert await read_bar(dut, CTRL, tag=10) is not None
    assert await read_bar(dut, STATUS, tag=11) is not None
    await write_bar(dut, RCTL, 0x3E)
    await write_bar(dut, TCTL, 0xFF)
    assert await read_bar(dut, RCTL, tag=12) == 0x3E
    assert await read_bar(dut, TCTL, tag=13) == 0xFF


@cocotb.test()
async def test_eth_dma_fake_rx_and_tx_loopback(dut):
    await reset(dut)

    # Intel rings have 128-byte minimum length (eight 16-byte descriptors).
    await host_poke(dut, 0x1000 // 4, 0x00005000)
    await host_poke(dut, 0x1000 // 4 + 4, 0x00005010)
    await host_poke(dut, 0x2000 // 4, 0x00006000)
    await host_poke(dut, 0x2000 // 4 + 2, 0x01000004)  # length=4, software-owned EOP
    await host_poke(dut, 0x6000 // 4, 0x474E4950)  # "PING" in host byte order
    await write_bar(dut, RDBAL_DMA, 0x1000)
    await write_bar(dut, RDLEN_DMA, 128)
    await write_bar(dut, TDBAL_DMA, 0x2000)
    await write_bar(dut, TDLEN_DMA, 128)
    await write_bar(dut, IMS_DMA, 1 << 7)

    # RDT=1 causes firmware to DMA the synthetic ARP frame to RX buffer 0.
    await write_bar(dut, RDT_DMA, 1)
    for _ in range(2500):
        await RisingEdge(dut.clk)
    assert await host_peek(dut, 0x5000 // 4) == 0xFFFFFFFF
    assert await host_peek(dut, 0x5000 // 4 + 3) == 0x01000608
    assert await host_peek(dut, 0x1000 // 4 + 2) == 60
    assert await host_peek(dut, 0x1000 // 4 + 3) == 3
    assert await read_bar(dut, RDH_DMA, tag=20) == 1

    # Make RX descriptor 1 available, then submit TX descriptor 0.
    await write_bar(dut, RDT_DMA, 2)
    await write_bar(dut, TDT_DMA, 1)
    for _ in range(4000):
        await RisingEdge(dut.clk)
    assert await host_peek(dut, 0x5010 // 4) == 0x474E4950
    assert await host_peek(dut, 0x1010 // 4 + 2) == 4
    assert await host_peek(dut, 0x1010 // 4 + 3) == 3
    assert await host_peek(dut, 0x2000 // 4 + 2) == 0x01000004
    assert await host_peek(dut, 0x2000 // 4 + 3) == 1
    assert await read_bar(dut, TDH_DMA, tag=21) == 1
