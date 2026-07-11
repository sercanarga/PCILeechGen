package devicemodel

import "time"

const CurrentSchemaVersion = 1

const ConfigBIR = -1

type AddressSpace string

const (
	SpaceConfig AddressSpace = "config"
	SpaceBAR    AddressSpace = "bar"
)

type BARType string

const (
	BARTypeIO       BARType = "io"
	BARTypeMem32    BARType = "mem32"
	BARTypeMem64    BARType = "mem64"
	BARTypeDisabled BARType = "disabled"
)

type AccessPolicy string

const (
	AccessRO       AccessPolicy = "ro"
	AccessRW       AccessPolicy = "rw"
	AccessRW1C     AccessPolicy = "rw1c"
	AccessW1C      AccessPolicy = AccessRW1C
	AccessW1S      AccessPolicy = "w1s"
	AccessW0C      AccessPolicy = "w0c"
	AccessRC       AccessPolicy = "rc"
	AccessReserved AccessPolicy = "reserved"
)

type ResetDomain string

const (
	ResetPowerOn     ResetDomain = "power_on"
	ResetFundamental ResetDomain = "fundamental"
	ResetFunction    ResetDomain = "function"
	ResetSoftware    ResetDomain = "software"
)

type ConfidenceLevel string

const (
	ConfidenceUnknown   ConfidenceLevel = "unknown"
	ConfidenceInferred  ConfidenceLevel = "inferred"
	ConfidenceMeasured  ConfidenceLevel = "measured"
	ConfidenceSpecified ConfidenceLevel = "specified"
)

type Model struct {
	SchemaVersion       int                   `json:"schema_version"`
	Name                string                `json:"name"`
	Functions           []Function            `json:"functions"`
	ConfigSpace         ConfigSpace           `json:"config_space"`
	Capabilities        []Capability          `json:"capabilities"`
	BARs                []BAR                 `json:"bars"`
	Registers           []Register            `json:"registers"`
	Interrupts          []InterruptDescriptor `json:"interrupts"`
	MSIX                *MSIXDescriptor       `json:"msix,omitempty"`
	Transformations     []Transformation      `json:"transformations"`
	UnsupportedFeatures []UnsupportedFeature  `json:"unsupported_features"`
	Confidence          Confidence            `json:"confidence"`
	Provenance          Provenance            `json:"provenance"`
}

type DeviceModel = Model

type Function struct {
	BDF               string `json:"bdf"`
	VendorID          uint16 `json:"vendor_id"`
	DeviceID          uint16 `json:"device_id"`
	SubsystemVendorID uint16 `json:"subsystem_vendor_id"`
	SubsystemDeviceID uint16 `json:"subsystem_device_id"`
	RevisionID        uint8  `json:"revision_id"`
	ClassCode         uint32 `json:"class_code"`
	HeaderType        uint8  `json:"header_type"`
}

type ConfigSpace struct {
	Size       uint32        `json:"size"`
	ResetImage []byte        `json:"reset_image"`
	Fields     []ConfigField `json:"fields"`
}

type ConfigField struct {
	Name       string       `json:"name"`
	Offset     uint16       `json:"offset"`
	Width      uint8        `json:"width"`
	Mask       uint64       `json:"mask"`
	Access     AccessPolicy `json:"access"`
	ResetValue uint64       `json:"reset_value"`
}

type Capability struct {
	ID         uint16 `json:"id"`
	Name       string `json:"name"`
	Version    uint8  `json:"version,omitempty"`
	Offset     uint16 `json:"offset"`
	NextOffset uint16 `json:"next_offset"`
	Length     uint16 `json:"length"`
	Extended   bool   `json:"extended"`
	Data       []byte `json:"data"`
}

type BAR struct {
	BIR          int     `json:"bir"`
	Type         BARType `json:"type"`
	Size         uint64  `json:"size"`
	SizeKnown    bool    `json:"size_known"`
	Prefetchable bool    `json:"prefetchable"`
	AddressWidth uint8   `json:"address_width"`
	PairBIR      *int    `json:"pair_bir,omitempty"`
	ResetImage   []byte  `json:"reset_image,omitempty"`
}

type Register struct {
	Name        string          `json:"name"`
	Space       AddressSpace    `json:"space"`
	BIR         int             `json:"bir"`
	Offset      uint64          `json:"offset"`
	Width       uint8           `json:"width"`
	ResetDomain ResetDomain     `json:"reset_domain"`
	ResetValue  uint64          `json:"reset_value"`
	Fields      []RegisterField `json:"fields"`
	Confidence  ConfidenceLevel `json:"confidence"`
}

type RegisterField struct {
	Name       string       `json:"name"`
	Mask       uint64       `json:"mask"`
	Access     AccessPolicy `json:"access"`
	ResetValue uint64       `json:"reset_value"`
}

type InterruptDescriptor struct {
	Kind             string `json:"kind"`
	CapabilityOffset uint16 `json:"capability_offset,omitempty"`
	Vectors          uint16 `json:"vectors"`
	Pin              uint8  `json:"pin,omitempty"`
	BIR              int    `json:"bir,omitempty"`
	TableOffset      uint64 `json:"table_offset,omitempty"`
	PBAOffset        uint64 `json:"pba_offset,omitempty"`
}

type MSIXDescriptor struct {
	CapabilityOffset uint16 `json:"capability_offset"`
	TableSize        uint16 `json:"table_size"`
	TableBIR         int    `json:"table_bir"`
	TableOffset      uint64 `json:"table_offset"`
	PBABIR           int    `json:"pba_bir"`
	PBAOffset        uint64 `json:"pba_offset"`
}

type Transformation struct {
	Kind        string `json:"kind"`
	Target      string `json:"target"`
	Description string `json:"description"`
}

type UnsupportedFeature struct {
	Name   string `json:"name"`
	Reason string `json:"reason"`
	Source string `json:"source"`
}

type Confidence struct {
	Overall  ConfidenceLevel `json:"overall"`
	Evidence []string        `json:"evidence"`
}

type Provenance struct {
	Source      string    `json:"source"`
	ToolVersion string    `json:"tool_version"`
	CollectedAt time.Time `json:"collected_at"`
	DonorBDF    string    `json:"donor_bdf"`
	Host        string    `json:"host,omitempty"`
}
