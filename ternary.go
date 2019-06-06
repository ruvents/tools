package tools

func DefaultStringValue(str string, defaultValue string) string {
	if str == "" {
		return defaultValue
	} else {
		return str
	}
}

func TerString(bool bool, true, false string) string {
	if bool {
		return true
	} else {
		return false
	}
}

func TerNumber(bool bool, true, false int) int {
	if bool {
		return true
	} else {
		return false
	}
}

func TerUINT8(bool bool, true, false uint8) uint8 {
	if bool {
		return true
	} else {
		return false
	}
}