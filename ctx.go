package nlp

import (
	"github.com/koykov/bitset"
	"github.com/koykov/fastconv"
)

const (
	flagClean = 0
	flagToken = 1
)

type Ctx[T Byteseq] struct {
	bitset.Bitset
	src  T
	srcB []byte

	cln  CleanerInterface[T]
	bufC []byte
	bufR []rune

	tkn  TokenizerInterface[T]
	bufT Tokens

	sd    ScriptDetectorInterface[T]
	bufSC []Script
	BufSP ScriptProba

	// BufLP LanguageProba
}

func NewCtx[T Byteseq]() *Ctx[T] {
	ctx := Ctx[T]{}
	return &ctx
}

func (ctx *Ctx[T]) SetText(text T) *Ctx[T] {
	ctx.SetBit(flagClean, false)
	ctx.SetBit(flagToken, false)
	ctx.src = text
	ctx.srcB = append(ctx.srcB[:0], q2b(ctx.src)...)
	return ctx
}

func (ctx Ctx[T]) GetText() T {
	return ctx.src
}

func (ctx Ctx[T]) GetRunes() []rune {
	if len(ctx.bufR) == 0 {
		ctx.bufR = fastconv.AppendB2R(ctx.bufR[:0], ctx.srcB)
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

func (ctx *Ctx[T]) Reset() *Ctx[T] {
	ctx.Bitset.Reset()
	ctx.src = ctx.src[:0]
	ctx.bufC = ctx.bufC[:0]
	ctx.bufR = ctx.bufR[:0]
	ctx.BufSP = ctx.BufSP[:0]
	// ctx.BufLP = ctx.BufLP[:0]
	ctx.bufT.Reset()
	return ctx
}
