package pjn

func Uint(value uint) Produce {
	if value == 0 {
		return produceNumberZero
	} else {
		return func(buf *Buffer) (err error) {
			buf.AppendUint(value)
			return nil
		}
	}
}

func NullableUint(ref *uint) Produce {
	if ref == nil {
		return produceNull
	} else {
		return Uint(*ref)
	}
}

func BindUint(ref *uint) Produce {
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

func BindNullableUint(ref **uint) Produce {
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
