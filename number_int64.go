package pjn

import (
	"bytes"
	"strconv"
)

func Int64(value int64) Produce {
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

func NullableInt64(ref *int64) Produce {
	if ref == nil {
		return produceNull
	} else {
		return Int64(*ref)
	}
}

func BindInt64(ref *int64) Produce {
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

func BindNullableInt64(ref **int64) Produce {
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
