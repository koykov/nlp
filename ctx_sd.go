package nlp

import "github.com/koykov/fastconv"

func (ctx *Ctx) WithScriptDetector(ds ScriptDetectorInterface) *Ctx {
	ctx.sd = ds
	return ctx
}

func (ctx *Ctx) DetectScript(text []byte) (Script, error) {
	return ctx.DetectScriptString(fastconv.B2S(text))
}

func (ctx *Ctx) DetectScriptString(text string) (Script, error) {
	sd := ctx.sd
	if sd == nil {
		sd = NewScriptDetector()
	}
	if len(text) > 0 {
		ctx.SetTextString(text)
	}
	if len(ctx.bufSC) == 0 {
		ctx.LimitScripts(ScriptsSupported())
	}
	return sd.Detect(ctx)
}

func (ctx *Ctx) DetectScriptProba(text []byte) (ScriptProba, error) {
	return ctx.DetectScriptStringProba(fastconv.B2S(text))
}

func (ctx *Ctx) DetectScriptStringProba(text string) (ScriptProba, error) {
	sd := ctx.sd
	if sd == nil {
		sd = NewScriptDetector()
	}
	if len(text) > 0 {
		ctx.SetTextString(text)
	}
	if len(ctx.bufSC) == 0 {
		ctx.LimitScripts(ScriptsSupported())
	}
	return sd.DetectProba(ctx)
}
