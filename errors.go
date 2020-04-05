package byteconv

import (
	"errorx"
)

var (
	nilArray  error = errorx.New("Warning", "Nil byte array", 0)
	wrongSize error = errorx.New("Error", "Wrong size byte array", 1)
	unsupType error = errorx.New("Error", "Unsupported type", 1)
)
