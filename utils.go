package byteconv

func sizeCheck(size int) error {
	switch size {
	case 0:
		return nilArray
	case 8, 16, 32, 64:
		return nil
	}
	return wrongSize
}
