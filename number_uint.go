package pjn

func Uint(value uint) Value {
	if value == 0 {
		return Zero
	} else {
		return func(buf *Buffer) (err error) {
			buf.AppendUint(value)
			return nil
		}
	}
}

func NullableUint(ref *uint) Value {
	if ref == nil {
		return Null
	} else {
		return Uint(*ref)
	}
}

func BindUint(ref *uint) Value {
	if ref == nil {
		return produceError(ErrNilReference)
	} else {
		return func(buf *Buffer) error {
			if *ref == 0 {
				buf.AppendByte(valueNumberZero)
			} else {
				buf.AppendUint(*ref)
			}
			return nil
		}
	}
}

func BindNullableUint(ref **uint) Value {
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
					buf.AppendUint(**ref)
				}
			}
			return nil
		}
	}
}
