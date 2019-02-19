package pjn

import (
	"bytes"
	"io"
)

type (
	Producer struct {
		buf Buffer
	}
	Produce func(*Buffer) error
)

func (p *Producer) Produce(value Produce) error {
	if p.buf.Buffer == nil {
		p.buf.Buffer = &bytes.Buffer{}
	}
	p.buf.Reset()
	return value(&p.buf)
}

func (p *Producer) WriteTo(writer io.Writer) (n int64, err error) {
	if writer == nil {
		return 0, io.ErrShortWrite
	}
	return p.buf.WriteTo(writer)
}

func (p *Producer) Bytes() []byte {
	return p.buf.Bytes()
}
