package nlp

type Tokens []int

func (t Tokens) Each(i int, fn func(p []byte)) {
	// ...
}

func (t Tokens) EachString(i int, fn func(p string)) {
	// ...
}

func (t Tokens) EachSpan(i int, fn func(a, b int)) {
	//
}
