package nlp

type Ctx struct {
	bufSP ScriptProba
	bufLP LanguageProba
}

func NewCtx() *Ctx {
	ctx := Ctx{}
	return &ctx
}

func (c *Ctx) Reset() {
	c.bufSP = c.bufSP[:0]
	c.bufLP = c.bufLP[:0]
}
