package pjn

func Float64(value float64) Value {
	if value == 0 {
		return Zero
	} else {
		return func(buf *Buffer) (err error) {
			buf.AppendFloat64(value)
			return nil
		}
	}
}

func NullableFloat64(ref *float64) Value {
	if ref == nil {
		return Null
	} else {
		return Float64(*ref)
	}
}

func BindFloat64(ref *float64) Value {
	if ref == nil {
		return produceError(ErrNilReference)
	} else {
		return func(buf *Buffer) error {
			if *ref == 0 {
				buf.AppendByte(valueNumberZero)
			} else {
				buf.AppendFloat64(*ref)
			}
			return nil
		}
	}
}

func BindNullableFloat64(ref **float64) Value {
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
					buf.AppendFloat64(**ref)
				}
			}
			return nil
		}
	}
}
