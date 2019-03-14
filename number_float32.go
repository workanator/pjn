package pjn

func Float32(value float32) Value {
	if value == 0 {
		return Zero
	} else {
		return func(buf *Buffer) (err error) {
			buf.AppendFloat32(value)
			return nil
		}
	}
}

func NullableFloat32(ref *float32) Value {
	if ref == nil {
		return Null
	} else {
		return Float32(*ref)
	}
}

func BindFloat32(ref *float32) Value {
	if ref == nil {
		return produceError(ErrNilReference)
	} else {
		return func(buf *Buffer) error {
			if *ref == 0 {
				buf.AppendByte(valueNumberZero)
			} else {
				buf.AppendFloat32(*ref)
			}
			return nil
		}
	}
}

func BindNullableFloat32(ref **float32) Value {
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
					buf.AppendFloat32(**ref)
				}
			}
			return nil
		}
	}
}
