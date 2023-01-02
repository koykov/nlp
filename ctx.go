package nlp

import "github.com/koykov/fastconv"

type Ctx struct {
	t string

	bufR []rune

	sd    ScriptDetector
	bufSC []Script
	BufSP ScriptProba

	BufLP LanguageProba
}

func NewCtx() *Ctx {
	ctx := Ctx{}
	return &ctx
}

func (ctx *Ctx) SetText(text []byte) *Ctx {
	return ctx.SetTextString(fastconv.B2S(text))
}

func (ctx Ctx) GetText() []byte {
	return fastconv.S2B(ctx.t)
}

func (ctx *Ctx) SetTextString(text string) *Ctx {
	ctx.t = text
	ctx.bufR = ctx.bufR[:0]
	for _, r := range text {
		if !mustSkip(r) {
			ctx.bufR = append(ctx.bufR, r)
		}
	}
	return ctx
}

func (ctx Ctx) GetTextString() string {
	return ctx.t
}

func (ctx Ctx) GetRunes() []rune {
	return ctx.bufR
}

func (ctx *Ctx) LimitScripts(list []Script) *Ctx {
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

func (ctx *Ctx) GetScripts() []Script {
	return ctx.bufSC
}

func (ctx *Ctx) Reset() *Ctx {
	ctx.t = ""
	ctx.bufR = ctx.bufR[:0]
	ctx.BufSP = ctx.BufSP[:0]
	ctx.BufLP = ctx.BufLP[:0]
	return ctx
}
