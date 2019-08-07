package tools

import "strconv"

func MakeStringFromUINT8(number uint8) string {
	return strconv.FormatUint(uint64(number), 10)
}

func MakeStringFromUINT16(number uint16) string {
	return strconv.FormatUint(uint64(number), 10)
}

func MakeStringFromUINT16(number uint16) string {
	return strconv.FormatUint(uint64(number), 10)
}

func MakeStringFromUINT32(number uint32) string {
	return strconv.FormatUint(uint64(number), 10)
}

func MakeStringFromUINT64(number uint64) string {
	return strconv.FormatUint(number, 10)
}

func MakeStringFromUINT(number uint) string {
	return strconv.FormatUint(uint64(number), 10)
}

func MakeStringFromINT8(number int8) string {
	return strconv.FormatInt(int64(number), 10)
}

func MakeStringFromINT32(number int32) string {
	return strconv.FormatInt(int64(number), 10)
}

func MakeStringFromINT64(number int64) string {
	return strconv.FormatInt(number, 10)
}

func MakeStringFromINT(number int) string {
	return strconv.FormatInt(int64(number), 10)
}

func MakeUINT64FromString(number string) uint64 {
	if result, err := strconv.ParseInt(number, 10, 64); err == nil {
		return uint64(result)
	} else {
		return 0
	}
}

func MakeUINT64FromStringUnsafe(number string) (uint64, error) {
	result, err := strconv.ParseInt(number, 10, 64)
	return uint64(result), err
}

func MakeUINT32FromString(number string) uint32 {
	if result, err := strconv.Atoi(number); err == nil {
		return uint32(result)
	} else {
		return 0
	}
}

func MakeUINT32FromStringUnsafe(number string) (uint32, error) {
	result, err := strconv.Atoi(number)
	return uint32(result), err
}

func MakeUINT16FromString(number string) uint16 {
	if result, err := strconv.Atoi(number); err == nil {
		return uint16(result)
	} else {
		return 0
	}
}

func MakeUINT16FromStringUnsafe(number string) (uint16, error) {
	result, err := strconv.Atoi(number)
	return uint16(result), err
}

func MakeUINT8FromString(number string) uint8 {
	if result, err := strconv.Atoi(number); err == nil {
		return uint8(result)
	} else {
		return 0
	}
}

func MakeUINT8FromStringUnsafe(number string) (uint8, error) {
	result, err := strconv.Atoi(number)
	return uint8(result), err
}

func MakeUINT32FromHEX(str string) (uint32, error) {
	res, err := strconv.ParseInt(str, 16, 32)
	return uint32(res), err
}

func MakeUINT16FromHEX(str string) (uint16, error) {
	res, err := strconv.ParseInt(str, 16, 32)
	return uint16(res), err
}

func MakeUINT8FromHEX(str string) (uint8, error) {
	res, err := strconv.ParseInt(str, 16, 32)
	return uint8(res), err
}
