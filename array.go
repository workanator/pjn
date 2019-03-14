package pjn

var (
	EmptyArray = produceEmptyArray
)

func produceEmptyArray(buf *Buffer) (err error) {
	buf.AppendBytes(emptyArray)
	return nil
}

func Array(producer ...Value) Value {
	if len(producer) == 0 {
		return EmptyArray
	}

	return func(buf *Buffer) (err error) {
		// Write opening bracket
		buf.AppendByte(beginArray)

		// Produce object content
		for n, p := range producer {
			if n > 0 {
				buf.AppendByte(valueSeparator)
			}
			if err = p(buf); err != nil {
				return ArrayProduceFailed(err)
			}
		}

		// Write closing bracket
		buf.AppendByte(endArray)

		return nil
	}
}
