package pjn

var (
	valueSeparator   = byte(',')
	nameSeparator    = byte(':')
	doubleQuote      = byte('"')
	beginObject      = byte('{')
	endObject        = byte('}')
	emptyObject      = []byte("{}")
	beginArray       = byte('[')
	endArray         = byte(']')
	emptyArray       = []byte("[]")
	valueNull        = []byte("null")
	valueFalse       = []byte("false")
	valueTrue        = []byte("true")
	valueNumberZero  = byte('0')
	valueEmptyString = []byte("\"\"")
)
