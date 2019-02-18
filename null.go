package jsons

import "bytes"

func Null() Produce {
	return produceNull
}

func produceNull(buf *bytes.Buffer) (err error) {
	_, _ = buf.Write(valueNull)
	return nil
}
