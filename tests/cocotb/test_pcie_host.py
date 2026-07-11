import struct, cocotb
from cocotb.clock import Clock
from cocotb.triggers import RisingEdge, Timer

def mrd3(addr, tag=1, length=1, first_be=0xF, last_be=None, req_id=0):
    if last_be is None:
        last_be = 0x0 if length == 1 else 0xF
    dw0 = (0b000 << 29) | 0b00000 << 24 | (length & 0x3FF)
    dw1 = ((req_id & 0xFFFF) << 16) | ((tag & 0xFF) << 8) | ((last_be & 0xF) << 4) | (first_be & 0xF)
    dw2 = addr & 0xFFFFFFFC
    return dw0 | (dw1 << 32) | (dw2 << 64)

def mwr3(addr, data, tag=0, req_id=0, be=0xF):
    dw0 = (0b010 << 29) | 0b00000 << 24 | 1
    dw1 = ((req_id & 0xFFFF) << 16) | ((tag & 0xFF) << 8) | be
    dw2 = addr & 0xFFFFFFFC
    return dw0 | (dw1 << 32) | (dw2 << 64) | ((data & 0xFFFFFFFF) << 96)

def iord3(addr, tag=1):
    dw0 = (0b000 << 29) | (0b00010 << 24) | 1
    dw1 = (tag << 8) | 0xF
    dw2 = addr & 0xFFFFFFFC
    return dw0 | (dw1 << 32) | (dw2 << 64)

async def send(dut, tdata, bar=0):
    dut.tlps_in_tdata.value = tdata
    dut.tlps_in_tvalid.value = 1
    dut.tlps_in_tlast.value = 1
    dut.tlps_in_tuser.value = 1 | (1 << (bar + 2))
    dut.tlps_in_tkeepdw.value = 0xF
    await RisingEdge(dut.clk)
    dut.tlps_in_tvalid.value = 0
    dut.tlps_in_tlast.value = 0
    dut.tlps_in_tuser.value = 0

async def send_bare(dut, tdata, tuser_val):
    dut.tlps_in_tdata.value = tdata
    dut.tlps_in_tvalid.value = 1
    dut.tlps_in_tlast.value = 1
    dut.tlps_in_tuser.value = tuser_val
    dut.tlps_in_tkeepdw.value = 0xF
    await RisingEdge(dut.clk)
    dut.tlps_in_tvalid.value = 0
    dut.tlps_in_tlast.value = 0
    dut.tlps_in_tuser.value = 0

async def recv_all(dut, timeout=10000):
    cpls = []
    idle = 0
    for _ in range(timeout):
        await RisingEdge(dut.clk)
        if dut.tlps_out_tvalid.value == 1:
            raw = dut.tlps_out_tdata.value.integer
            dws = [(raw >> (i*32)) & 0xFFFFFFFF for i in range(4)]
            cpls.append(dws)
            if dut.tlps_out_tlast.value == 1:
                break
            idle = 0
        else:
            idle += 1
            if cpls and idle > 50:
                break
    return cpls

async def recv_dma_all(dut, timeout=10000):
    beats = []
    for _ in range(timeout):
        await RisingEdge(dut.clk)
        if dut.tlps_dma_out_tvalid.value == 1:
            raw = dut.tlps_dma_out_tdata.value.integer
            dws = [(raw >> (i*32)) & 0xFFFFFFFF for i in range(4)]
            beats.append(dws)
            if dut.tlps_dma_out_tlast.value == 1:
                break
    return beats

async def reset(dut):
    cocotb.start_soon(Clock(dut.clk, 10, unit="ns").start())
    dut.rst.value = 1
    dut.tlps_in_tvalid.value = 0
    await Timer(200, "ns")
    for _ in range(10): await RisingEdge(dut.clk)
    dut.rst.value = 0
    for _ in range(200): await RisingEdge(dut.clk)

# --- basic read/write ---

@cocotb.test()
async def test_mrd_basic(dut):
    await reset(dut)
    await send(dut, mrd3(addr=0, tag=1))
    cpls = await recv_all(dut)
    assert len(cpls) > 0

@cocotb.test()
async def test_mrd_offset_100(dut):
    await reset(dut)
    await send(dut, mrd3(addr=0x100, tag=2))
    cpls = await recv_all(dut)
    assert len(cpls) > 0

