package nlp

import (
	"sort"
)

type DetectAlgo uint

const (
	DetectAlgoHalf DetectAlgo = iota
	DetectAlgoDistributed
	DetectAlgoFull
)

type Detector struct {
	algo DetectAlgo
}

func NewDetector() Detector {
	return Detector{algo: DetectAlgoFull}
}

func NewDetectorWithAlgo(algo DetectAlgo) Detector {
	return Detector{algo: algo}
}

func (d Detector) Detect(ctx *Ctx) (Script, error) {
	return d.DetectString(ctx)
}

func (d Detector) DetectProba(ctx *Ctx) (ScriptProba, error) {
	return d.DetectProbaString(ctx)
}

func (d Detector) DetectString(ctx *Ctx) (Script, error) {
	if err := d.dsProba(ctx); err != nil {
		return 0, err
	}
	var (
		mx float32
		mi int
	)
	_ = ctx.bufSP[len(ctx.bufSP)-1]
	for i := 0; i < len(ctx.bufSP); i++ {
		if score := ctx.bufSP[i].Score; score > mx {
			mx, mi = score, i
		}
	}
	return ctx.bufSP[mi].Script, nil
}

func (d Detector) DetectProbaString(ctx *Ctx) (ScriptProba, error) {
	if err := d.dsProba(ctx); err != nil {
		return nil, err
	}
	sort.Sort(&ctx.bufSP)
	return ctx.bufSP, nil
}

func (d Detector) dsProba(ctx *Ctx) error {
	runes := ctx.GetRunes()
	l := len(runes)
	if l == 0 {
		return ErrEmptyInput
	}
	if len(ctx.bufSP) == 0 {
		ctx.LimitScripts(ScriptsSupported())
	}
	s := l
	if d.algo == DetectAlgoHalf {
		l /= 2
	}
	if d.algo == DetectAlgoDistributed {
		s = distStep(l)
	}
	scripts := ctx.GetScripts()
	for i := 0; i < len(runes); i++ {
		for j := 0; j < len(scripts); j++ {
			if scripts[j].Evaluate(runes[i]) {
				//
			}
		}
	}
	ctx.R.Each(func(i int, r rune) {
		for j := 0; j < len(ctx.bufSP); j++ {
			if ctx.bufSP[j].Script.Evaluate(r) {
				ctx.bufSP[j].Score += 1
			}
		}
	}, s)
	for i := 0; i < len(ctx.bufSP); i++ {
		ctx.bufSP[i].Score /= float32(l)
	}
	return nil
}

func distStep(l int) int {
	if l < 8 {
		return 1
	}
	if l < 32 {
		return 2
	}
	if l < 128 {
		return 4
	}
	return 8
}
