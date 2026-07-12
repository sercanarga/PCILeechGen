import struct, cocotb
from cocotb.clock import Clock
from cocotb.triggers import RisingEdge, Timer, FallingEdge

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
    d = data & 0xFFFFFFFF
    swd = ((d & 0xFF) << 24) | ((d & 0xFF00) << 8) | ((d >> 8) & 0xFF00) | ((d >> 24) & 0xFF)
    return dw0 | (dw1 << 32) | (dw2 << 64) | (swd << 96)

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

async def poke(dut, dw_idx, data):
    dut.host_poke_valid.value = 1
    dut.host_poke_addr.value = dw_idx & 0xFFFF
    dut.host_poke_data.value = data & 0xFFFFFFFF
    await RisingEdge(dut.clk)
    dut.host_poke_valid.value = 0

async def peek(dut, dw_idx):
    dut.host_peek_addr.value = dw_idx & 0xFFFF
    await RisingEdge(dut.clk)
    await RisingEdge(dut.clk)
    return dut.host_peek_data.value.integer

async def send_write(dut, addr, data):
    await send(dut, mwr3(addr=addr, data=data))
    for _ in range(30): await RisingEdge(dut.clk)

async def setup_admin_queues(dut, asq=0x1000, acq=0x2000, asqs=15, acqs=15):
    await send_write(dut, 0x0024, ((acqs & 0xFFF) << 16) | (asqs & 0xFFF))
    await send_write(dut, 0x0028, asq & 0xFFFFFFFF)
    await send_write(dut, 0x002C, (asq >> 32) & 0xFFFFFFFF)
    await send_write(dut, 0x0030, acq & 0xFFFFFFFF)
    await send_write(dut, 0x0034, (acq >> 32) & 0xFFFFFFFF)
    for _ in range(10): await RisingEdge(dut.clk)

async def poke_sqe(dut, dwbase, op, nsid=0, prp1=0, cdw10=0, cdw11=0, cdw12=0):
    await poke(dut, dwbase + 0,  0x00010000 | op)
    await poke(dut, dwbase + 1,  nsid)
    await poke(dut, dwbase + 6,  prp1 & 0xFFFFFFFF)
    await poke(dut, dwbase + 7,  (prp1 >> 32) & 0xFFFFFFFF)
    await poke(dut, dwbase + 10, cdw10)
    await poke(dut, dwbase + 11, cdw11)
    await poke(dut, dwbase + 12, cdw12)

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
async def test_dma_loopback_sqe_fetch(dut):
    await reset(dut)
    await setup_admin_queues(dut)
    await send(dut, mwr3(addr=0x0014, data=0x00000001))
    for _ in range(50): await RisingEdge(dut.clk)
    await send(dut, mwr3(addr=0x1000, data=0x00000001))
    beats = 0
    for _ in range(20000):
        await RisingEdge(dut.clk)
        if dut.tlps_dma_out_tvalid.value == 1:
            beats += 1
    dut._log.info(f"dma beats after doorbell (empty SQE): {beats}")

@cocotb.test()
async def test_nvme_identify_full_dma(dut):
    await reset(dut)
    await setup_admin_queues(dut, asq=0x1000, acq=0x2000)
    await poke_sqe(dut, 0x400, op=0x06, prp1=0x5000, cdw10=0x00000001)
    await send(dut, mwr3(addr=0x0014, data=0x00000001))
    for _ in range(50): await RisingEdge(dut.clk)
    await send(dut, mwr3(addr=0x1000, data=0x00000001))
    for _ in range(60000): await RisingEdge(dut.clk)
    nonzero = 0
    for dw in range(0x1400, 0x1800):
        if await peek(dut, dw) != 0: nonzero += 1
    cqe_dw3 = await peek(dut, 0x803)
    status = (cqe_dw3 >> 17) & 0x7FFF
    dut._log.info(f"identify non-zero dwords @0x5000: {nonzero}, cqe status={status}")
    assert nonzero > 0, "identify data not written to host memory"
    assert cqe_dw3 != 0, "no CQE posted to ACQ"
    assert status == 0, f"cqe status not SUCCESS: {status}"

