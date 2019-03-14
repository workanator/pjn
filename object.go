package pjn

import (
	"errors"
)

var (
	EmptyObject = produceEmptyObject
)

func produceEmptyObject(buf *Buffer) (err error) {
	buf.AppendBytes(emptyObject)
	return nil
}

func Object(producer ...Value) Value {
	if len(producer) == 0 {
		return EmptyObject
	}

	return func(buf *Buffer) (err error) {
		// Write opening bracket
		buf.AppendByte(beginObject)

		// Value object content
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

func Member(name string, value Value) Value {
	if len(name) == 0 {
		return produceError(MemberProduceFailed(errors.New("empty member name")))
	}
	if value == nil {
		return produceError(MemberProduceFailed(errors.New("nil member value producer")))
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
