package nlp

import "github.com/koykov/fastconv"

func (ctx *Ctx) WithCleaner(cln CleanerInterface) *Ctx {
	ctx.cln = cln
	ctx.SetBit(flagClean, false)
	return ctx
}

func (ctx *Ctx) Clean() *Ctx {
	if ctx.CheckBit(flagClean) {
		return ctx
	}
	defer ctx.SetBit(flagClean, true)

	ctx.bufR = ctx.chkCln().CleanString(ctx.bufR[:0], ctx.src)
	ctx.bufC = fastconv.AppendR2B(ctx.bufC[:0], ctx.bufR)

	return ctx
}

func (ctx *Ctx) CleanBytes(p []byte) []byte {
	return ctx.SetText(p).
		Clean().
		GetClean()
}

func (ctx *Ctx) CleanString(s string) string {
	return ctx.SetTextString(s).
		Clean().
		GetTextString()
}

func (ctx Ctx) GetClean() []byte {
	return ctx.bufC
}

func (ctx Ctx) GetCleanString() string {
	return fastconv.B2S(ctx.bufC)
}

func (ctx *Ctx) chkCln() CleanerInterface {
	if ctx.cln == nil {
		ctx.cln = NewCleaner()
	}
	return ctx.cln
}