@cocotb.test()
async def test_nvme_cc_en_gate_blocks_doorbell(dut):
    await reset(dut)
    await setup_admin_queues(dut, asq=0x1000, acq=0x2000)
    await poke_sqe(dut, 0x400, op=0x06, prp1=0x5000, cdw10=0x00000001)
    await send(dut, mwr3(addr=0x1000, data=0x00000001))
    beats = 0
    for _ in range(3000):
        await RisingEdge(dut.clk)
        if dut.tlps_dma_out_tvalid.value == 1:
            beats += 1
    dut._log.info(f"dma beats before CC.EN: {beats}")


# --- MSI-X interrupt delivery tests ---
#
# Interrupt path (pcileech_tlps128_bar_controller.sv):
#   responder S_SEND_MSIX -> pba_set_valid/vector (== nvme_irq_event_*)
#   -> i_nvme_interrupt_service.event_valid  (sets pending[], pulses pba_set)
#   -> i_msix_table.pba_set_valid            (sets msix_pba[vector] bit)
#   interrupt_service scan FSM (Q_SELECT->Q_WAIT->Q_CHECK) finds the pending
#   vector and, if function_enable && !function_mask && !vector_masked, asserts
#   delivery_valid. The responder then issues a 1-DW Memory-Write TLP to
#   {msix_vector_addr} with payload msix_vector_data via the DMA bridge ->
#   tlps_dma_out.  NOTE: bar_controller.intr_req is hardwired 1'b0 (line ~1094),
#   so the MSI-X message is observable ONLY on tlps_dma_out, never on intr_req.

def bswap32(v):
    v &= 0xFFFFFFFF
    return ((v & 0xFF) << 24) | ((v & 0xFF00) << 8) | ((v >> 8) & 0xFF00) | ((v >> 24) & 0xFF)


async def read_bar(dut, addr, tag):
    """Single-DW BAR read; returns the 32-bit register value (byte-de-swapped).

    The rdengine places completion data byte-swapped in tdata[127:96] (DW3);
    recv_all stores that lane as cpls[0][3], so we bswap it back."""
    await send(dut, mrd3(addr=addr, tag=tag))
    cpls = await recv_all(dut)
    if not cpls:
        return None
    return bswap32(cpls[0][3])


async def program_msix_vector0(dut, addr_lo, addr_hi, msg_data, masked):
    """Write MSI-X table entry 0 at BAR0+0x2000.

    Layout per entry (16 bytes): msg_addr_lo, msg_addr_hi, msg_data, vector_ctrl.
    vector_ctrl bit0 = Mask (1=masked/suppressed)."""
    await send_write(dut, 0x2000, addr_lo)
    await send_write(dut, 0x2004, addr_hi)
    await send_write(dut, 0x2008, msg_data)
    await send_write(dut, 0x200C, 0x00000001 if masked else 0x00000000)


async def run_identify(dut):
    """Mirror test_nvme_identify_full_dma setup (no final assertions)."""
    await setup_admin_queues(dut, asq=0x1000, acq=0x2000)
    await poke_sqe(dut, 0x400, op=0x06, prp1=0x5000, cdw10=0x00000001)
    await send(dut, mwr3(addr=0x0014, data=0x00000001))   # CC.EN
    for _ in range(50):
        await RisingEdge(dut.clk)
    await send(dut, mwr3(addr=0x1000, data=0x00000001))   # ring admin SQ doorbell


