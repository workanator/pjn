package pjn

var (
	valueSeparator   = byte(',')
	nameSeparator    = byte(':')
	doubleQuote      = byte('"')
	beginObject      = byte('{')
	endObject        = byte('}')
	beginArray       = byte('[')
	endArray         = byte(']')
	valueNull        = []byte("null")
	valueFalse       = []byte("false")
	valueTrue        = []byte("true")
	valueNumberZero  = byte('0')
	valueEmptyString = []byte("\"\"")
)

func produceBoolTrue(buf *Buffer) (err error) {
	buf.AppendBytes(valueTrue)
	return nil
}

func produceBoolFalse(buf *Buffer) (err error) {
	buf.AppendBytes(valueFalse)
	return nil
}

func produceNull(buf *Buffer) (err error) {
	buf.AppendBytes(valueNull)
	return nil
}

func produceNumberZero(buf *Buffer) (err error) {
	buf.AppendByte(valueNumberZero)
	return nil
}

func produceEmptyString(buf *Buffer) (err error) {
	buf.AppendBytes(valueEmptyString)
	return nil
}
