package formats

import (
	"bytes"
)

func TrimEmptySpace(buf []byte) []byte {
	buf = bytes.ReplaceAll(buf, []byte(" "), []byte(""))
	buf = bytes.ReplaceAll(buf, []byte("\t"), []byte(""))
	buf = bytes.ReplaceAll(buf, []byte("\n"), []byte(""))

	return buf
}
