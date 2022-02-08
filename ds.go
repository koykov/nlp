package nlp

import (
	"sort"

	"github.com/koykov/fastconv"
)

type DetectScriptAlgo uint

const (
	DetectScriptHalf DetectScriptAlgo = iota
	DetectScriptDistributed
	DetectScriptFull
)

type ScriptScore struct {
	Script Script
	Score  float32
}

type ScriptProba []ScriptScore

func (s ScriptProba) Len() int {
	return len(s)
}

func (s ScriptProba) Less(i, j int) bool {
	return s[i].Score > s[j].Score
}

func (s *ScriptProba) Swap(i, j int) {
	(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
}

func DetectScript(ctx *Ctx, text []byte) (Script, error) {
	return DetectScriptString(ctx, fastconv.B2S(text))
}

func DetectScriptString(ctx *Ctx, text string) (Script, error) {
	if err := dsProba(ctx, text); err != nil {
		return 0, err
	}
	var (
		max  float32
		maxi int
	)
	for i := 0; i < len(ctx.bufSP); i++ {
		if score := ctx.bufSP[i].Score; score > max {
			max, maxi = score, i
		}
	}
	return ctx.bufSP[maxi].Script, nil
}

func DetectScriptProba(ctx *Ctx, text []byte) (ScriptProba, error) {
	return DetectScriptStringProba(ctx, fastconv.B2S(text))
}

func DetectScriptStringProba(ctx *Ctx, text string) (ScriptProba, error) {
	if err := dsProba(ctx, text); err != nil {
		return nil, err
	}
	sort.Sort(&ctx.bufSP)
	return ctx.bufSP, nil
}

func dsProba(ctx *Ctx, text string) error {
	for _, r := range text {
		if !mustSkip(r) {
			ctx.bufR = append(ctx.bufR, r)
		}
	}
	if len(ctx.bufR) == 0 {
		return ErrEmptyInput
	}
	if len(ctx.bufSP) == 0 {
		ctx.LimitScripts(ScriptsSupported())
	}
	l, s := len(ctx.bufR), 1
	if ctx.dsa == DetectScriptHalf {
		l /= 2
	}
	if ctx.dsa == DetectScriptDistributed {
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
