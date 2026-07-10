package services

type PowerState uint8

const (
	D0 PowerState = iota
	D1
	D2
	D3Hot
)

const D3 = D3Hot

type Inputs struct {
	Reset    bool
	FLR      bool
	DState   PowerState
	MSE      bool
	BME      bool
	Turnoff  bool
	LinkDown bool
}

type Outputs struct {
	DeviceReset bool
	IOEnabled   bool
	DMAEnabled  bool
	Quiesce     bool
	Generation  uint32
}

type Lifecycle struct {
	generation uint32
	inReset    bool
}

func (l *Lifecycle) Apply(in Inputs) Outputs {
	deviceReset := in.Reset || in.FLR
	if deviceReset && !l.inReset {
		l.generation++
	}
	l.inReset = deviceReset

	active := !deviceReset && in.DState == D0 && !in.Turnoff && !in.LinkDown
	ioEnabled := active && in.MSE
	dmaEnabled := active && in.BME
	return Outputs{
		DeviceReset: deviceReset,
		IOEnabled:   ioEnabled,
		DMAEnabled:  dmaEnabled,
		Quiesce:     !dmaEnabled,
		Generation:  l.generation,
	}
}
