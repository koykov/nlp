package nlp

import (
	"github.com/koykov/bitset"
	"github.com/koykov/byteseq"
	"github.com/koykov/fastconv"
)

const (
	flagClean = 0
	flagMod   = 1
	flagToken = 2
)

type Ctx[T byteseq.Byteseq] struct {
	bitset.Bitset
	src  T
	buf  []byte
	bufR []rune

	mod []Modifier[T]
	cln []Cleaner[T]

	tkn  Tokenizer[T]
	bufT Tokens

	sd    ScriptDetector[T]
	bufSC []Script
	BufSP ScriptProba

	// BufLP LanguageProba

	err error
}

func NewCtx[T byteseq.Byteseq]() *Ctx[T] {
	ctx := Ctx[T]{}
	return &ctx
}

func (ctx *Ctx[T]) SetText(text T) *Ctx[T] {
	ctx.SetBit(flagClean, false)
	ctx.SetBit(flagToken, false)
	ctx.src = text
	ctx.buf = append(ctx.buf[:0], ctx.src...)
	return ctx
}

func (ctx *Ctx[T]) GetOriginText() T {
	return ctx.src
}

func (ctx *Ctx[T]) GetText() T {
	return byteseq.B2Q[T](ctx.buf)
}

func (ctx *Ctx[T]) GetRunes() []rune {
	if len(ctx.bufR) == 0 {
		ctx.bufR = fastconv.AppendB2R(ctx.bufR[:0], ctx.buf)
	}
	return ctx.bufR
}

func (ctx *Ctx[T]) LimitScripts(list []Script) *Ctx[T] {
	ctx.bufSC = append(ctx.bufSC[:0], list...)
	ctx.BufSP = ctx.BufSP[:0]
	for i := 0; i < len(list); i++ {
		ctx.BufSP = append(ctx.BufSP, ScriptScore{
			Script: list[i],
			Score:  0,
		})
	}
	return ctx
}

func (ctx *Ctx[T]) GetScripts() []Script {
	return ctx.bufSC
}

func (ctx *Ctx[T]) GetError() error {
	return ctx.err
}

func (ctx *Ctx[T]) Reset() *Ctx[T] {
	ctx.Bitset.Reset()
	ctx.src = ctx.src[:0]
	ctx.buf = ctx.buf[:0]
	ctx.bufR = ctx.bufR[:0]
	ctx.BufSP = ctx.BufSP[:0]
	// ctx.BufLP = ctx.BufLP[:0]
	ctx.bufT.Reset()
	ctx.ResetCleaners()
	ctx.ResetModifiers()
	ctx.err = nil
	return ctx
}

func (ctx *Ctx[T]) chkSrcErr() bool {
	if len(ctx.buf) == 0 {
		ctx.err = ErrEmptyInput
		return true
	}
	return false
}
