package nlp

import "github.com/koykov/fastconv"

func (ctx *Ctx) WithCleaner(cln CleanerInterface) *Ctx {
	ctx.cln = cln
	ctx.SetBit(flagClean, false)
	return ctx
}

func (ctx *Ctx) Clean(p []byte) []byte {
	return fastconv.S2B(ctx.CleanString(fastconv.B2S(p)))
}

func (ctx *Ctx) CleanString(s string) string {
	if ctx.CheckBit(flagClean) {
		return fastconv.B2S(ctx.bufC)
	}
	defer ctx.SetBit(flagClean, true)
	ctx.bufR = ctx.chkCln().CleanString(ctx.bufR[:0], s)
	ctx.bufC = fastconv.AppendR2B(ctx.bufC[:0], ctx.bufR)
	return fastconv.B2S(ctx.bufC)
}

func (ctx *Ctx) chkCln() CleanerInterface {
	if ctx.cln == nil {
		ctx.cln = NewCleaner()
	}
	return ctx.cln
}
