import struct
import cocotb
from cocotb.clock import Clock
from cocotb.triggers import RisingEdge, Timer

def mrd3(addr, tag=1, length=1, first_be=0xF, last_be=0x0, req_id=0):
    dw0 = (0b000 << 29) | (0b00000 << 24) | (length & 0x3FF)
    dw1 = ((req_id & 0xFFFF) << 16) | ((tag & 0xFF) << 8) | ((last_be & 0xF) << 4) | (first_be & 0xF)
    dw2 = (addr & 0xFFFFFFFC)
    return dw0 | (dw1 << 32) | (dw2 << 64)

def mwr3(addr, data, tag=0, req_id=0):
    dw0 = (0b010 << 29) | (0b00000 << 24) | 1
    dw1 = ((req_id & 0xFFFF) << 16) | ((tag & 0xFF) << 8) | 0xF
    dw2 = (addr & 0xFFFFFFFC)
    dw3 = data
    return dw0 | (dw1 << 32) | (dw2 << 64) | (dw3 << 96)

async def send(dut, tdata, bar=0):
    tuser = 1 | (1 << (bar + 2))
    dut.tlps_in_tdata.value = tdata
    dut.tlps_in_tvalid.value = 1
    dut.tlps_in_tlast.value = 1
    dut.tlps_in_tuser.value = tuser
    dut.tlps_in_tkeepdw.value = 0xF
    await RisingEdge(dut.clk)
    dut.tlps_in_tvalid.value = 0
    dut.tlps_in_tlast.value = 0
    dut.tlps_in_tuser.value = 0

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
async def test_mrd_debug(dut):
    """MRd to BAR0 offset 0, monitor all output streams."""
    await reset(dut)

    # Monitor cycle-by-cycle for 200 cycles after send
    dut._log.info("Sending MRd to BAR0 offset 0")
    await send(dut, mrd3(addr=0x0000, tag=1))

    out_hits = 0
    dma_hits = 0
    for i in range(500):
        await RisingEdge(dut.clk)
        if dut.tlps_out_tvalid.value == 1:
            out_hits += 1
            raw = dut.tlps_out_tdata.value.integer
            dws = [(raw >> (j*32)) & 0xFFFFFFFF for j in range(4)]
            if out_hits <= 5:
                dut._log.info(f"  cycle {i}: tlps_out DWs={[hex(d) for d in dws]}")
        if dut.tlps_dma_out_tvalid.value == 1:
            dma_hits += 1
            if dma_hits <= 3:
                dut._log.info(f"  cycle {i}: tlps_dma_out active")

    dut._log.info(f"Results: tlps_out hits={out_hits}, tlps_dma_out hits={dma_hits}")
    if out_hits > 0:
        dut._log.info("PASS: completion received on tlps_out")
    elif dma_hits > 0:
        dut._log.warning("Completion went to tlps_dma_out, not tlps_out")
    else:
        dut._log.error("No response on any output stream")

@cocotb.test()
async def test_mwr_doorbell(dut):
    """MWr BAR0 doorbell, no crash."""
    await reset(dut)
    await send(dut, mwr3(addr=0x1000, data=0x00000000))
    for _ in range(50):
        await RisingEdge(dut.clk)
    dut._log.info("PASS: no crash")
