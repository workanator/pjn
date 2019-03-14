package pjn

func Uint16(value uint16) Value {
	if value == 0 {
		return Zero
	} else {
		return func(buf *Buffer) (err error) {
			buf.AppendUint16(value)
			return nil
		}
	}
}

func NullableUint16(ref *uint16) Value {
	if ref == nil {
		return Null
	} else {
		return Uint16(*ref)
	}
}

func BindUint16(ref *uint16) Value {
	if ref == nil {
		return produceError(ErrNilReference)
	} else {
		return func(buf *Buffer) error {
			if *ref == 0 {
				buf.AppendByte(valueNumberZero)
			} else {
				buf.AppendUint16(*ref)
			}
			return nil
		}
	}
}

func BindNullableUint16(ref **uint16) Value {
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
					buf.AppendUint16(**ref)
				}
			}
			return nil
		}
	}
}
