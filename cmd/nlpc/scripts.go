package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"go/format"
	"io/ioutil"
	"strings"
	"unicode"
)

type scriptsModule struct{}

type scriptsTuple struct {
	Name      string   `json:"name"`
	Weight    uint     `json:"weight"`
	Languages []string `json:"languages"`
}

func (m scriptsModule) Validate(input, _ string) error {
	if len(input) == 0 {
		return fmt.Errorf("param -input is required")
	}
	return nil
}

func (m scriptsModule) Compile(w moduleWriter, input, target string) (err error) {
	if len(target) == 0 {
		target = "scripts_repo.go"
	}

	var body []byte
	if body, err = ioutil.ReadFile(input); err != nil {
		return
	}
	if len(body) == 0 {
		err = errors.New("nothing to parse")
		return
	}
	var tuples []scriptsTuple
	if err = json.Unmarshal(body, &tuples); err != nil {
		return
	}

	_, _ = w.WriteString("import \"unicode\"\n\nconst (\n")
	for i := 0; i < len(tuples); i++ {
		t := &tuples[i]
		t.Name = strings.ReplaceAll(t.Name, " ", "_")
		if _, ok := unicode.Scripts[t.Name]; !ok {
			err = fmt.Errorf("unknown script name: %s", t.Name)
			return
		}
		_, _ = w.WriteString("Script")
		_, _ = w.WriteString(t.Name)
		if i == 0 {
			_, _ = w.WriteString(" Script = iota")
		}
		_ = w.WriteByte('\n')
	}
	_, _ = w.WriteString(")\n\n")

	_, _ = w.WriteString("var (\n__sre_buf = []SRE{\n")
	for i := 0; i < len(tuples); i++ {
		t := &tuples[i]
		_, _ = w.WriteString("SRE{Evaluate: __sreEval")
		_, _ = w.WriteString(t.Name)
		_, _ = w.WriteString("},\n")
	}
	_, _ = w.WriteString("}\n")
	_, _ = w.WriteString("__st_buf = []*unicode.RangeTable{\n")
	for i := 0; i < len(tuples); i++ {
		t := &tuples[i]
		_, _ = w.WriteString("unicode.")
		_, _ = w.WriteString(t.Name)
		_, _ = w.WriteString(",\n")
	}
	_, _ = w.WriteString("}\n)\n")

	for i := 0; i < len(tuples); i++ {
		t := &tuples[i]
		_, _ = w.WriteString("func __sreEval")
		_, _ = w.WriteString(t.Name)
		_, _ = w.WriteString("(r rune) bool {\n")

		rt := unicode.Scripts[t.Name]
		l16, l32 := len(rt.R16), len(rt.R32)
		if l16 > 0 {
			rl16 := rt.R16[l16-1]
			_, _ = w.WriteString("if r <= rune(0x" + fmt.Sprintf("%04x", rl16.Hi) + ") {\n")
			_, _ = w.WriteString("l16 := len(unicode." + t.Name + ".R16)\n")
			_, _ = w.WriteString("r16 := uint16(r)\n")
			_, _ = w.WriteString("if l16 <= sreLinearMax || r16 <= unicode.MaxLatin1 {\n")
			for i := 0; i < l16; i++ {
				rt16 := rt.R16[i]
				_, _ = w.WriteString(fmt.Sprintf("if r16 < 0x%04x {return false}\n", rt16.Lo))
				_, _ = w.WriteString(fmt.Sprintf("if r16 <= 0x%04x {\n", rt16.Hi))
				if rt16.Stride == 1 {
					_, _ = w.WriteString("return true\n")
				} else {
					_, _ = w.WriteString(fmt.Sprintf("return (r16-0x%04x)/0x%04x==0\n", rt16.Lo, rt16.Stride))
				}
				_, _ = w.WriteString("}\n")
			}
			_, _ = w.WriteString("}\n")
			_, _ = w.WriteString("return sreEvalBinary16(unicode." + t.Name + ".R16, r16)\n")
			_, _ = w.WriteString("}\n")
		}
		if l32 > 0 {
			_, _ = w.WriteString("if r >= rune(0x" + fmt.Sprintf("%05x", rt.R32[0].Lo) + ") {\n")
			_, _ = w.WriteString("l32 := len(unicode." + t.Name + ".R32)\n")
			_, _ = w.WriteString("r32 := uint32(r)\n")
			_, _ = w.WriteString("if l32 <= sreLinearMax {\n")
			for i := 0; i < l32; i++ {
				rt32 := rt.R32[i]
				_, _ = w.WriteString(fmt.Sprintf("if r32 < 0x%05x {return false}\n", rt32.Lo))
				_, _ = w.WriteString(fmt.Sprintf("if r32 <= 0x%05x {\n", rt32.Hi))
				if rt32.Stride == 1 {
					_, _ = w.WriteString("return true\n")
				} else {
					_, _ = w.WriteString(fmt.Sprintf("return (r32-0x%05x)/0x%05x==0\n", rt32.Lo, rt32.Stride))
				}
				_, _ = w.WriteString("}\n")
			}
			_, _ = w.WriteString("}\n")
			_, _ = w.WriteString("return sreEvalBinary32(unicode." + t.Name + ".R32, r32)\n")
			_, _ = w.WriteString("}\n")
		}

		if l32 > 0 {
			//
		}

		_, _ = w.WriteString("return false\n}\n\n")
	}

	source := w.Bytes()
	var fmtSource []byte
	if fmtSource, err = format.Source(source); err != nil {
		return
	}

	err = ioutil.WriteFile(target, fmtSource, 0644)

	return
}
