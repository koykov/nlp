package nlp

func (ctx *Ctx[T]) WithModifier(mod Modifier[T]) *Ctx[T] {
	ctx.mod = mod
	ctx.SetBit(flagMod, false)
	return ctx
}

func (ctx *Ctx[T]) Modify() *Ctx[T] {
	if ctx.chkSrcErr() {
		return ctx
	}
	if ctx.CheckBit(flagMod) {
		return ctx
	}
	defer ctx.SetBit(flagMod, true)
	ctx.bufT = ctx.chkTkn().AppendTokenize(ctx.bufT, T(ctx.buf))
	return ctx
}

func (ctx *Ctx[T]) ModifyT(t T) *Ctx[T] {
	return ctx.SetText(t).
		Modify()
}

func (ctx *Ctx[T]) chkMod() Modifier[T] {
	if ctx.mod == nil {
		mod := DummyModifier[T]{}
		ctx.WithModifier(&mod)
	}
	return ctx.mod
}
