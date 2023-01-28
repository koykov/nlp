package nlp

import (
	"sort"

	"github.com/koykov/byteseq"
)

type ScriptDetectAlgo uint

const (
	ScriptDetectAlgoHalf ScriptDetectAlgo = iota
	ScriptDetectAlgoDistributed
	ScriptDetectAlgoFull
)

// UnicodeScriptDetector is a builtin detector of writing scripts.
type UnicodeScriptDetector[T byteseq.Byteseq] struct {
	algo ScriptDetectAlgo
}

func NewUnicodeScriptDetector[T byteseq.Byteseq]() UnicodeScriptDetector[T] {
	return UnicodeScriptDetector[T]{algo: ScriptDetectAlgoFull}
}

func NewUnicodeScriptDetectorWithAlgo[T byteseq.Byteseq](algo ScriptDetectAlgo) UnicodeScriptDetector[T] {
	return UnicodeScriptDetector[T]{algo: algo}
}

func (d UnicodeScriptDetector[T]) Detect(ctx *Ctx[T]) (Script, error) {
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

func (d UnicodeScriptDetector[T]) DetectProba(ctx *Ctx[T]) (ScriptProba, error) {
	if err := d.dsProba(ctx); err != nil {
		return nil, err
	}
	sort.Sort(&ctx.BufSP)
	return ctx.BufSP, nil
}

func (d UnicodeScriptDetector[T]) dsProba(ctx *Ctx[T]) error {
	runes := ctx.GetRunes()
	l := len(runes)
	if l == 0 {
		return ErrEmptyInput
	}

	s := 1
	if d.algo == ScriptDetectAlgoHalf {
		l /= 2
	}
	if d.algo == ScriptDetectAlgoDistributed {
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
