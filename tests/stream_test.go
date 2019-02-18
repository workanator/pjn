package tests

import (
	"testing"

	"github.com/workanator/jsons"
)

func TestAll(t *testing.T) {
	s := jsons.NewStream()
	err := s.Produce(
		jsons.Array(
			jsons.Object(
				jsons.Member("heights", jsons.Array(
					jsons.Int(1),
					jsons.Int(2),
					jsons.Int(3),
				)),
				jsons.Member("categories", jsons.Object(
					jsons.Member("$one", jsons.Int(1)),
					jsons.Member("$two", jsons.Int(2)),
				)),
				jsons.Member("awesome", jsons.Bool(true)),
			),
			jsons.Object(
				jsons.Member("heights", jsons.Array(
					jsons.Int(9),
					jsons.Int(8),
					jsons.Null(),
				)),
				jsons.Member("categories", jsons.Object(
					jsons.Member("$three", jsons.Int(3)),
					jsons.Member("$four", jsons.Int(4)),
					jsons.Member("$zero", jsons.Int(0)),
				)),
				jsons.Member("awesome", jsons.Bool(false)),
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
