package nlp

func (ctx *Ctx[T]) WithScriptDetector(ds ScriptDetector[T]) *Ctx[T] {
	ctx.sd = ds
	ctx.SetBit(flagSD, false)
	return ctx
}

func (ctx *Ctx[T]) DetectScript() *Ctx[T] {
	if ctx.chkSrcErr() {
		return ctx
	}
	if ctx.CheckBit(flagSD) {
		return ctx
	}
	defer ctx.SetBit(flagSD, true)
	ctx.scr, ctx.err = ctx.chkSD().Detect(ctx)
	return ctx
}

func (ctx *Ctx[T]) DetectScriptT(text T) (Script, error) {
	ctx.SetText(text).
		DetectScript()
	return ctx.scr, ctx.err
}

func (ctx *Ctx[T]) DetectScriptProba() *Ctx[T] {
	if ctx.chkSrcErr() {
		return ctx
	}
	if ctx.CheckBit(flagSDP) {
		return ctx
	}
	defer ctx.SetBit(flagSDP, true)
	ctx.BufSP, ctx.err = ctx.chkSD().DetectProba(ctx)
	return ctx
}

func (ctx *Ctx[T]) DetectScriptProbaT(text T) (ScriptProba, error) {
	ctx.SetText(text).
		DetectScriptProba()
	return ctx.BufSP, ctx.err
}

func (ctx *Ctx[T]) chkSD() ScriptDetector[T] {
	if ctx.sd == nil {
		ctx.sd = NewUnicodeScriptDetector[T]()
	}
	if len(ctx.bufSC) == 0 {
		ctx.LimitScripts(ScriptsSupported())
	}
	return ctx.sd
}
