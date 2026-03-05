package svgen

// MSIXConfig parameters for MSI-X SV template.
type MSIXConfig struct {
	NumVectors  int
	TableOffset uint32
	PBAOffset   uint32
}
