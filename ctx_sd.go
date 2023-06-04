package nlp

func (ctx *Ctx[T]) WithScriptDetector(ds ScriptDetector[T]) *Ctx[T] {
	ctx.sd = ds
	ctx.SetBit(flagSD, false)
	return ctx
}

func (ctx *Ctx[T]) LimitScripts(list []Script) *Ctx[T] {
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

func (ctx *Ctx[T]) GetScriptsLimit() []Script {
	return ctx.bufSC
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

func (ctx *Ctx[T]) GetScript() Script {
	return ctx.scr
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

func (ctx *Ctx[T]) GetScriptProba() ScriptProba {
	return ctx.BufSP
}

func (ctx *Ctx[T]) ResetScriptDetector() *Ctx[T] {
	ctx.sd = nil
	return ctx
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
