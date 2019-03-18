package pjn

func RuneNoEscape(value rune) Value {
	if value == 0 {
		return produceEmptyString
	} else {
		return func(buf *Buffer) (err error) {
			buf.AppendUnescapedRune(value)
			return nil
		}
	}
}

func NullableRuneNoEscape(ref *rune) Value {
	if ref == nil {
		return produceNull
	} else {
		return RuneNoEscape(*ref)
	}
}

func BindRuneNoEscape(ref *rune) Value {
	if ref == nil {
		return produceError(ErrNilReference)
	} else {
		return func(buf *Buffer) error {
			if *ref == 0 {
				buf.AppendBytes(valueEmptyString)
			} else {
				buf.AppendUnescapedRune(*ref)
			}
			return nil
		}
	}
}

func BindNullableRuneNoEscape(ref **rune) Value {
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
					buf.AppendUnescapedRune(**ref)
				}
			}
			return nil
		}
	}
}