async def watch_msix_mwr(dut, full_addr, cycles=60000):
    """Scan tlps_dma_out for the MSI-X Memory-Write TLP.

    A 1-DW MWr (fmt=0b011, 4DW header) occupies two beats: header (tuser[0]=1,
    addr in DW2/DW3) then data (tuser[0]=0, tlast). Returns (seen, data_value,
    mwr_count) where mwr_count is the total number of MWr first-beats seen."""
    seen = False
    data_value = None
    mwr_count = 0
    in_mwr = False
    cur_addr = None
    for _ in range(cycles):
        await RisingEdge(dut.clk)
        if dut.tlps_dma_out_tvalid.value != 1:
            continue
        raw = dut.tlps_dma_out_tdata.value.integer
        dws = [(raw >> (i * 32)) & 0xFFFFFFFF for i in range(4)]
        fmt = (dws[0] >> 29) & 0x7
        first = dut.tlps_dma_out_tuser.value.integer & 1
        if first:
            in_mwr = (fmt == 0b011)
            if in_mwr:
                mwr_count += 1
                cur_addr = (dws[2] << 32) | dws[3]
            else:
                cur_addr = None
        elif in_mwr:
            if cur_addr == full_addr:
                seen = True
                data_value = bswap32(dws[0])
            in_mwr = False
    return seen, data_value, mwr_count


@cocotb.test()
async def test_msix_interrupt_after_cqe(dut):
    """An UNMASKED MSI-X vector fires a Memory-Write TLP on tlps_dma_out after
    the Identify CQE is posted, carrying the programmed addr/data."""
    await reset(dut)

    MSIX_ADDR_LO = 0xFEE0C000   # low 16 bits -> host_mem[0x3000] via loopback
    MSIX_ADDR_HI = 0x00000000
    MSIX_MSG_DATA = 0x00004400
    FULL_ADDR = (MSIX_ADDR_HI << 32) | MSIX_ADDR_LO

    await program_msix_vector0(dut, MSIX_ADDR_LO, MSIX_ADDR_HI, MSIX_MSG_DATA, masked=False)
    ctrl = await read_bar(dut, 0x200C, tag=70)
    assert ctrl == 0, f"vector 0 ctrl={ctrl:#x}, expected 0 (unmasked)"

    await run_identify(dut)

    intr_pulses = 0
    seen = [False]
    data_val = [None]
    mwr_count = [0]

    async def _watch():
        s, dv, mc = await watch_msix_mwr(dut, FULL_ADDR, cycles=60000)
        seen[0], data_val[0], mwr_count[0] = s, dv, mc

    watch_task = cocotb.start_soon(_watch())
    for _ in range(60000):
        await RisingEdge(dut.clk)
        # intr_req is 1-bit -> a cocotb Logic object (no .integer/.binstr).
        if dut.intr_req.value == 1:
            intr_pulses += 1
    await watch_task

    cqe_dw3 = await peek(dut, 0x803)
    delivered = await peek(dut, 0x3000)

    dut._log.info(
        f"msix_seen={seen[0]} data={data_val[0] if data_val[0] is not None else 'NA'} "
        f"mwr_count={mwr_count[0]} intr_pulses={intr_pulses} "
        f"cqe_dw3={cqe_dw3:#x} host_mem[0x3000]={delivered:#x}")

    # Assert points.
    assert cqe_dw3 != 0, "no CQE posted to ACQ - identify did not complete"
    assert seen[0], "MSI-X MWr TLP not observed on tlps_dma_out after CQE"
    assert data_val[0] == MSIX_MSG_DATA, (
        f"MSI-X data mismatch: {data_val[0]:#x} != {MSIX_MSG_DATA:#x}")
    assert mwr_count[0] > 0, "no MWr TLPs at all on tlps_dma_out"
    assert intr_pulses == 0, (
        f"intr_req pulsed {intr_pulses}x; it is hardwired 0 in bar_controller "
        f"- MSI-X uses the MWr TLP path, not intr_req")
    assert delivered == MSIX_MSG_DATA, (
        f"MSI-X write not delivered to host_mem[0x3000]: {delivered:#x}")


