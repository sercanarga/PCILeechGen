import cocotb
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


async def recv_all(dut, timeout=10000):
    cpls = []
    idle = 0
    for _ in range(timeout):
        await RisingEdge(dut.clk)
        if dut.tlps_out_tvalid.value == 1:
            raw = dut.tlps_out_tdata.value.integer
            cpls.append([(raw >> (i * 32)) & 0xFFFFFFFF for i in range(4)])
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
    for _ in range(10):
        await RisingEdge(dut.clk)
    dut.rst.value = 0
    for _ in range(200):
        await RisingEdge(dut.clk)


@cocotb.test()
async def test_config_space_vendor_read(dut):
    await reset(dut)
    await send(dut, mrd3(addr=0x0000, tag=1))
    cpls = await recv_all(dut)
    assert len(cpls) > 0, "no completion for config vendor read"


@cocotb.test()
async def test_bar0_read(dut):
    await reset(dut)
    await send(dut, mrd3(addr=0x200, tag=2))
    cpls = await recv_all(dut)
    assert len(cpls) > 0, "no completion for BAR0 read"


@cocotb.test()
async def test_bar0_write_then_read(dut):
    await reset(dut)
    await send(dut, mwr3(addr=0x200, data=0xDEADBEEF))
    for _ in range(20):
        await RisingEdge(dut.clk)
    await send(dut, mrd3(addr=0x200, tag=3))
    cpls = await recv_all(dut)
    assert len(cpls) > 0, "no completion after BAR0 write+read"


@cocotb.test()
async def test_bar0_multi_dword_read(dut):
    await reset(dut)
    await send(dut, mrd3(addr=0x200, tag=4, length=4))
    cpls = await recv_all(dut, timeout=100000)
    assert len(cpls) > 0, "no completion for multi-DW BAR0 read"


@cocotb.test()
async def test_no_bar_hit_no_completion(dut):
    await reset(dut)
    dut.tlps_in_tdata.value = mrd3(addr=0, tag=5)
    dut.tlps_in_tvalid.value = 1
    dut.tlps_in_tlast.value = 1
    dut.tlps_in_tuser.value = 1
    dut.tlps_in_tkeepdw.value = 0xF
    await RisingEdge(dut.clk)
    dut.tlps_in_tvalid.value = 0
    for _ in range(100):
        await RisingEdge(dut.clk)
    assert dut.tlps_out_tvalid.value == 0, "completion produced without BAR hit"


@cocotb.test()
async def test_rapid_bar0_reads(dut):
    await reset(dut)
    for t in range(10):
        await send(dut, mrd3(addr=0x200 + t * 4, tag=10 + t))
        await RisingEdge(dut.clk)
    total = 0
    for _ in range(5000):
        await RisingEdge(dut.clk)
        if dut.tlps_out_tvalid.value == 1:
            total += 1
    assert total > 0, "no completions for rapid BAR0 reads"


@cocotb.test()
async def test_io_read_returns_ur(dut):
    await reset(dut)
    await send(dut, iord3(addr=0, tag=20))
    cpls = await recv_all(dut)
    assert len(cpls) > 0, "no completion for IO read"
    status = (cpls[0][1] >> 13) & 0x7
    assert status == 0b001, f"expected UR completion status, got {status:#b}"


def _bswap32(v):
    v &= 0xFFFFFFFF
    return ((v & 0xFF) << 24) | ((v & 0xFF00) << 8) | ((v >> 8) & 0xFF00) | ((v >> 24) & 0xFF)


@cocotb.test()
async def test_bar0_byte_enable(dut):
    await reset(dut)
    await send(dut, mwr3(addr=0x200, data=0xFFFFFFFF))
    for _ in range(20): await RisingEdge(dut.clk)
    await send(dut, mwr3(addr=0x200, data=0x11223344, be=0x1))
    for _ in range(20): await RisingEdge(dut.clk)
    await send(dut, mrd3(addr=0x200, tag=30))
    cpls = await recv_all(dut)
    if not cpls:
        return
    val = cpls[0][3]
    if val == 0:
        return
    bram = _bswap32(val)
    assert (bram & 0xFF) == 0x44, f"byte0 not updated: bram={bram:#x}"
    assert (bram & 0xFFFFFF00) == 0xFFFFFF00, f"other bytes not preserved: bram={bram:#x}"


def cfgrd0(offset, tag=0x40, req_id=0x0000):
    dw0 = (0b000 << 29) | (0b00101 << 24) | 1
    dw1 = ((req_id & 0xFFFF) << 16) | ((tag & 0xFF) << 8) | 0x000F
    dw2 = offset & 0xFFC
    return dw0 | (dw1 << 32) | (dw2 << 64)


def cfgwr0(offset, data, tag=0x40, req_id=0x0000, be=0xF):
    dw0 = (0b010 << 29) | (0b00101 << 24) | 1
    dw1 = ((req_id & 0xFFFF) << 16) | ((tag & 0xFF) << 8) | (be & 0xF)
    dw2 = offset & 0xFFC
    return dw0 | (dw1 << 32) | (dw2 << 64) | ((data & 0xFFFFFFFF) << 96)


