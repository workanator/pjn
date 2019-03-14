package pjn

var (
	True  Value = produceBoolTrue
	False Value = produceBoolFalse
)

func produceBoolTrue(buf *Buffer) (err error) {
	buf.AppendBytes(valueTrue)
	return nil
}

func produceBoolFalse(buf *Buffer) (err error) {
	buf.AppendBytes(valueFalse)
	return nil
}

func Bool(value bool) Value {
	if value {
		return True
	} else {
		return False
	}
}

func NullableBool(ref *bool) Value {
	if ref == nil {
		return Null
	} else {
		return Bool(*ref)
	}
}

func BindBool(ref *bool) Value {
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

func BindNullableBool(ref **bool) Value {
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
