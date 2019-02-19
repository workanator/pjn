package tests

//easyjson:json
type TestObj struct {
	Heights    []int
	Categories struct {
		One int
		Two int
	}
	Awesome bool
	Name    string
	Percent float32
}