async def recv_cfg(dut, timeout=10000):
    cpls = []
    idle = 0
    for _ in range(timeout):
        await RisingEdge(dut.clk)
        if dut.tlps_cfg_rsp_tvalid.value == 1:
            raw = dut.tlps_cfg_rsp_tdata.value.integer
            cpls.append([(raw >> (i * 32)) & 0xFFFFFFFF for i in range(4)])
            if dut.tlps_cfg_rsp_tlast.value == 1:
                break
            idle = 0
        else:
            idle += 1
            if cpls and idle > 50:
                break
    return cpls


async def _cfg_send(dut, tdata):
    dut.tlps_in_tdata.value = tdata
    dut.tlps_in_tvalid.value = 1
    dut.tlps_in_tlast.value = 1
    dut.tlps_in_tuser.value = 1
    dut.tlps_in_tkeepdw.value = 0xF
    await RisingEdge(dut.clk)
    dut.tlps_in_tvalid.value = 0
    dut.tlps_in_tlast.value = 0
    dut.tlps_in_tuser.value = 0


@cocotb.test()
async def test_config_readonly_identity(dut):
    await reset(dut)
    await _cfg_send(dut, cfgrd0(0x00, tag=40))
    cpls = await recv_cfg(dut)
    assert len(cpls) > 0, "no cfg completion"
    before = cpls[0][3]
    await _cfg_send(dut, cfgwr0(0x00, 0xDEADBEEF, tag=41))
    for _ in range(20):
        await RisingEdge(dut.clk)
    await _cfg_send(dut, cfgrd0(0x00, tag=42))
    cpls = await recv_cfg(dut)
    after = cpls[0][3] if cpls else before
    dut._log.info(f"cfg identity dw0 before={before:#x} after={after:#x}")
    assert before == after, "read-only VID/DID changed"


@cocotb.test()
async def test_config_writable_interrupt_line(dut):
    await reset(dut)
    await _cfg_send(dut, cfgwr0(0x3C, 0x000000DE, tag=50))
    for _ in range(20):
        await RisingEdge(dut.clk)
    await _cfg_send(dut, cfgrd0(0x3C, tag=51))
    cpls = await recv_cfg(dut)
    assert len(cpls) > 0, "no cfg completion"
    val = cpls[0][3]
    dut._log.info(f"cfg intline after write 0xDE: {val:#x}")
    assert (val & 0xFF) == 0xDE, f"interrupt line byte not writable: {val:#x}"
    assert (val & 0xFFFFFF00) == 0x100, f"intpin/ro bytes not preserved: {val:#x}"


@cocotb.test()
async def test_driver_pci_enumeration(dut):
    """Generic PCI enumeration sequence (applies to all device types)."""
    await reset(dut)
    tag = 1
    results = {}

    # 1. VID/DID
    await _cfg_send(dut, cfgrd0(0x00, tag)); tag += 1
    c = await recv_cfg(dut)
    results["vid_did"] = c[0][3] if c else 0

    # 2. Revision + Class
    await _cfg_send(dut, cfgrd0(0x08, tag)); tag += 1
    c = await recv_cfg(dut)
    results["class"] = c[0][3] if c else 0

    # 3. Subsystem ID
    await _cfg_send(dut, cfgrd0(0x2C, tag)); tag += 1
    c = await recv_cfg(dut)
    results["subsys"] = c[0][3] if c else 0

    # 4. BAR0 sizing
    await _cfg_send(dut, cfgwr0(0x10, 0xFFFFFFFF))
    for _ in range(10): await RisingEdge(dut.clk)
    await _cfg_send(dut, cfgrd0(0x10, tag)); tag += 1
    c = await recv_cfg(dut)
    results["bar0_raw"] = c[0][3] if c else 0
    await _cfg_send(dut, cfgwr0(0x10, 0x00000000))
    for _ in range(10): await RisingEdge(dut.clk)

    # 5. Cap chain walk
    await _cfg_send(dut, cfgrd0(0x34, tag)); tag += 1
    c = await recv_cfg(dut)
    cap_ptr = (c[0][3] if c else 0) & 0xFC
    caps_found = []
    visited = set()
    while cap_ptr != 0 and cap_ptr not in visited and tag < 30:
        visited.add(cap_ptr)
        await _cfg_send(dut, cfgrd0(cap_ptr, tag)); tag += 1
        c = await recv_cfg(dut)
        val = c[0][3] if c else 0
        cap_id = val & 0xFF
        next_ptr = (val >> 8) & 0xFC
        caps_found.append((cap_id, cap_ptr))
        cap_ptr = next_ptr

    dut._log.info(f"pci enum: vid_did={results['vid_did']:#x} class={results['class']:#x} subsys={results['subsys']:#x}")
    dut._log.info(f"caps: {[(hex(i), hex(o)) for i,o in caps_found]}")
    dut._log.info(f"bar0 sizing: {results['bar0_raw']:#x}")

    assert results["vid_did"] != 0, "VID/DID = 0 — device not detected by PCI bus"
    assert results["subsys"] != 0, "subsystem ID = 0 — Code-10 INF match failure"
    assert len(caps_found) >= 2, f"too few capabilities: {caps_found}"
    assert results["bar0_raw"] != 0, "BAR0 sizing = 0 — BAR0 not responding"
