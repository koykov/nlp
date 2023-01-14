package nlp

type TokenizerInterface interface {
	Tokenize(p []byte)
	TokenizeString(s string)
	Each(i int, fn func([]byte))
	EachString(i int, fn func(string))
	EachSpan(i int, fn func(int, int))
}
