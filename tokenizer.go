package nlp

type Tokenizer interface {
	Tokenize(p []byte)
	TokenizeString(s string)
	Each(i int, fn func([]byte))
	EachString(i int, fn func(string))
	EachRange(i int, fn func(int, int))
}
