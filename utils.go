package byteconv

import "errors"

var (
	errNullArray error = errors.New("nil byte array")
)

func sizeCheck(size int) error {
	switch size {
	case 0:
		return errNullArray
	case 8, 16, 32, 64:
		return nil
	}
	return wrongSize
}
