package pjn

import (
	"errors"
)

func Object(producer ...Produce) Produce {
	return func(buf *Buffer) (err error) {
		// Write opening bracket
		buf.AppendByte(beginObject)

		// Produce object content
		for n, p := range producer {
			if n > 0 {
				buf.AppendByte(valueSeparator)
			}
			if err = p(buf); err != nil {
				return ObjectProduceFailed(err)
			}
		}

		// Write closing bracket
		buf.AppendByte(endObject)

		return nil
	}
}

func Member(name string, value Produce) Produce {
	if len(name) == 0 {
		return produceError(MemberProduceFailed(errors.New("member name is empty")))
	}

	return func(buf *Buffer) (err error) {
		// Write name and value separator
		buf.AppendByte(doubleQuote)
		buf.AppendString(name)
		buf.AppendByte(doubleQuote)
		buf.AppendByte(nameSeparator)

		// Write value
		return value(buf)
	}
}
