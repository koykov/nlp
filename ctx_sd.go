package nlp

import "github.com/koykov/fastconv"

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
	if len(text) > 0 {
		ctx.SetTextString(text)
	}
	if len(ctx.bufSC) == 0 {
		ctx.LimitScripts(ScriptsSupported())
	}
	return ctx.sd.DetectString(ctx)
}

func (ctx *Ctx) DetectScriptProba(text []byte) (ScriptProba, error) {
	return ctx.DetectScriptStringProba(fastconv.B2S(text))
}

func (ctx *Ctx) DetectScriptStringProba(text string) (ScriptProba, error) {
	if ctx.sd == nil {
		return nil, ErrNoScriptDetector
	}
	if len(text) > 0 {
		ctx.SetTextString(text)
	}
	if len(ctx.bufSC) == 0 {
		ctx.LimitScripts(ScriptsSupported())
	}
	return ctx.sd.DetectProbaString(ctx)
}
