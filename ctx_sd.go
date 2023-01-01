package nlp

import "github.com/koykov/fastconv"

type ScriptDetector interface {
	Detect(ctx *Ctx, text []byte) (Script, error)
	DetectProba(ctx *Ctx, text []byte) (ScriptProba, error)
	DetectString(ctx *Ctx, text string) (Script, error)
	DetectProbaString(ctx *Ctx, text string) (ScriptProba, error)
}

func (ctx *Ctx) WithScriptDetector(ds ScriptDetector) *Ctx {
	ctx.sd = ds
	return ctx
}

func (ctx *Ctx) DetectScript(text []byte) (Script, error) {
	return ctx.DetectScriptString(fastconv.B2S(text))
}

func (ctx *Ctx) DetectScriptString(text string) (Script, error) {
	if ctx.sd == nil {
		return 0, ErrNoScriptDetector
	}
	return ctx.sd.DetectString(ctx, text)
}

func (ctx *Ctx) DetectScriptProba(text []byte) (ScriptProba, error) {
	return ctx.DetectScriptStringProba(fastconv.B2S(text))
}

func (ctx *Ctx) DetectScriptStringProba(text string) (ScriptProba, error) {
	if ctx.sd == nil {
		return nil, ErrNoScriptDetector
	}
	return ctx.sd.DetectProbaString(ctx, text)
}
