package pjn

import (
	"bytes"
	"strconv"
)

func Int(value int) Produce {
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

func NullableInt(ref *int) Produce {
	if ref == nil {
		return produceNull
	} else {
		return Int(*ref)
	}
}

func BindInt(ref *int) Produce {
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

func BindNullableInt(ref **int) Produce {
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