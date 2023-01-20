package nlp

func (ctx *Ctx) WithTokenizer(tkn TokenizerInterface) *Ctx {
	ctx.tkn = tkn
	ctx.SetBit(flagToken, false)
	return ctx
}

func (ctx *Ctx) Tokenize() *Ctx {
	if ctx.CheckBit(flagToken) {
		return ctx
	}
	defer ctx.SetBit(flagToken, true)
	ctx.bufT = ctx.chkTkn().TokenizeBytes(ctx.bufT, ctx.bufC)
	return ctx
}

func (ctx *Ctx) TokenizeBytes(p []byte) Tokens {
	return ctx.SetText(p).Tokenize().GetTokens()
}

func (ctx *Ctx) TokenizeString(s string) Tokens {
	return ctx.SetTextString(s).Tokenize().GetTokens()
}

func (ctx Ctx) GetTokens() Tokens {
	return ctx.bufT
}

func (ctx *Ctx) chkTkn() TokenizerInterface {
	if ctx.tkn == nil {
		ctx.tkn = NewTokenizer()
	}
	return ctx.tkn
}
