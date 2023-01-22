package nlp

import "github.com/koykov/fastconv"

func (ctx *Ctx[T]) WithCleaner(cln CleanerInterface[T]) *Ctx[T] {
	ctx.cln = cln
	ctx.SetBit(flagClean, false)
	return ctx
}

func (ctx *Ctx[T]) Clean() *Ctx[T] {
	if ctx.CheckBit(flagClean) {
		return ctx
	}
	defer ctx.SetBit(flagClean, true)

	ctx.bufR = ctx.chkCln().Clean(ctx.bufR[:0], ctx.src)
	ctx.buf = fastconv.AppendR2B(ctx.buf[:0], ctx.bufR)

	return ctx
}

func (ctx *Ctx[T]) CleanString(x T) T {
	return ctx.SetText(x).
		Clean().
		GetClean()
}

func (ctx Ctx[T]) GetClean() T {
	return T(ctx.buf)
}

func (ctx *Ctx[T]) chkCln() CleanerInterface[T] {
	if ctx.cln == nil {
		cln := NewCleaner[T]()
		ctx.cln = cln
	}
	return ctx.cln
}
