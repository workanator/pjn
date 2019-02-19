package pjn

func Array(producer ...Produce) Produce {
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