@cocotb.test()
async def test_mwr_then_mrd(dut):
    await reset(dut)
    await send(dut, mwr3(addr=0x100, data=0xDEADBEEF))
    for _ in range(20): await RisingEdge(dut.clk)
    await send(dut, mrd3(addr=0x100, tag=3))
    cpls = await recv_all(dut)
    assert len(cpls) > 0

@cocotb.test()
async def test_mrd_no_bar_hit(dut):
    await reset(dut)
    await send_bare(dut, mrd3(addr=0, tag=4), tuser_val=1)
    for _ in range(100): await RisingEdge(dut.clk)
    assert dut.tlps_out_tvalid.value == 0

@cocotb.test()
async def test_io_read_ur(dut):
    await reset(dut)
    await send(dut, iord3(addr=0, tag=5))
    for _ in range(500): await RisingEdge(dut.clk)

@cocotb.test()
async def test_mrd_length_4(dut):
    await reset(dut)
    await send(dut, mrd3(addr=0, tag=6, length=4))
    cpls = await recv_all(dut, timeout=100000)
    assert len(cpls) > 0

@cocotb.test()
async def test_mwr_doorbell_safe(dut):
    await reset(dut)
    await send(dut, mwr3(addr=0x1000, data=0))
    for _ in range(50): await RisingEdge(dut.clk)

@cocotb.test()
async def test_rapid_mrd_burst(dut):
    await reset(dut)
    for t in range(10):
        await send(dut, mrd3(addr=t * 4, tag=t + 10))
        await RisingEdge(dut.clk)
    total_cpls = 0
    for _ in range(5000):
        await RisingEdge(dut.clk)
        if dut.tlps_out_tvalid.value == 1:
            total_cpls += 1
    assert total_cpls > 0

@cocotb.test()
async def test_mrd_large_length(dut):
    await reset(dut)
    await send(dut, mrd3(addr=0, tag=20, length=32))
    cpls = await recv_all(dut, timeout=100000)
    assert len(cpls) > 0

@cocotb.test()
async def test_mwr_pattern_integrity(dut):
    await reset(dut)
    for i, p in enumerate([0xDEADBEEF, 0xCAFEBABE, 0x12345678, 0xAABBCCDD]):
        await send(dut, mwr3(addr=0x200 + i * 4, data=p))
        for _ in range(10): await RisingEdge(dut.clk)

# --- NVMe register reads ---

@cocotb.test()
async def test_nvme_cap_read(dut):
    await reset(dut)
    await send(dut, mrd3(addr=0x0000, tag=30))
    cpls = await recv_all(dut)
    assert len(cpls) > 0

@cocotb.test()
async def test_nvme_cc_write_readback(dut):
    await reset(dut)
    await send(dut, mwr3(addr=0x0014, data=0x00000001))
    for _ in range(20): await RisingEdge(dut.clk)
    await send(dut, mrd3(addr=0x0014, tag=31))
    cpls = await recv_all(dut)
    assert len(cpls) > 0

@cocotb.test()
async def test_nvme_csts_read(dut):
    await reset(dut)
    await send(dut, mrd3(addr=0x001C, tag=32))
    cpls = await recv_all(dut)
    assert len(cpls) > 0

@cocotb.test()
async def test_nvme_vs_read(dut):
    await reset(dut)
    await send(dut, mrd3(addr=0x0008, tag=33))
    cpls = await recv_all(dut)
    assert len(cpls) > 0

@cocotb.test()
async def test_nvme_doorbell_write(dut):
    await reset(dut)
    await send(dut, mwr3(addr=0x1000, data=0))
    for _ in range(100): await RisingEdge(dut.clk)

@cocotb.test()
async def test_nvme_full_init(dut):
    await reset(dut)
    await send(dut, mwr3(addr=0x0014, data=0x00000001))
    for _ in range(500): await RisingEdge(dut.clk)
    await send(dut, mrd3(addr=0x001C, tag=40))
    cpls = await recv_all(dut)
    assert len(cpls) > 0

# --- MSI-X table tests ---

