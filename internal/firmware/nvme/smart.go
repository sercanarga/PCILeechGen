package nvme

import (
	"crypto/rand"
	"math/big"
)

// SMART holds plausible donor SMART/Health wear values that seed the admin
// responder's counters so a fresh clone reports realistic usage.
type SMART struct {
	DataUnitsRead     uint64
	DataUnitsWritten  uint64
	HostReadCommands  uint64
	HostWriteCommands uint64
	ControllerBusyMin uint64
	PowerCycles       uint32
	PowerOnHours      uint32
	UnsafeShutdowns   uint32
	MediaErrors       uint32
	ErrorLogEntries   uint32
}

// BuildSMART generates randomized SMART/Health wear values. Counters are
// correlated (data units and host commands scale with power-on hours) so the
// clone resembles a consistently worn drive, not independent noise.
func BuildSMART() *SMART {
	hours := randU32(500, 15000)
	cycles := randU32(100, 1000)

	// NVMe "data unit" = 1000 x 512B blocks.
	writeUnitsPerHour := randU64(50, 5000)
	readUnitsPerHour := randU64(100, 8000)
	dataUnitsWritten := uint64(hours) * writeUnitsPerHour
	dataUnitsRead := uint64(hours) * readUnitsPerHour

	writeCmds := dataUnitsWritten * randU64(5, 50)
	readCmds := dataUnitsRead * randU64(10, 100)

	busyMin := (uint64(hours) * 60 * uint64(randU32(1, 8))) / 100

	return &SMART{
		DataUnitsRead:     dataUnitsRead,
		DataUnitsWritten:  dataUnitsWritten,
		HostReadCommands:  readCmds,
		HostWriteCommands: writeCmds,
		ControllerBusyMin: busyMin,
		PowerCycles:       cycles,
		PowerOnHours:      hours,
		UnsafeShutdowns:   randU32(1, 30),
		MediaErrors:       0,
		ErrorLogEntries:   randU32(0, 8),
	}
}

func randU32(min, max uint32) uint32 {
	n, _ := rand.Int(rand.Reader, big.NewInt(int64(max-min)))
	return min + uint32(n.Int64())
}

func randU64(min, max uint64) uint64 {
	n, _ := rand.Int(rand.Reader, new(big.Int).SetUint64(max-min))
	return min + n.Uint64()
}
