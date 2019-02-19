package pjn

func Int16(value int16) Produce {
	if value == 0 {
		return produceNumberZero
	} else {
		return func(buf *Buffer) (err error) {
			buf.AppendInt16(value)
			return nil
		}
	}
}

func NullableInt16(ref *int16) Produce {
	if ref == nil {
		return produceNull
	} else {
		return Int16(*ref)
	}
}

func BindInt16(ref *int16) Produce {
	if ref == nil {
		return produceError(ErrNilReference)
	} else {
		return func(buf *Buffer) error {
			if *ref == 0 {
				buf.AppendByte(valueNumberZero)
			} else {
				buf.AppendInt16(*ref)
			}
			return nil
		}
	}
}

func BindNullableInt16(ref **int16) Produce {
	if ref == nil {
		return produceError(ErrNilReference)
	} else {
		return func(buf *Buffer) error {
			if *ref == nil {
				buf.AppendBytes(valueNull)
			} else {
				if **ref == 0 {
					buf.AppendByte(valueNumberZero)
				} else {
					buf.AppendInt16(**ref)
				}
			}
			return nil
		}
	}
}
