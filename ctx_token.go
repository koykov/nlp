package nlp

func (ctx *Ctx[T]) WithTokenizer(tkn Tokenizer[T]) *Ctx[T] {
	ctx.tkn = tkn
	ctx.SetBit(flagToken, false)
	return ctx
}

func (ctx *Ctx[T]) Tokenize() *Ctx[T] {
	if ctx.chkSrcErr() {
		return ctx
	}
	if ctx.CheckBit(flagToken) {
		return ctx
	}
	defer ctx.SetBit(flagToken, true)
	ctx.bufT = ctx.chkTkn().Tokenize(ctx.bufT, T(ctx.buf))
	return ctx
}

func (ctx *Ctx[T]) TokenizeT(t T) Tokens {
	return ctx.SetText(t).
		Tokenize().
		GetTokens()
}

func (ctx Ctx[T]) GetTokens() Tokens {
	return ctx.bufT
}

func (ctx *Ctx[T]) chkTkn() Tokenizer[T] {
	if ctx.tkn == nil {
		tkn := NewStringTokenizer[T](DefaultTokenSeparator, false, true)
		ctx.tkn = tkn
	}
	return ctx.tkn
}