package nlp

func (ctx *Ctx) WithTokenizer(tkn TokenizerInterface) *Ctx {
	ctx.tkn = tkn
	return ctx
}

func (ctx *Ctx) Tokenize(p []byte) Tokens {
	return ctx.chkTkn().Tokenize(ctx.bufT[:0], p)
}

func (ctx *Ctx) TokenizeString(s string) Tokens {
	return ctx.chkTkn().TokenizeString(ctx.bufT[:0], s)
}

func (ctx *Ctx) chkTkn() TokenizerInterface {
	if ctx.tkn == nil {
		ctx.tkn = NewTokenizer()
	}
	return ctx.tkn
}
