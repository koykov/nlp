package nlp

import (
	"sort"

	"github.com/koykov/fastconv"
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

func (d Detector) Detect(ctx *Ctx, text []byte) (Script, error) {
	return d.DetectString(ctx, fastconv.B2S(text))
}

func (d Detector) DetectProba(ctx *Ctx, text []byte) (ScriptProba, error) {
	return d.DetectProbaString(ctx, fastconv.B2S(text))
}

func (d Detector) DetectString(ctx *Ctx, text string) (Script, error) {
	if err := d.dsProba(ctx, text); err != nil {
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

func (d Detector) DetectProbaString(ctx *Ctx, text string) (ScriptProba, error) {
	if err := d.dsProba(ctx, text); err != nil {
		return nil, err
	}
	sort.Sort(&ctx.bufSP)
	return ctx.bufSP, nil
}

func (d Detector) dsProba(ctx *Ctx, text string) error {
	ctx.bufferize(text)
	if len(ctx.bufR) == 0 {
		return ErrEmptyInput
	}
	if len(ctx.bufSP) == 0 {
		ctx.LimitScripts(ScriptsSupported())
	}
	l, s := len(ctx.bufR), 1
	if d.algo == DetectAlgoHalf {
		l /= 2
	}
	if d.algo == DetectAlgoDistributed {
		s = dsStep(l)
	}
	_, _ = ctx.bufR[l-1], ctx.bufSP[len(ctx.bufSP)-1]
	for i := 0; i < l; i += s {
		for j := 0; j < len(ctx.bufSP); j++ {
			if ctx.bufSP[j].Script.Evaluate(ctx.bufR[i]) {
				ctx.bufSP[j].Score += 1
			}
		}
	}
	for i := 0; i < len(ctx.bufSP); i++ {
		ctx.bufSP[i].Score /= float32(l)
	}
	return nil
}

func dsStep(l int) int {
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
