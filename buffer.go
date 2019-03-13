package pjn

import (
	"bytes"
	"strconv"
)

type Buffer struct {
	*bytes.Buffer
	backBuf [128]byte
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

func (buf *Buffer) AppendEscapedString(s string) {
	_, _ = buf.WriteRune('"')
	_, _ = buf.WriteString(s)
	_, _ = buf.WriteRune('"')
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

func (buf *Buffer) AppendFloat32(n float32) {
	b := buf.backBuf[:0]
	b = strconv.AppendFloat(b, float64(n), 'g', -1, 32)
	_, _ = buf.Write(b)
}

func (buf *Buffer) AppendFloat64(n float64) {
	b := buf.backBuf[:0]
	b = strconv.AppendFloat(b, float64(n), 'g', -1, 64)
	_, _ = buf.Write(b)
}
