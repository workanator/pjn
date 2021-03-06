package pjn

func Uint32(value uint32) Value {
	if value == 0 {
		return Zero
	} else {
		return func(buf *Buffer) (err error) {
			buf.AppendUint32(value)
			return nil
		}
	}
}

func NullableUint32(ref *uint32) Value {
	if ref == nil {
		return Null
	} else {
		return Uint32(*ref)
	}
}

func BindUint32(ref *uint32) Value {
	if ref == nil {
		return produceError(ErrNilReference)
	} else {
		return func(buf *Buffer) error {
			if *ref == 0 {
				buf.AppendByte(valueNumberZero)
			} else {
				buf.AppendUint32(*ref)
			}
			return nil
		}
	}
}

func BindNullableUint32(ref **uint32) Value {
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
					buf.AppendUint32(**ref)
				}
			}
			return nil
		}
	}
}
