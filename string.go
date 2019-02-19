package pjn

func Str(value string) Produce {
	if len(value) == 0 {
		return produceEmptyString
	} else {
		return func(buf *Buffer) (err error) {
			buf.AppendString(value)
			return nil
		}
	}
}

func NullableStr(ref *string) Produce {
	if ref == nil {
		return produceNull
	} else {
		return Str(*ref)
	}
}

func BindStr(ref *string) Produce {
	if ref == nil {
		return produceError(ErrNilReference)
	} else {
		return func(buf *Buffer) error {
			if len(*ref) == 0 {
				buf.AppendBytes(valueEmptyString)
			} else {
				buf.AppendString(*ref)
			}
			return nil
		}
	}
}

func BindNullableStr(ref **string) Produce {
	if ref == nil {
		return produceError(ErrNilReference)
	} else {
		return func(buf *Buffer) error {
			if *ref == nil {
				buf.AppendBytes(valueNull)
			} else {
				if len(**ref) == 0 {
					buf.AppendBytes(valueEmptyString)
				} else {
					buf.AppendString(**ref)
				}
			}
			return nil
		}
	}
}
