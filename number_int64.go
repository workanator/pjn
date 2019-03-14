package pjn

func Int64(value int64) Value {
	if value == 0 {
		return Zero
	} else {
		return func(buf *Buffer) (err error) {
			buf.AppendInt64(value)
			return nil
		}
	}
}

func NullableInt64(ref *int64) Value {
	if ref == nil {
		return Null
	} else {
		return Int64(*ref)
	}
}

func BindInt64(ref *int64) Value {
	if ref == nil {
		return produceError(ErrNilReference)
	} else {
		return func(buf *Buffer) error {
			if *ref == 0 {
				buf.AppendByte(valueNumberZero)
			} else {
				buf.AppendInt64(*ref)
			}
			return nil
		}
	}
}

func BindNullableInt64(ref **int64) Value {
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
					buf.AppendInt64(**ref)
				}
			}
			return nil
		}
	}
}
