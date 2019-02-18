package jsons

import (
	"bytes"
	"io"
)

type (
	Stream struct {
		buf *bytes.Buffer
	}
	Produce func(*bytes.Buffer) error
)

var (
	valueSeparator  = byte(',')
	nameSeparator   = byte(':')
	doubleQuote     = byte('"')
	beginObject     = byte('{')
	endObject       = byte('}')
	beginArray      = byte('[')
	endArray        = byte(']')
	valueNull       = []byte("null")
	valueNumberZero = byte('0')
)

func NewStream() *Stream {
	return &Stream{
		buf: &bytes.Buffer{},
	}
}

func NewStreamWithBuffer(buf *bytes.Buffer) *Stream {
	return &Stream{
		buf: buf,
	}
}

func (s *Stream) Produce(value Produce) error {
	if s.buf == nil {
		s.buf = &bytes.Buffer{}
	}
	s.buf.Reset()
	return value(s.buf)
}

func (s *Stream) WriteTo(writer io.Writer) (n int64, err error) {
	if writer == nil {
		return 0, io.ErrShortWrite
	} else if s.buf == nil {
		return 0, io.ErrUnexpectedEOF
	}
	return s.buf.WriteTo(writer)
}

func (s *Stream) Bytes() []byte {
	return s.buf.Bytes()
}

func (s *Stream) String() string {
	return s.buf.String()
}
