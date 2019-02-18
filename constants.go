package pjn

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
