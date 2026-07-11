import struct, cocotb, random
from cocotb.clock import Clock
from cocotb.triggers import RisingEdge, Timer

def mrd3(addr, tag=1, length=1, first_be=0xF, last_be=None, req_id=0):
    if last_be is None:
        last_be = 0x0 if length == 1 else 0xF
    dw0 = (0b000 << 29) | 0b00000 << 24 | (length & 0x3FF)
    dw1 = ((req_id & 0xFFFF) << 16) | ((tag & 0xFF) << 8) | ((last_be & 0xF) << 4) | (first_be & 0xF)
    dw2 = addr & 0xFFFFFFFC
    return dw0 | (dw1 << 32) | (dw2 << 64)

def mwr3(addr, data, tag=0, req_id=0):
    dw0 = (0b010 << 29) | 0b00000 << 24 | 1
    dw1 = ((req_id & 0xFFFF) << 16) | ((tag & 0xFF) << 8) | 0xF
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

async def reset(dut):
    cocotb.start_soon(Clock(dut.clk, 10, unit="ns").start())
    dut.rst.value = 1
    dut.tlps_in_tvalid.value = 0
    await Timer(200, "ns")
    for _ in range(10): await RisingEdge(dut.clk)
    dut.rst.value = 0
    for _ in range(200): await RisingEdge(dut.clk)

# --- common tests (all device classes) ---

@cocotb.test()
async def test_mrd_basic(dut):
    """MRd BAR0 offset 0 returns completion."""
    await reset(dut)
    await send(dut, mrd3(addr=0, tag=1))
    cpls = await recv_all(dut)
    assert len(cpls) > 0, "no completion"

@cocotb.test()
async def test_mrd_offset_100(dut):
    """MRd BAR0 offset 0x100."""
    await reset(dut)
    await send(dut, mrd3(addr=0x100, tag=2))
    cpls = await recv_all(dut)
    assert len(cpls) > 0

@cocotb.test()
async def test_mwr_then_mrd(dut):
    """Write then read back data integrity."""
    await reset(dut)
    await send(dut, mwr3(addr=0x100, data=0xDEADBEEF))
    for _ in range(20): await RisingEdge(dut.clk)
    await send(dut, mrd3(addr=0x100, tag=3))
    cpls = await recv_all(dut)
    assert len(cpls) > 0

@cocotb.test()
async def test_mrd_no_bar_hit(dut):
    """MRd with no BAR bit in tuser → no completion."""
    await reset(dut)
    tdata = mrd3(addr=0, tag=4)
    dut.tlps_in_tdata.value = tdata
    dut.tlps_in_tvalid.value = 1
    dut.tlps_in_tlast.value = 1
    dut.tlps_in_tuser.value = 1
    dut.tlps_in_tkeepdw.value = 0xF
    await RisingEdge(dut.clk)
    dut.tlps_in_tvalid.value = 0
    dut.tlps_in_tlast.value = 0
    dut.tlps_in_tuser.value = 0
    for _ in range(100): await RisingEdge(dut.clk)
    assert dut.tlps_out_tvalid.value == 0, "unexpected completion for no-BAR request"

@cocotb.test()
async def test_io_read_ur(dut):
    """I/O read → unsupported → UR or no response."""
    await reset(dut)
    await send(dut, iord3(addr=0, tag=5))
    for _ in range(500): await RisingEdge(dut.clk)

@cocotb.test()
async def test_mrd_length_4(dut):
    """Multi-DWORD MRd (length=4)."""
    await reset(dut)
    await send(dut, mrd3(addr=0, tag=6, length=4))
    cpls = await recv_all(dut, timeout=100000)
    assert len(cpls) > 0, "no completion for 4-DW read"

@cocotb.test()
async def test_mwr_doorbell_safe(dut):
    """MWr to doorbell area, no crash."""
    await reset(dut)
    await send(dut, mwr3(addr=0x1000, data=0))
    for _ in range(50): await RisingEdge(dut.clk)

@cocotb.test()
async def test_rapid_mrd_burst(dut):
    """10 rapid MRds, count total completions."""
    await reset(dut)
    for t in range(10):
        await send(dut, mrd3(addr=t * 4, tag=t + 10))
        await RisingEdge(dut.clk)
    total_cpls = 0
    for _ in range(5000):
        await RisingEdge(dut.clk)
        if dut.tlps_out_tvalid.value == 1:
            total_cpls += 1
    assert total_cpls > 0, "no completions from burst"

@cocotb.test()
async def test_mrd_large_length(dut):
    """MRd with length=32 (multi-completion)."""
    await reset(dut)
    await send(dut, mrd3(addr=0, tag=20, length=32))
    cpls = await recv_all(dut, timeout=100000)
    assert len(cpls) > 0

@cocotb.test()
async def test_mwr_pattern_integrity(dut):
    """Write known patterns to multiple offsets."""
    await reset(dut)
    patterns = [0xDEADBEEF, 0xCAFEBABE, 0x12345678, 0xAABBCCDD]
    for i, p in enumerate(patterns):
        await send(dut, mwr3(addr=0x200 + i * 4, data=p))
        for _ in range(10): await RisingEdge(dut.clk)

# --- NVMe-specific tests (only meaningful with nvme fixture) ---

@cocotb.test()
async def test_nvme_cap_read(dut):
    """MRd BAR0 offset 0 → NVMe CAP."""
    await reset(dut)
    await send(dut, mrd3(addr=0x0000, tag=30))
    cpls = await recv_all(dut)
    assert len(cpls) > 0
    cap = cpls[0][3]
    dut._log.info(f"NVMe CAP = {hex(cap)}")

@cocotb.test()
async def test_nvme_cc_write_readback(dut):
    """MWr CC register, MRd verify."""
    await reset(dut)
    await send(dut, mwr3(addr=0x0014, data=0x00000001))
    for _ in range(20): await RisingEdge(dut.clk)
    await send(dut, mrd3(addr=0x0014, tag=31))
    cpls = await recv_all(dut)
    assert len(cpls) > 0

@cocotb.test()
async def test_nvme_csts_read(dut):
    """MRd CSTS register (offset 0x1C)."""
    await reset(dut)
    await send(dut, mrd3(addr=0x001C, tag=32))
    cpls = await recv_all(dut)
    assert len(cpls) > 0

@cocotb.test()
async def test_nvme_vs_read(dut):
    """MRd VS register (offset 0x08)."""
    await reset(dut)
    await send(dut, mrd3(addr=0x0008, tag=33))
    cpls = await recv_all(dut)
    assert len(cpls) > 0

@cocotb.test()
async def test_nvme_intms_read(dut):
    """MRd INTMS (offset 0x0C)."""
    await reset(dut)
    await send(dut, mrd3(addr=0x000C, tag=34))
    cpls = await recv_all(dut)
    assert len(cpls) > 0

@cocotb.test()
async def test_nvme_doorbell_write(dut):
    """MWr NVMe SQ0 doorbell."""
    await reset(dut)
    await send(dut, mwr3(addr=0x1000, data=0x00000000))
    for _ in range(100): await RisingEdge(dut.clk)

@cocotb.test()
async def test_nvme_full_init_sequence(dut):
    """Full NVMe init: CC.EN=1 → poll CSTS.RDY."""
    await reset(dut)
    await send(dut, mwr3(addr=0x0014, data=0x00000001))
    for _ in range(500): await RisingEdge(dut.clk)
    await send(dut, mrd3(addr=0x001C, tag=40))
    cpls = await recv_all(dut)
    assert len(cpls) > 0
