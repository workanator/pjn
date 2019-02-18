package pjn

import (
	"bytes"
	"strconv"
)

func Uint64(value uint64) Produce {
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

func NullableUint64(ref *uint64) Produce {
	if ref == nil {
		return produceNull
	} else {
		return Uint64(*ref)
	}
}

func BindUint64(ref *uint64) Produce {
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

func BindNullableUint64(ref **uint64) Produce {
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
