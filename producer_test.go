package pjn

import (
	"testing"
)

func TestProducer_Produce(t *testing.T) {
	s := Producer{}
	err := s.Produce(
		Array(
			Object(
				Member("heights", Array(
					Int(1),
					Int(2),
					Int(3),
				)),
				Member("categories", Object(
					Member("$one", Int(1)),
					Member("$two", Int(2)),
				)),
				Member("awesome", Bool(true)),
			),
			Object(
				Member("heights", Array(
					Int(9),
					Int(8),
					Null,
				)),
				Member("categories", Object(
					Member("$three", Int(3)),
					Member("$four", Int(4)),
					Member("$zero", Int(0)),
				)),
				Member("awesome", Bool(false)),
			),
		),
	)
	if err != nil {
		t.Log(err)
		t.Fail()
	} else {
		t.Log(s.String())
	}
}
