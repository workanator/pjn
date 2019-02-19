package pjn

func Int32(value int32) Produce {
	if value == 0 {
		return produceNumberZero
	} else {
		return func(buf *Buffer) (err error) {
			buf.AppendInt32(value)
			return nil
		}
	}
}

func NullableInt32(ref *int32) Produce {
	if ref == nil {
		return produceNull
	} else {
		return Int32(*ref)
	}
}

func BindInt32(ref *int32) Produce {
	if ref == nil {
		return produceError(ErrNilReference)
	} else {
		return func(buf *Buffer) error {
			if *ref == 0 {
				buf.AppendByte(valueNumberZero)
			} else {
				buf.AppendInt32(*ref)
			}
			return nil
		}
	}
}

func BindNullableInt32(ref **int32) Produce {
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
					buf.AppendInt32(**ref)
				}
			}
			return nil
		}
	}
}