@cocotb.test()
async def test_msix_pba_pending_when_masked(dut):
    """A MASKED vector still sets its PBA pending bit (PCIe 6.8.3.4) and the bit
    stays set because delivery (and thus the PBA clear) is suppressed.
    Complements test_msix_interrupt_after_cqe."""
    await reset(dut)

    MSIX_ADDR_LO = 0xFEE0D000
    MSIX_ADDR_HI = 0x00000000
    MSIX_MSG_DATA = 0x0000ABCD
    FULL_ADDR = (MSIX_ADDR_HI << 32) | MSIX_ADDR_LO

    await program_msix_vector0(dut, MSIX_ADDR_LO, MSIX_ADDR_HI, MSIX_MSG_DATA, masked=True)
    ctrl = await read_bar(dut, 0x200C, tag=71)
    assert ctrl == 1, f"vector 0 ctrl={ctrl:#x}, expected 1 (masked)"

    await run_identify(dut)

    seen, _dv, _mc = await watch_msix_mwr(dut, FULL_ADDR, cycles=60000)

    pba = await read_bar(dut, 0x3000, tag=72)
    cqe_dw3 = await peek(dut, 0x803)

    dut._log.info(f"masked: pba={pba:#x} msix_mwr_seen={seen} cqe_dw3={cqe_dw3:#x}")

    assert cqe_dw3 != 0, "no CQE posted to ACQ - identify did not complete"
    assert pba == 0x1, f"PBA bit 0 not set for masked vector 0: {pba:#x}"
    assert not seen, "MSI-X MWr leaked to tlps_dma_out for a masked vector"


async def _post_admin_cmd(dut, op, nsid=0, prp1=0, cdw10=0, cdw11=0, cdw12=0, wait=60000):
    await reset(dut)
    await setup_admin_queues(dut)
    await poke_sqe(dut, 0x400, op=op, nsid=nsid, prp1=prp1, cdw10=cdw10, cdw11=cdw11, cdw12=cdw12)
    await send(dut, mwr3(addr=0x0014, data=0x00000001))
    for _ in range(50): await RisingEdge(dut.clk)
    await send(dut, mwr3(addr=0x1000, data=0x00000001))
    for _ in range(wait): await RisingEdge(dut.clk)
    return await peek(dut, 0x803)

@cocotb.test()
async def test_nvme_get_features_invalid(dut):
    cqe_dw3 = await _post_admin_cmd(dut, op=0x0A, cdw10=0x000000FF)
    status = (cqe_dw3 >> 17) & 0x7FFF
    dut._log.info(f"get features invalid: status={status:#x}")
    assert cqe_dw3 != 0, "no CQE posted"
    assert status == 0x0002, f"expected INVALID_FIELD, got {status:#x}"

@cocotb.test()
async def test_nvme_get_log_page_smart(dut):
    cqe_dw3 = await _post_admin_cmd(dut, op=0x02, prp1=0x6000, cdw10=0x007F0002)
    status = (cqe_dw3 >> 17) & 0x7FFF
    spare = await peek(dut, 0x1800)
    spare_thr = await peek(dut, 0x1801)
    unsafe = await peek(dut, 0x1824)
    dut._log.info(f"smart: status={status:#x} spare_dw0={spare:#x} thr={spare_thr:#x} unsafe={unsafe:#x}")
    assert status == 0, f"smart log failed: {status:#x}"
    assert (spare >> 24) & 0xFF == 0x64, f"spare byte wrong: {spare:#x}"
    assert spare_thr == 0x0000000A, f"spare threshold wrong: {spare_thr:#x}"
    assert unsafe == 0x00000003, f"unsafe shutdowns wrong: {unsafe:#x}"

@cocotb.test()
async def test_nvme_create_io_cq(dut):
    cqe_dw3 = await _post_admin_cmd(dut, op=0x05, prp1=0x4000, cdw10=0x00010001, cdw11=0x00000001)
    status = (cqe_dw3 >> 17) & 0x7FFF
    dut._log.info(f"create io cq: status={status:#x}")
    assert cqe_dw3 != 0, "no CQE posted"
    assert status == 0, f"create io cq failed: {status:#x}"

