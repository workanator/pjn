package pjn

import "bytes"

var (
	valueSeparator  = byte(',')
	nameSeparator   = byte(':')
	doubleQuote     = byte('"')
	beginObject     = byte('{')
	endObject       = byte('}')
	beginArray      = byte('[')
	endArray        = byte(']')
	valueNull       = []byte("null")
	valueFalse      = []byte("false")
	valueTrue       = []byte("true")
	valueNumberZero = byte('0')
)

func produceBoolTrue(buf *bytes.Buffer) (err error) {
	_, _ = buf.Write(valueTrue)
	return nil
}

func produceBoolFalse(buf *bytes.Buffer) (err error) {
	_, _ = buf.Write(valueFalse)
	return nil
}

func produceNull(buf *bytes.Buffer) (err error) {
	_, _ = buf.Write(valueNull)
	return nil
}

func produceNumberZero(buf *bytes.Buffer) (err error) {
	_ = buf.WriteByte(valueNumberZero)
	return nil
}
