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
	_ = ctx.BufSP[len(ctx.BufSP)-1]
	for i := 0; i < len(ctx.BufSP); i++ {
		if score := ctx.BufSP[i].Score; score > mx {
			mx, mi = score, i
		}
	}
	return ctx.BufSP[mi].Script, nil
}

func (d Detector) DetectProbaString(ctx *Ctx) (ScriptProba, error) {
	if err := d.dsProba(ctx); err != nil {
		return nil, err
	}
	sort.Sort(&ctx.BufSP)
	return ctx.BufSP, nil
}

func (d Detector) dsProba(ctx *Ctx) error {
	runes := ctx.GetRunes()
	l := len(runes)
	if l == 0 {
		return ErrEmptyInput
	}
	s := 1
	if d.algo == DetectAlgoHalf {
		l /= 2
	}
	if d.algo == DetectAlgoDistributed {
		s = distStep(l)
	}
	scripts := ctx.GetScripts()
	sl := len(scripts)
	if sl == 0 {
		return nil
	}
	ctx.BufSP = ctx.BufSP[:0]
	_ = scripts[sl-1]
	for i := 0; i < len(scripts); i++ {
		ctx.BufSP = append(ctx.BufSP, ScriptScore{Script: scripts[i]})
	}
	_ = runes[l-1]
	for i := 0; i < len(runes); i += s {
		for j := 0; j < len(scripts); j++ {
			if scripts[j].Evaluate(runes[i]) {
				ctx.BufSP[j].Score += 1
			}
		}
	}
	for i := 0; i < len(ctx.BufSP); i++ {
		ctx.BufSP[i].Score /= float32(l)
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
