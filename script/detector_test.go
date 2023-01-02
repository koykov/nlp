package script

import (
	"testing"

	"github.com/koykov/nlp"
)

type dsStage struct {
	key, text string
	script    nlp.Script
	err       error
}

var (
	dsStages = []dsStage{
		{key: "no script", text: "012345678987654321!", err: nlp.ErrEmptyInput},
		{key: "pure latin", text: "Hello, world!", script: nlp.ScriptLatin},
		{key: "pure cyrillic", text: "Привет, мир!", script: nlp.ScriptCyrillic},
		{key: "pure georgian", text: "ქართული ენა მსოფლიო", script: nlp.ScriptGeorgian},
		{key: "pure arabic", text: "ككل حوالي 1.6، ومعظم الناس", script: nlp.ScriptArabic},
		{key: "pure ethiopic", text: "የኢትዮጵያ ፌዴራላዊ ዴሞክራሲያዊሪፐብሊክ", script: nlp.ScriptEthiopic},
		{key: "pure hebrew", text: "היסטוריה והתפתחות של האלפבית העברי", script: nlp.ScriptHebrew},
		{key: "pure han", text: "県見夜上温国阪題富販", script: nlp.ScriptHan},
		{key: "pure bengali", text: "আমি ভালো আছি, ধন্যবাদ!", script: nlp.ScriptBengali},
		{key: "mixed cyrillic and latin", text: "Английское слово fluctuate означает \"неустойчивый\"", script: nlp.ScriptCyrillic},
		{key: "mixed latin and cyrillic", text: "Russian word собственник means proprietor", script: nlp.ScriptLatin},
	}
)

func TestDetectScript(t *testing.T) {
	for _, stage := range dsStages {
		t.Run(stage.key, func(t *testing.T) {
			ctx := nlp.NewCtx().WithScriptDetector(NewDetector())
			script, err := ctx.DetectScriptString(stage.text)
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

func BenchmarkDetectScript(b *testing.B) {
	for _, stage := range dsStages {
		b.Run(stage.key, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				ctx := nlp.AcquireCtx().WithScriptDetector(NewDetectorWithAlgo(DetectAlgoDistributed))
				script, err := ctx.DetectScriptString(stage.text)
				if err != nil {
					if err != stage.err {
						b.Error(err)
					}
				} else if script != stage.script {
					b.Errorf("detect script failed: need %d, got %d", stage.script, script)
				}
				nlp.ReleaseCtx(ctx)
			}
		})
	}
}

func BenchmarkDetectScriptProba(b *testing.B) {
	for _, stage := range dsStages {
		b.Run(stage.key, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				ctx := nlp.AcquireCtx().WithScriptDetector(NewDetectorWithAlgo(DetectAlgoDistributed))
				proba, err := ctx.DetectScriptStringProba(stage.text)
				if err != nil {
					if err != stage.err {
						b.Error(err)
					}
				} else if script := proba[0].Script; script != stage.script {
					b.Errorf("detect script failed: need %d, got %d", stage.script, script)
				}
				nlp.ReleaseCtx(ctx)
			}
		})
	}
}
