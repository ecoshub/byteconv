package byteconv

import "unsafe"

func ToBytes(val interface{}) ([]byte, error) {
	return encoderSwitch(val)
}

func encoderSwitch(val interface{}) ([]byte, error) {
	switch val.(type) {
	case bool:
		return boolEncoder(val.(bool)), nil
	case string:
		return stringEncoder(val.(string)), nil
	case float64:
		return floatEncoder(val.(float64)), nil
	case float32:
		return floatEncoder(float64(val.(float32))), nil
	case int:
		return intEncoder(val.(int)), nil
	case int8:
		return intEncoder(int(val.(int8))), nil
	case int16:
		return intEncoder(int(val.(int16))), nil
	case int32:
		return intEncoder(int(val.(int32))), nil
	case int64:
		return intEncoder(int(val.(int64))), nil
	case uint:
		return intEncoder(int(val.(uint))), nil
	case uint8:
		return intEncoder(int(val.(uint8))), nil
	case uint16:
		return intEncoder(int(val.(uint16))), nil
	case uint32:
		return intEncoder(int(val.(uint32))), nil
	case uint64:
		return intEncoder(int(val.(uint64))), nil
	}
	return nil, unsupType
}

func boolEncoder(val bool) []byte {
	if val {
		return []byte{1}
	}
	return []byte{0}
}

func stringEncoder(val string) []byte {
	return *(*[]byte)(unsafe.Pointer(&val))
}

func intEncoder(val int) []byte {
	size := int(unsafe.Sizeof(val))
	point := uintptr(unsafe.Pointer(&val))
	return coreEncoder(size, point)
}

func floatEncoder(val float64) []byte {
	size := int(unsafe.Sizeof(val))
	point := uintptr(unsafe.Pointer(&val))
	return coreEncoder(size, point)
}

func coreEncoder(size int, point uintptr) []byte {
	arr := make([]byte, size)
	for i := 0; i < size; i++ {
		byt := *(*uint8)(unsafe.Pointer(point + uintptr(i)))
		arr[i] = byt
	}
	return arr
}
