package nlp

type Ctx struct {
	dsa DetectScriptAlgo

	bufR  []rune
	bufSP ScriptProba
	bufLP LanguageProba
}

func NewCtx() *Ctx {
	ctx := Ctx{}
	return &ctx
}

func (c *Ctx) SetDetectScriptAlgo(algo DetectScriptAlgo) {
	c.dsa = algo
}

func (c *Ctx) LimitScripts(list []Script) {
	l := len(list)
	if l == 0 {
		return
	}
	if len(c.bufSP) > 0 {
		c.bufSP = c.bufSP[:0]
	}
	_ = list[l-1]
	for i := 0; i < l; i++ {
		c.bufSP = append(c.bufSP, ScriptScore{
			Script: list[i],
			Score:  0,
		})
	}
}

func (c *Ctx) Reset() {
	c.bufR = c.bufR[:0]
	c.bufSP = c.bufSP[:0]
	c.bufLP = c.bufLP[:0]
}
