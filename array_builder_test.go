package pjn

import (
	"testing"
)

func TestArrayBuilder_Build(t *testing.T) {
	s := Producer{}
	arr := ArrayBuilder{}.
		Push(Int(1)).
		Push(Str("TWO!")).
		Push(Array(
			Int(1),
			Int(2),
			Int(3),
		)).
		Build()
	err := s.Produce(arr)
	if err != nil {
		t.Log(err)
		t.Fail()
	} else {
		t.Log(s.String())
	}
}
