package nlp

import "github.com/koykov/fastconv"

type ScriptScore struct {
	Script Script
	Score  float32
}

type ScriptProba []ScriptScore

func (s ScriptProba) Len() int {
	return len(s)
}

func (s ScriptProba) Less(i, j int) bool {
	return s[i].Score < s[j].Score
}

func (s *ScriptProba) Swap(i, j int) {
	(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
}

func DetectScript(ctx *Ctx, text []byte) (Script, error) {
	// ...
	return 0, nil
}

func DetectScriptString(ctx *Ctx, text string) (Script, error) {
	return DetectScript(ctx, fastconv.S2B(text))
}

func DetectScriptProba(ctx *Ctx, text []byte) (ScriptProba, error) {
	// ...
	return nil, nil
}

func DetectScriptStringProba(ctx *Ctx, text string) (ScriptProba, error) {
	return DetectScriptProba(ctx, fastconv.S2B(text))
}
