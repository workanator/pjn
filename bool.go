package pjn

import "bytes"

func Bool(value bool) Produce {
	if value {
		return produceBoolTrue
	} else {
		return produceBoolFalse
	}
}

func produceBoolTrue(buf *bytes.Buffer) (err error) {
	_, _ = buf.Write(valueTrue)
	return nil
}

func produceBoolFalse(buf *bytes.Buffer) (err error) {
	_, _ = buf.Write(valueFalse)
	return nil
}
