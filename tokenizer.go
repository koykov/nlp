package nlp

type TokenizerInterface interface {
	Tokenize(ctx *Ctx, p []byte) *Tokens
	TokenizeString(ctx *Ctx, s string) *Tokens
}

const tknSep = " \n\t"

type Tokenizer struct {
	sep string
	bl  bool
	eof bool
}

func NewTokenizer() Tokenizer {
	return Tokenizer{sep: tknSep}
}

func NewTokenizerWithOptions(sep string, keepBlank bool, discardEOF bool) Tokenizer {
	return Tokenizer{
		sep: sep,
		bl:  keepBlank,
		eof: discardEOF,
	}
}

func (t Tokenizer) Tokenize(ctx *Ctx, p []byte) *Tokens {
	ctx.BufT.SetSource(p)
	return t.tokenize(&ctx.BufT)
}

func (t Tokenizer) TokenizeString(ctx *Ctx, s string) *Tokens {
	ctx.BufT.SetSourceString(s)
	return t.tokenize(&ctx.BufT)
}

func (t Tokenizer) tokenize(buf *Tokens) *Tokens {
	// ...
	return buf
}
