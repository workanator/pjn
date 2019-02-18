package pjn

import (
	"bytes"
	"errors"
)

func Object(producer ...Produce) Produce {
	return func(buf *bytes.Buffer) (err error) {
		// Write opening bracket
		_ = buf.WriteByte(beginObject)

		// Produce object content
		for n, p := range producer {
			if n > 0 {
				_ = buf.WriteByte(valueSeparator)
			}
			if err = p(buf); err != nil {
				return ObjectProduceFailed(err)
			}
		}

		// Write closing bracket
		_ = buf.WriteByte(endObject)

		return nil
	}
}

func Member(name string, value Produce) Produce {
	if len(name) == 0 {
		return produceError(MemberProduceFailed(errors.New("member name is empty")))
	}

	return func(buf *bytes.Buffer) (err error) {
		// Write name and value separator
		_ = buf.WriteByte(doubleQuote)
		_, _ = buf.WriteString(name)
		_ = buf.WriteByte(doubleQuote)
		_ = buf.WriteByte(nameSeparator)

		// Write value
		return value(buf)
	}
}
