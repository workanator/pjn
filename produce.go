package pjn

import (
	"bytes"
	"io"
)

type (
	Producer struct {
		buf *bytes.Buffer
	}
	Produce func(*bytes.Buffer) error
)

func NewProducer() *Producer {
	return &Producer{
		buf: &bytes.Buffer{},
	}
}

func NewProducerWithBuffer(buf *bytes.Buffer) *Producer {
	return &Producer{
		buf: buf,
	}
}

func (p *Producer) Produce(value Produce) error {
	if p.buf == nil {
		p.buf = &bytes.Buffer{}
	}
	p.buf.Reset()
	return value(p.buf)
}

func (p *Producer) ProduceToBuffer(buf *bytes.Buffer, value Produce) error {
	p.buf = buf
	return p.Produce(value)
}

func (p *Producer) WriteTo(writer io.Writer) (n int64, err error) {
	if writer == nil {
		return 0, io.ErrShortWrite
	} else if p.buf == nil {
		return 0, io.ErrUnexpectedEOF
	}
	return p.buf.WriteTo(writer)
}

func (p *Producer) CaptureBuffer() (buf *bytes.Buffer) {
	buf, p.buf = p.buf, nil
	return
}

func (p *Producer) Bytes() []byte {
	return p.buf.Bytes()
}

func (p *Producer) String() string {
	return p.buf.String()
}
