package tests

import (
	"testing"

	"github.com/workanator/pjn"
)

func TestAll(t *testing.T) {
	s := pjn.Producer{}
	err := s.Produce(
		pjn.Array(
			pjn.Object(
				pjn.Member("heights", pjn.Array(
					pjn.Int(1),
					pjn.Int(2),
					pjn.Int(3),
				)),
				pjn.Member("categories", pjn.Object(
					pjn.Member("$one", pjn.Int(1)),
					pjn.Member("$two", pjn.Int(2)),
				)),
				pjn.Member("awesome", pjn.Bool(true)),
			),
			pjn.Object(
				pjn.Member("heights", pjn.Array(
					pjn.Int(9),
					pjn.Int(8),
					pjn.Null,
				)),
				pjn.Member("categories", pjn.Object(
					pjn.Member("$three", pjn.Int(3)),
					pjn.Member("$four", pjn.Int(4)),
					pjn.Member("$zero", pjn.Int(0)),
				)),
				pjn.Member("awesome", pjn.Bool(false)),
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

func TestObjectBuilder(t *testing.T) {
	s := pjn.Producer{}
	obj := pjn.ObjectBuilder{}.
		Member("one", pjn.Int(1)).
		Member("$two", pjn.Str("TWO!")).
		Member("heights", pjn.Array(
			pjn.Int(1),
			pjn.Int(2),
			pjn.Int(3),
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

func TestArrayBuilder(t *testing.T) {
	s := pjn.Producer{}
	arr := pjn.ArrayBuilder{}.
		Push(pjn.Int(1)).
		Push(pjn.Str("TWO!")).
		Push(pjn.Array(
			pjn.Int(1),
			pjn.Int(2),
			pjn.Int(3),
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
