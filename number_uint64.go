package pjn

func Uint64(value uint64) Value {
	if value == 0 {
		return Zero
	} else {
		return func(buf *Buffer) (err error) {
			buf.AppendUint64(value)
			return nil
		}
	}
}

func NullableUint64(ref *uint64) Value {
	if ref == nil {
		return Null
	} else {
		return Uint64(*ref)
	}
}

func BindUint64(ref *uint64) Value {
	if ref == nil {
		return produceError(ErrNilReference)
	} else {
		return func(buf *Buffer) error {
			if *ref == 0 {
				buf.AppendByte(valueNumberZero)
			} else {
				buf.AppendUint64(*ref)
			}
			return nil
		}
	}
}

func BindNullableUint64(ref **uint64) Value {
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
					buf.AppendUint64(**ref)
				}
			}
			return nil
		}
	}
}
