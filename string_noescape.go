package pjn

func StrNoEscape(value string) Value {
	if len(value) == 0 {
		return EmptyStr
	} else {
		return func(buf *Buffer) (err error) {
			buf.AppendUnescapedString(value)
			return nil
		}
	}
}

func NullableStrNoEscape(ref *string) Value {
	if ref == nil {
		return produceNull
	} else {
		return StrNoEscape(*ref)
	}
}

func BindStrNoEscape(ref *string) Value {
	if ref == nil {
		return produceError(ErrNilReference)
	} else {
		return func(buf *Buffer) error {
			if len(*ref) == 0 {
				buf.AppendBytes(valueEmptyString)
			} else {
				buf.AppendUnescapedString(*ref)
			}
			return nil
		}
	}
}

func BindNullableStrNoEscape(ref **string) Value {
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
					buf.AppendUnescapedString(**ref)
				}
			}
			return nil
		}
	}
}
