package pjn

import (
	"bytes"
	"strconv"
)

func Int16(value int16) Produce {
	if value == 0 {
		return produceNumberZero
	} else {
		s := strconv.FormatInt(int64(value), 10)
		return func(buf *bytes.Buffer) (err error) {
			_, _ = buf.WriteString(s)
			return nil
		}
	}
}

func NullableInt16(ref *int16) Produce {
	if ref == nil {
		return produceNull
	} else {
		return Int16(*ref)
	}
}

func BindInt16(ref *int16) Produce {
	if ref == nil {
		return produceError(ErrNilReference)
	} else {
		return func(buf *bytes.Buffer) error {
			if *ref == 0 {
				_ = buf.WriteByte(valueNumberZero)
			} else {
				_, _ = buf.WriteString(strconv.FormatInt(int64(*ref), 10))
			}
			return nil
		}
	}
}

func BindNullableInt16(ref **int16) Produce {
	if ref == nil {
		return produceError(ErrNilReference)
	} else {
		return func(buf *bytes.Buffer) error {
			if *ref == nil {
				_, _ = buf.Write(valueNull)
			} else {
				if **ref == 0 {
					_ = buf.WriteByte(valueNumberZero)
				} else {
					_, _ = buf.WriteString(strconv.FormatInt(int64(**ref), 10))
				}
			}
			return nil
		}
	}
}
