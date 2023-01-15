package nlp

type TokenizerInterface interface {
	Tokenize(ctx *Ctx, p []byte) Tokens
	TokenizeString(ctx *Ctx, s string) Tokens
}

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
