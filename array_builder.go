package pjn

import "errors"

type ArrayBuilder struct {
	p Value
}

func (ab ArrayBuilder) Push(value Value) ArrayBuilder {
	if value == nil {
		return ArrayBuilder{
			p: produceError(ValueProduceFailed(errors.New("nil value producer"))),
		}
	}

	return ArrayBuilder{
		p: func(buffer *Buffer) (err error) {
			if ab.p != nil {
				if err = ab.p(buffer); err != nil {
					return err
				}
				_ = buffer.WriteByte(valueSeparator)
			}
			return value(buffer)
		},
	}
}

func (ab ArrayBuilder) Build() Value {
	if ab.p == nil {
		return EmptyArray
	}
	return Array(ab.p)
}
