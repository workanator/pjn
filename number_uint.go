package pjn

import (
	"bytes"
	"strconv"
)

func Uint(value uint) Produce {
	if value == 0 {
		return produceNumberZero
	} else {
		s := strconv.FormatUint(uint64(value), 10)
		return func(buf *bytes.Buffer) (err error) {
			_, _ = buf.WriteString(s)
			return nil
		}
	}
}

func NullableIUint(ref *uint) Produce {
	if ref == nil {
		return produceNull
	} else {
		return Uint(*ref)
	}
}

func BindUint(ref *uint) Produce {
	if ref == nil {
		return produceError(ErrNilReference)
	} else {
		return func(buf *bytes.Buffer) error {
			if *ref == 0 {
				_ = buf.WriteByte(valueNumberZero)
			} else {
				_, _ = buf.WriteString(strconv.FormatUint(uint64(*ref), 10))
			}
			return nil
		}
	}
}

func BindNullableUint(ref **uint) Produce {
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
					_, _ = buf.WriteString(strconv.FormatUint(uint64(**ref), 10))
				}
			}
			return nil
		}
	}
}
