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


def bswap32(v):
    v &= 0xFFFFFFFF
    return ((v & 0xFF) << 24) | ((v & 0xFF00) << 8) | ((v >> 8) & 0xFF00) | ((v >> 24) & 0xFF)


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
    for _ in range(10): await RisingEdge(dut.clk)
    dut.rst.value = 0
    for _ in range(200): await RisingEdge(dut.clk)


async def read_bar(dut, addr, tag, bar=0):
    await send(dut, mrd3(addr=addr, tag=tag), bar=bar)
    cpls = await recv_all(dut)
    if not cpls:
        return None
    return bswap32(cpls[0][3])


async def write_bar(dut, addr, data, bar=0):
    await send(dut, mwr3(addr=addr, data=data), bar=bar)
    for _ in range(20): await RisingEdge(dut.clk)
