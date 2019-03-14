package pjn

import "errors"

type ObjectBuilder struct {
	p Value
}

func (ob ObjectBuilder) Member(name string, value Value) ObjectBuilder {
	if len(name) == 0 {
		return ObjectBuilder{
			p: produceError(MemberProduceFailed(errors.New("empty member name"))),
		}
	}
	if value == nil {
		return ObjectBuilder{
			p: produceError(MemberProduceFailed(errors.New("nil member value producer"))),
		}
	}

	return ObjectBuilder{
		p: func(buffer *Buffer) (err error) {
			if ob.p != nil {
				if err = ob.p(buffer); err != nil {
					return err
				}
				_ = buffer.WriteByte(valueSeparator)
			}
			return Member(name, value)(buffer)
		},
	}
}

func (ob ObjectBuilder) Build() Value {
	if ob.p == nil {
		return EmptyObject
	}
	return Object(ob.p)
}
