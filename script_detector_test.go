package nlp

import (
	"testing"
)

type dsStage struct {
	key, text string
	script    Script
	err       error
}

var (
	dsStages = []dsStage{
		{key: "no script", text: "012345678987654321!", err: ErrEmptyInput},
		{key: "pure latin", text: "Hello, world!", script: ScriptLatin},
		{key: "pure cyrillic", text: "Привет, мир!", script: ScriptCyrillic},
		{key: "pure georgian", text: "ქართული ენა მსოფლიო", script: ScriptGeorgian},
		{key: "pure arabic", text: "ككل حوالي 1.6، ومعظم الناس", script: ScriptArabic},
		{key: "pure ethiopic", text: "የኢትዮጵያ ፌዴራላዊ ዴሞክራሲያዊሪፐብሊክ", script: ScriptEthiopic},
		{key: "pure hebrew", text: "היסטוריה והתפתחות של האלפבית העברי", script: ScriptHebrew},
		{key: "pure han", text: "県見夜上温国阪題富販", script: ScriptHan},
		{key: "pure bengali", text: "আমি ভালো আছি, ধন্যবাদ!", script: ScriptBengali},
		{key: "mixed cyrillic and latin", text: "Английское слово fluctuate означает \"неустойчивый\"", script: ScriptCyrillic},
		{key: "mixed latin and cyrillic", text: "Russian word собственник means proprietor", script: ScriptLatin},
	}
)

func TestDetectScript(t *testing.T) {
	for _, stage := range dsStages {
		t.Run(stage.key, func(t *testing.T) {
			ctx := NewCtx[string]().WithScriptDetector(NewScriptDetector[string]())
			script, err := ctx.DetectScript(stage.text)
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
				ctx := AcquireCtx[string]()
				script, err := ctx.DetectScript(stage.text)
				if err != nil {
					if err != stage.err {
						b.Error(err)
					}
				} else if script != stage.script {
					b.Errorf("detect script failed: need %d, got %d", stage.script, script)
				}
				ReleaseCtx(ctx)
			}
		})
	}
}

func BenchmarkDetectScriptProba(b *testing.B) {
	for _, stage := range dsStages {
		b.Run(stage.key, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				ctx := AcquireCtx[string]().WithScriptDetector(NewScriptDetectorWithAlgo[string](ScriptDetectAlgoDistributed))
				proba, err := ctx.DetectScriptProba(stage.text)
				if err != nil {
					if err != stage.err {
						b.Error(err)
					}
				} else if script := proba[0].Script; script != stage.script {
					b.Errorf("detect script failed: need %d, got %d", stage.script, script)
				}
				ReleaseCtx(ctx)
			}
		})
	}
}
