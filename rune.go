package pjn

func Rune(value rune) Value {
	if value == 0 {
		return produceEmptyString
	} else {
		return func(buf *Buffer) (err error) {
			buf.AppendEscapedRune(value)
			return nil
		}
	}
}

func NullableRune(ref *rune) Value {
	if ref == nil {
		return produceNull
	} else {
		return Rune(*ref)
	}
}

func BindRune(ref *rune) Value {
	if ref == nil {
		return produceError(ErrNilReference)
	} else {
		return func(buf *Buffer) error {
			if *ref == 0 {
				buf.AppendBytes(valueEmptyString)
			} else {
				buf.AppendEscapedRune(*ref)
			}
			return nil
		}
	}
}

func BindNullableRune(ref **rune) Value {
	if ref == nil {
		return produceError(ErrNilReference)
	} else {
		return func(buf *Buffer) error {
			if *ref == nil {
				buf.AppendBytes(valueNull)
			} else {
				if **ref == 0 {
					buf.AppendBytes(valueEmptyString)
				} else {
					buf.AppendEscapedRune(**ref)
				}
			}
			return nil
		}
	}
}
