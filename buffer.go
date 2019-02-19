package pjn

import (
	"bytes"
	"strconv"
)

type Buffer struct {
	*bytes.Buffer
}

func (buf *Buffer) AppendByte(b byte) {
	_ = buf.WriteByte(b)
}

func (buf *Buffer) AppendBytes(p []byte) {
	_, _ = buf.Write(p)
}

func (buf *Buffer) AppendString(s string) {
	_, _ = buf.WriteString(s)
}

func (buf *Buffer) AppendInt(n int) {
	_, _ = buf.WriteString(strconv.FormatInt(int64(n), 10))
}

func (buf *Buffer) AppendUint(n uint) {
	_, _ = buf.WriteString(strconv.FormatUint(uint64(n), 10))
}

func (buf *Buffer) AppendInt16(n int16) {
	_, _ = buf.WriteString(strconv.FormatInt(int64(n), 10))
}

func (buf *Buffer) AppendUint16(n uint16) {
	_, _ = buf.WriteString(strconv.FormatUint(uint64(n), 10))
}

func (buf *Buffer) AppendInt32(n int32) {
	_, _ = buf.WriteString(strconv.FormatInt(int64(n), 10))
}

func (buf *Buffer) AppendUint32(n uint32) {
	_, _ = buf.WriteString(strconv.FormatUint(uint64(n), 10))
}

func (buf *Buffer) AppendInt64(n int64) {
	_, _ = buf.WriteString(strconv.FormatInt(int64(n), 10))
}

func (buf *Buffer) AppendUint64(n uint64) {
	_, _ = buf.WriteString(strconv.FormatUint(uint64(n), 10))
}
