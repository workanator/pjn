package pjn

var (
	EmptyStr Value = produceEmptyString
)

func produceEmptyString(buf *Buffer) (err error) {
	buf.AppendBytes(valueEmptyString)
	return nil
}

func Str(value string) Value {
	if len(value) == 0 {
		return EmptyStr
	} else {
		return func(buf *Buffer) (err error) {
			buf.AppendEscapedString(value)
			return nil
		}
	}
}

func NullableStr(ref *string) Value {
	if ref == nil {
		return produceNull
	} else {
		return Str(*ref)
	}
}

func BindStr(ref *string) Value {
	if ref == nil {
		return produceError(ErrNilReference)
	} else {
		return func(buf *Buffer) error {
			if len(*ref) == 0 {
				buf.AppendBytes(valueEmptyString)
			} else {
				buf.AppendEscapedString(*ref)
			}
			return nil
		}
	}
}

func BindNullableStr(ref **string) Value {
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
					buf.AppendEscapedString(**ref)
				}
			}
			return nil
		}
	}
}
