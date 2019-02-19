package pjn

func Int(value int) Produce {
	if value == 0 {
		return produceNumberZero
	} else {
		return func(buf *Buffer) (err error) {
			buf.AppendInt(value)
			return nil
		}
	}
}

func NullableInt(ref *int) Produce {
	if ref == nil {
		return produceNull
	} else {
		return Int(*ref)
	}
}

func BindInt(ref *int) Produce {
	if ref == nil {
		return produceError(ErrNilReference)
	} else {
		return func(buf *Buffer) error {
			if *ref == 0 {
				buf.AppendByte(valueNumberZero)
			} else {
				buf.AppendInt(*ref)
			}
			return nil
		}
	}
}

func BindNullableInt(ref **int) Produce {
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
					buf.AppendInt(**ref)
				}
			}
			return nil
		}
	}
}
