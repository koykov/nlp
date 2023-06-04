package nlp

import (
	"github.com/koykov/byteseq"
	"github.com/koykov/fastconv"
)

func (ctx *Ctx[T]) WithCleaner(cln Cleaner[T]) *Ctx[T] {
	ctx.cln = append(ctx.cln, cln)
	ctx.SetBit(flagClean, false)
	return ctx
}

func (ctx *Ctx[T]) Clean() *Ctx[T] {
	if ctx.chkSrcErr() {
		return ctx
	}
	if ctx.CheckBit(flagClean) {
		return ctx
	}
	defer ctx.SetBit(flagClean, true)

	if len(ctx.cln) == 0 {
		ctx.cln = append(ctx.cln, NewUnicodeCleaner[T](DefaultCleanMask))
	}
	for i := 0; i < len(ctx.cln); i++ {
		ctx.bufR = ctx.cln[i].AppendClean(ctx.bufR[:0], byteseq.B2Q[T](ctx.buf))
		ctx.buf = fastconv.AppendR2B(ctx.buf[:0], ctx.bufR)
	}

	return ctx
}

func (ctx *Ctx[T]) ResetCleaners() *Ctx[T] {
	ctx.cln = ctx.cln[:0]
	return ctx
}
