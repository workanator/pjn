package tests

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/mailru/easyjson/jwriter"
	"github.com/workanator/pjn"
)

var (
	testObj = TestObj{
		Heights: []int{1, 2, 3},
		Categories: struct {
			One int
			Two int
		}{One: 1, Two: 2},
		Awesome: true,
		Percent: 0.99,
	}
)

func BenchmarkPjnStaticObject(b *testing.B) {
	s := pjn.Producer{}
	obj := pjn.Object(
		pjn.Member("heights", pjn.Array(
			pjn.Int(1),
			pjn.Int(2),
			pjn.Int(3),
		)),
		pjn.Member("categories", pjn.Object(
			pjn.Member("one", pjn.BindInt(&testObj.Categories.One)),
			pjn.Member("two", pjn.BindInt(&testObj.Categories.Two)),
		)),
		pjn.Member("awesome", pjn.Bool(false)),
		pjn.Member("percent", pjn.Float32(0.99)),
	)
	for i := 0; i < b.N; i++ {
		if err := s.Produce(obj); err != nil {
			b.Log(err)
			b.Fail()
		}
		_ = s.Bytes()
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

func BenchmarkEasyjsonMarshalObject(b *testing.B) {
	for i := 0; i < b.N; i++ {
		w := jwriter.Writer{}
		testObj.MarshalEasyJSON(&w)
	}
}
