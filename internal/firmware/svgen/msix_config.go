package svgen

// MSIXConfig parameters for MSI-X SV template.
type MSIXConfig struct {
	NumVectors  int
	TableBIR    int
	TableOffset uint32
	PBABIR      int
	PBAOffset   uint32
}

func (c *MSIXConfig) TableEnd() uint64 {
	return uint64(c.TableOffset) + uint64(c.NumVectors)*16
}

func (c *MSIXConfig) PBAEnd() uint64 {
	return uint64(c.PBAOffset) + uint64((c.NumVectors+63)/64)*8
}
