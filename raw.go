package pjn

func RawBytes(value []byte) Value {
	if len(value) == 0 {
		return produceError(ErrEmptySlice)
	} else {
		return func(buf *Buffer) (err error) {
			buf.AppendBytes(value)
			return nil
		}
	}
}
