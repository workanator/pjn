package pjn

import "bytes"

func Array(producer ...Produce) Produce {
	return func(buf *bytes.Buffer) (err error) {
		// Write opening bracket
		_ = buf.WriteByte(beginArray)

		// Produce object content
		for n, p := range producer {
			if n > 0 {
				_ = buf.WriteByte(valueSeparator)
			}
			if err = p(buf); err != nil {
				return ArrayProduceFailed(err)
			}
		}

		// Write closing bracket
		_ = buf.WriteByte(endArray)

		return nil
	}
}
