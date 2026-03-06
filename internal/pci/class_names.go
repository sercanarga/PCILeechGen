package pci

// pciSubClassNames maps (base_class << 8 | sub_class) to human-readable names.
var pciSubClassNames = map[uint16]string{
	// Mass Storage
	0x0101: "IDE interface",
	0x0104: "RAID bus controller",
	0x0106: "SATA controller",
	0x0107: "Serial Attached SCSI controller",
	0x0108: "Non-Volatile memory controller",
	// Network
	0x0200: "Ethernet controller",
	0x0280: "Network controller",
	// Display
	0x0300: "VGA compatible controller",
	0x0302: "3D controller",
	// Multimedia
	0x0400: "Multimedia video controller",
	0x0401: "Multimedia audio controller",
	0x0403: "Audio device",
	// Memory
	0x0500: "RAM memory",
	0x0580: "Memory controller",
	// Bridge
	0x0600: "Host bridge",
	0x0601: "ISA bridge",
	0x0604: "PCI bridge",
	0x0680: "Bridge",
	// Communication
	0x0700: "Serial controller",
	0x0780: "Communication controller",
	// System Peripheral
	0x0800: "PIC",
	0x0880: "System peripheral",
	// Serial Bus
	0x0C03: "USB controller",
	0x0C05: "SMBus",
	// Wireless
	0x0D00: "IRDA controller",
	0x0D11: "Bluetooth",
	0x0D80: "Wireless controller",
	// Signal Processing
	0x1180: "Signal processing controller",
	// Processing Accelerator
	0x1200: "Processing accelerator",
}

// pciBaseClassNames maps base_class to a fallback human-readable name.
var pciBaseClassNames = map[uint8]string{
	0x00: "Unclassified device",
	0x01: "Mass storage controller",
	0x02: "Network controller",
	0x03: "Display controller",
	0x04: "Multimedia controller",
	0x05: "Memory controller",
	0x06: "Bridge",
	0x07: "Communication controller",
	0x08: "System peripheral",
	0x09: "Input device controller",
	0x0A: "Docking station",
	0x0B: "Processor",
	0x0C: "Serial bus controller",
	0x0D: "Wireless controller",
	0x0E: "Intelligent controller",
	0x0F: "Satellite communication controller",
	0x10: "Encryption controller",
	0x11: "Signal processing controller",
	0x12: "Processing accelerator",
	0xFF: "Unassigned class",
}
