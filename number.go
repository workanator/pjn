package jsons

import (
	"bytes"
	"strconv"
)

func Int(value int) Produce {
	if value == 0 {
		return produceNumberZero
	} else {
		s := strconv.Itoa(value)
		return func(buf *bytes.Buffer) (err error) {
			_, _ = buf.WriteString(s)
			return nil
		}
	}
}

func IntVar(ref *int) Produce {
	if ref == nil {
		return produceError(ErrNilReference)
	} else {
		return func(buf *bytes.Buffer) error {
			if *ref == 0 {
				_ = buf.WriteByte(valueNumberZero)
			} else {
				_, _ = buf.WriteString(strconv.Itoa(*ref))
			}
			return nil
		}
	}
}

func OptionalIntVar(ref **int) Produce {
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
					_, _ = buf.WriteString(strconv.Itoa(**ref))
				}
			}
			return nil
		}
	}
}

func produceNumberZero(buf *bytes.Buffer) (err error) {
	_ = buf.WriteByte(valueNumberZero)
	return nil
}
