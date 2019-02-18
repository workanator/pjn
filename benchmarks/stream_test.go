package tests

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/workanator/jsons"
)

var (
	testObj = TestObj{
		Heights: []int{1, 2, 3},
		Categories: struct {
			One int
			Two int
		}{One: 1, Two: 2},
		Awesome: true,
	}
)

func BenchmarkJsonsObject(b *testing.B) {
	s := jsons.NewStream()
	obj := jsons.Object(
		jsons.Member("heights", jsons.Array(
			jsons.Int(1),
			jsons.Int(2),
			jsons.Int(3),
		)),
		jsons.Member("categories", jsons.Object(
			jsons.Member("one", jsons.IntVar(&testObj.Categories.One)),
			jsons.Member("two", jsons.IntVar(&testObj.Categories.Two)),
		)),
		jsons.Member("awesome", jsons.Bool(false)),
	)
	for i := 0; i < b.N; i++ {
		if err := s.Produce(obj); err != nil {
			b.Log(err)
			b.Fail()
		}
	}
}

func BenchmarkStandardMarshalObject(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, err := json.Marshal(&testObj); err != nil {
			b.Log(err)
			b.Fail()
		}
	}
}

func BenchmarkStandardEncodeObject(b *testing.B) {
	buf := &bytes.Buffer{}
	encoder := json.NewEncoder(buf)
	for i := 0; i < b.N; i++ {
		buf.Reset()
		if err := encoder.Encode(&testObj); err != nil {
			b.Log(err)
			b.Fail()
		}
	}
}

func BenchmarkEasyjsonObject(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, err := testObj.MarshalJSON(); err != nil {
			b.Log(err)
			b.Fail()
		}
	}
}
