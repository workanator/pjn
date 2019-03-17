package pjn

import (
	"testing"
)

func TestObjectBuilder_Build(t *testing.T) {
	s := Producer{}
	obj := ObjectBuilder{}.
		Member("one", Int(1)).
		Member("$two", Str("TWO!")).
		Member("heights", Array(
			Int(1),
			Int(2),
			Int(3),
		)).
		Build()
	err := s.Produce(obj)
	if err != nil {
		t.Log(err)
		t.Fail()
	} else {
		t.Log(s.String())
	}
}
