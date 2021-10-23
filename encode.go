package byteconv

import (
	"errors"
	"unsafe"
)

var (
	errTypeNotSupported error = errors.New("unsupported type")
)

func ToBytes(val interface{}) ([]byte, error) {
	return encoderSwitch(val)
}

func encoderSwitch(val interface{}) ([]byte, error) {
	switch t := val.(type) {
	case bool:
		return boolEncoder(t), nil
	case string:
		return stringEncoder(t), nil
	case float64:
		return floatEncoder(t), nil
	case float32:
		return floatEncoder(float64(t)), nil
	case int:
		return intEncoder(t), nil
	case int8:
		return intEncoder(int(t)), nil
	case int16:
		return intEncoder(int(t)), nil
	case int32:
		return intEncoder(int(t)), nil
	case int64:
		return intEncoder(int(t)), nil
	case uint:
		return intEncoder(int(t)), nil
	case uint8:
		return intEncoder(int(t)), nil
	case uint16:
		return intEncoder(int(t)), nil
	case uint32:
		return intEncoder(int(t)), nil
	case uint64:
		return intEncoder(int(t)), nil
	default:
		return nil, errTypeNotSupported
	}
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