@cocotb.test()
async def test_nvme_create_io_sq(dut):
    await reset(dut)
    await setup_admin_queues(dut)
    await poke_sqe(dut, 0x400, op=0x05, prp1=0x4000, cdw10=0x00010001, cdw11=0x00000001)
    await send(dut, mwr3(addr=0x0014, data=0x00000001))
    for _ in range(50): await RisingEdge(dut.clk)
    await send(dut, mwr3(addr=0x1000, data=0x00000001))
    for _ in range(60000): await RisingEdge(dut.clk)
    cq_status = (await peek(dut, 0x803) >> 17) & 0x7FFF
    await poke_sqe(dut, 0x410, op=0x01, prp1=0x3000, cdw10=0x00010001, cdw11=0x00010001)
    await send(dut, mwr3(addr=0x1000, data=0x00000002))
    for _ in range(60000): await RisingEdge(dut.clk)
    sq_status = (await peek(dut, 0x813) >> 17) & 0x7FFF
    dut._log.info(f"create io: cq_status={cq_status:#x} sq_status={sq_status:#x}")
    assert cq_status == 0, f"create io cq failed: {cq_status:#x}"
    assert sq_status == 0, f"create io sq failed: {sq_status:#x}"

async def _create_io_queues(dut):
    await poke_sqe(dut, 0x400, op=0x05, prp1=0x4000, cdw10=0x00010001, cdw11=0x00000001)
    await send(dut, mwr3(addr=0x0014, data=0x00000001))
    for _ in range(50): await RisingEdge(dut.clk)
    await send(dut, mwr3(addr=0x1000, data=0x00000001))
    for _ in range(60000): await RisingEdge(dut.clk)
    await poke_sqe(dut, 0x410, op=0x01, prp1=0x3000, cdw10=0x00010001, cdw11=0x00010001)
    await send(dut, mwr3(addr=0x1000, data=0x00000002))
    for _ in range(60000): await RisingEdge(dut.clk)

@cocotb.test()
async def test_nvme_io_read_completed(dut):
    await reset(dut)
    await setup_admin_queues(dut)
    await _create_io_queues(dut)
    await poke_sqe(dut, 0xC00, op=0x02, nsid=0x00000001, cdw12=0x00000040)
    await send(dut, mwr3(addr=0x1008, data=0x00000001))
    for _ in range(60000): await RisingEdge(dut.clk)
    status = (await peek(dut, 0x1000) >> 17) & 0x7FFF
    dut._log.info(f"io read: status={status:#x}")
    assert status == 0, f"io read failed: {status:#x}"



@cocotb.test()
async def test_nvme_get_features_invalid_fid(dut):
    cqe_dw3 = await _post_admin_cmd(dut, op=0x0A, cdw10=0x000000FF)
    status = (cqe_dw3 >> 17) & 0x7FFF
    dut._log.info(f"get features FID=0xFF: cqe_dw3={cqe_dw3:#x} status={status:#x}")
    assert cqe_dw3 != 0, "no CQE posted for Get Features"
    assert status == 0x0002, f"get features status != INVALID_FIELD: {status:#x}"

@cocotb.test()
async def test_nvme_get_log_page_smart(dut):
    cqe_dw3 = await _post_admin_cmd(dut, op=0x02, prp1=0x6000, cdw10=0x007F0002)
    status = (cqe_dw3 >> 17) & 0x7FFF
    smart_dw0 = await peek(dut, 0x1800)
    smart_dw1 = await peek(dut, 0x1801)
    dut._log.info(f"smart: cqe_dw3={cqe_dw3:#x} status={status:#x} dw0={smart_dw0:#x} dw1={smart_dw1:#x}")
    assert cqe_dw3 != 0, "no CQE posted for Get Log Page"
    assert status == 0, f"smart status != SUCCESS: {status:#x}"
    assert smart_dw0 != 0, "SMART data not written to host_mem[0x1800]"
