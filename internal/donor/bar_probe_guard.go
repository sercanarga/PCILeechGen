package donor

func isNetworkClass(classCode uint32) bool {
	return (classCode>>16)&0xFF == 0x02
}

func shouldProbeBAR(enabled bool, classCode uint32, content []byte) bool {
	if !enabled {
		return false
	}
	if len(content) > 0 && isAllFF(content) {
		return false
	}
	if isNetworkClass(classCode) {
		return false
	}
	return true
}
