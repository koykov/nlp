package nlp

import "github.com/koykov/fastconv"

func (ctx *Ctx[T]) WithCleaner(cln Cleaner[T]) *Ctx[T] {
	ctx.cln = cln
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

	ctx.bufR = ctx.chkCln().AppendClean(ctx.bufR[:0], ctx.src)
	ctx.buf = fastconv.AppendR2B(ctx.buf[:0], ctx.bufR)

	return ctx
}

func (ctx *Ctx[T]) CleanT(t T) T {
	return ctx.SetText(t).
		Clean().
		GetClean()
}

func (ctx Ctx[T]) GetClean() T {
	return T(ctx.buf)
}

func (ctx *Ctx[T]) chkCln() Cleaner[T] {
	if ctx.cln == nil {
		cln := NewUnicodeCleaner[T](DefaultCleanMask)
		ctx.cln = cln
	}
	return ctx.cln
}
