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
	if len(text) > 0 {
		ctx.SetTextString(text)
	}
	return ctx.chkSD().Detect(ctx)
}

func (ctx *Ctx) DetectScriptProba(text []byte) (ScriptProba, error) {
	return ctx.DetectScriptStringProba(fastconv.B2S(text))
}

func (ctx *Ctx) DetectScriptStringProba(text string) (ScriptProba, error) {
	if len(text) > 0 {
		ctx.SetTextString(text)
	}
	return ctx.chkSD().DetectProba(ctx)
}

func (ctx *Ctx) chkSD() ScriptDetectorInterface {
	if ctx.sd == nil {
		ctx.sd = NewScriptDetector()
	}
	if len(ctx.bufSC) == 0 {
		ctx.LimitScripts(ScriptsSupported())
	}
	return ctx.sd
}
