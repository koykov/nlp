package nlp

func (ctx *Ctx[T]) WithScriptDetector(ds ScriptDetector[T]) *Ctx[T] {
	ctx.sd = ds
	return ctx
}

func (ctx *Ctx[T]) DetectScript(text T) (Script, error) {
	if len(text) > 0 {
		ctx.SetText(text)
	}
	return ctx.chkSD().Detect(ctx)
}

func (ctx *Ctx[T]) DetectScriptProba(text T) (ScriptProba, error) {
	if len(text) > 0 {
		ctx.SetText(text)
	}
	return ctx.chkSD().DetectProba(ctx)
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
