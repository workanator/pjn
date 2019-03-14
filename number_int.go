package pjn

func Int(value int) Value {
	if value == 0 {
		return Zero
	} else {
		return func(buf *Buffer) (err error) {
			buf.AppendInt(value)
			return nil
		}
	}
}

func NullableInt(ref *int) Value {
	if ref == nil {
		return Null
	} else {
		return Int(*ref)
	}
}

func BindInt(ref *int) Value {
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

func BindNullableInt(ref **int) Value {
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
