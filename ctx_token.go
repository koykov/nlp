package nlp

import "github.com/koykov/byteseq"

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
	ctx.bufT = ctx.chkTkn().AppendTokenize(ctx.bufT, byteseq.B2Q[T](ctx.buf))
	return ctx
}

func (ctx *Ctx[T]) GetTokens() Tokens {
	return ctx.bufT
}

func (ctx *Ctx[T]) ResetTokenizer() *Ctx[T] {
	ctx.tkn = nil
	return ctx
}

func (ctx *Ctx[T]) chkTkn() Tokenizer[T] {
	if ctx.tkn == nil {
		tkn := NewStringTokenizer[T](DefaultTokenSeparator, TokenizerBlankLinesDiscard)
		ctx.WithTokenizer(tkn)
	}
	return ctx.tkn
}
