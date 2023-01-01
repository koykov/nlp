package nlp

type Ctx struct {
	sd ScriptDetector

	bufR  []rune
	bufOR []rune
	bufSP ScriptProba
	bufLP LanguageProba
}

func NewCtx() *Ctx {
	ctx := Ctx{}
	return &ctx
}

func (ctx *Ctx) bufferize(text string) {
	ctx.bufOR, ctx.bufR = ctx.bufOR[:0], ctx.bufR[:0]
	for _, r := range text {
		ctx.bufOR = append(ctx.bufOR, r)
		if !mustSkip(r) {
			ctx.bufR = append(ctx.bufR, r)
		}
	}
}

func (ctx *Ctx) LimitScripts(list []Script) {
	l := len(list)
	if l == 0 {
		return
	}
	_ = list[l-1]
	ctx.bufSP = ctx.bufSP[:0]
	for i := 0; i < l; i++ {
		ctx.bufSP = append(ctx.bufSP, ScriptScore{
			Script: list[i],
			Score:  0,
		})
	}
}

func (ctx *Ctx) Reset() {
	ctx.bufR = ctx.bufR[:0]
	ctx.bufOR = ctx.bufOR[:0]
	ctx.bufSP = ctx.bufSP[:0]
	ctx.bufLP = ctx.bufLP[:0]
}
