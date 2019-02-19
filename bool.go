package pjn

func Bool(value bool) Produce {
	if value {
		return produceBoolTrue
	} else {
		return produceBoolFalse
	}
}

func NullableBool(ref *bool) Produce {
	if ref == nil {
		return produceNull
	} else {
		return Bool(*ref)
	}
}

func BindBool(ref *bool) Produce {
	if ref == nil {
		return produceError(ErrNilReference)
	} else {
		return func(buf *Buffer) error {
			if *ref {
				buf.AppendBytes(valueTrue)
			} else {
				buf.AppendBytes(valueFalse)
			}
			return nil
		}
	}
}

func BindNullableBool(ref **bool) Produce {
	if ref == nil {
		return produceError(ErrNilReference)
	} else {
		return func(buf *Buffer) error {
			if *ref == nil {
				buf.AppendBytes(valueNull)
			} else {
				if **ref {
					buf.AppendBytes(valueTrue)
				} else {
					buf.AppendBytes(valueFalse)
				}
			}
			return nil
		}
	}
}
