package nlp

import (
	"testing"
)

type dsStage struct {
	key    string
	script Script
	err    error
}

var (
	dsStages = []dsStage{
		{key: "012345678987654321!", err: ErrEmptyInput},
		{key: "Hello, world!", script: ScriptLatin},
		{key: "Привет, мир!", script: ScriptCyrillic},
		{key: "ქართული ენა მსოფლიო", script: ScriptGeorgian},
		{key: "ككل حوالي 1.6، ومعظم الناس", script: ScriptArabic},
		{key: "የኢትዮጵያ ፌዴራላዊ ዴሞክራሲያዊሪፐብሊክ", script: ScriptEthiopic},
		{key: "היסטוריה והתפתחות של האלפבית העברי", script: ScriptHebrew},
		{key: "県見夜上温国阪題富販", script: ScriptHan},
		{key: "আমি ভালো আছি, ধন্যবাদ!", script: ScriptBengali},
		{key: "Английское слово fluctuate означает \"неустойчивый\"", script: ScriptCyrillic},
		{key: "Russian word собственник means proprietor", script: ScriptLatin},
	}
)

func TestDetectScript(t *testing.T) {
	for _, stage := range dsStages {
		t.Run(stage.key, func(t *testing.T) {
			ctx := NewCtx()
			ctx.SetDetectScriptAlgo(DetectScriptFull)
			script, err := DetectScriptString(ctx, stage.key)
			if err != nil {
				if err != stage.err {
					t.Error(err)
				}
				return
			}
			if script != stage.script {
				t.Errorf("detect script failed: need %d, got %d", stage.script, script)
			}
		})
	}
}