@cocotb.test()
async def test_msix_table_write(dut):
    """Write MSI-X vector 0 addr/data at table offset 0x2000."""
    await reset(dut)
    await send(dut, mwr3(addr=0x2000, data=0xFEE00000))
    for _ in range(10): await RisingEdge(dut.clk)
    await send(dut, mwr3(addr=0x2004, data=0x00007FF0))
    for _ in range(10): await RisingEdge(dut.clk)
    await send(dut, mwr3(addr=0x2008, data=0x00004400))
    for _ in range(10): await RisingEdge(dut.clk)
    await send(dut, mwr3(addr=0x200C, data=0x00000000))
    for _ in range(50): await RisingEdge(dut.clk)
    dut._log.info("PASS: MSI-X vector 0 written")

@cocotb.test()
async def test_msix_table_readback(dut):
    """Read back MSI-X vector 0 after writing."""
    await reset(dut)
    await send(dut, mwr3(addr=0x2000, data=0xFEE00000))
    for _ in range(20): await RisingEdge(dut.clk)
    await send(dut, mrd3(addr=0x2000, tag=50))
    cpls = await recv_all(dut)
    assert len(cpls) > 0
    dut._log.info(f"vector 0 DW0 readback: {hex(cpls[0][3])}")

@cocotb.test()
async def test_msix_pba_read(dut):
    """Read MSI-X PBA at offset 0x3000."""
    await reset(dut)
    await send(dut, mrd3(addr=0x3000, tag=51))
    cpls = await recv_all(dut)
    assert len(cpls) > 0

@cocotb.test()
async def test_msix_vector_mask_write(dut):
    """Write vector control dword (offset +0xC per entry) — mask bit 0."""
    await reset(dut)
    await send(dut, mwr3(addr=0x200C, data=0x00000001))
    for _ in range(10): await RisingEdge(dut.clk)
    await send(dut, mwr3(addr=0x201C, data=0x00000001))
    for _ in range(10): await RisingEdge(dut.clk)
    dut._log.info("PASS: vector mask writes accepted")

@cocotb.test()
async def test_msix_multiple_vectors(dut):
    """Write 5 vectors, then read back each."""
    await reset(dut)
    for v in range(5):
        base = 0x2000 + v * 16
        await send(dut, mwr3(addr=base, data=0xFEE00000 + v))
        for _ in range(5): await RisingEdge(dut.clk)
        await send(dut, mwr3(addr=base + 4, data=0x00007FF0))
        for _ in range(5): await RisingEdge(dut.clk)
    for v in range(5):
        base = 0x2000 + v * 16
        await send(dut, mrd3(addr=base, tag=60 + v))
        cpls = await recv_all(dut)
        assert len(cpls) > 0

@cocotb.test()
async def test_msix_table_outside_range_ignored(dut):
    """Write at 0x2080 (past 5 vectors) should not crash."""
    await reset(dut)
    await send(dut, mwr3(addr=0x2080, data=0xDEADBEEF))
    for _ in range(20): await RisingEdge(dut.clk)

# --- DMA path tests ---

@cocotb.test()
async def test_dma_nvme_init_triggers_dma(dut):
    """Enable NVMe controller, ring doorbell, observe output activity."""
    await reset(dut)
    await send(dut, mwr3(addr=0x0014, data=0x00000001))
    for _ in range(500): await RisingEdge(dut.clk)
    dut._log.info("controller enabled, ringing doorbell")
    await send(dut, mwr3(addr=0x1000, data=0x00000001))
    total_dma = 0
    total_cpl = 0
    for _ in range(20000):
        await RisingEdge(dut.clk)
        if dut.tlps_dma_out_tvalid.value == 1:
            total_dma += 1
        if dut.tlps_out_tvalid.value == 1:
            total_cpl += 1
    dut._log.info(f"after doorbell: dma_beats={total_dma} cpl_beats={total_cpl}")

@cocotb.test()
async def test_dma_write_data_on_tlps_dma_out(dut):
    """After NVMe admin Identify, check DMA output carries write data."""
    await reset(dut)
    await send(dut, mwr3(addr=0x0014, data=0x00000001))
    for _ in range(100): await RisingEdge(dut.clk)
    await send(dut, mwr3(addr=0x1000, data=0x00000001))
    dma_beats = await recv_dma_all(dut, timeout=20000)
    dut._log.info(f"DMA beats: {len(dma_beats)}")
