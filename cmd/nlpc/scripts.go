package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"go/format"
	"io/ioutil"
	"sort"
	"strings"
)

type scriptsModule struct{}

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
	var tuples []languagesTuple
	if err = json.Unmarshal(body, &tuples); err != nil {
		return
	}

	uniq := make(map[string]struct{})
	var list []string
	for i := 0; i < len(tuples); i++ {
		tuple := &tuples[i]
		for j := 0; j < len(tuple.Scripts); j++ {
			script := tuple.Scripts[j]
			script = strings.ReplaceAll(script, " ", "_")
			if _, ok := uniq[script]; ok {
				continue
			}
			uniq[script] = struct{}{}
			list = append(list, script)
		}
	}
	if len(list) == 0 {
		return
	}
	sort.Strings(list)

	_, _ = w.WriteString("import \"unicode\"\n\nconst (\n")
	for i := 0; i < len(list); i++ {
		_, _ = w.WriteString("Script")
		_, _ = w.WriteString(list[i])
		if i == 0 {
			_, _ = w.WriteString(" Script = iota")
		}
		_ = w.WriteByte('\n')
	}
	_, _ = w.WriteString(")\n\n")

	_, _ = w.WriteString("var (\n__st_buf = []*unicode.RangeTable{\n")
	for i := 0; i < len(list); i++ {
		_, _ = w.WriteString("unicode.")
		_, _ = w.WriteString(list[i])
		_, _ = w.WriteString(",\n")
	}
	_, _ = w.WriteString("}\n)\n")

	source := w.Bytes()
	var fmtSource []byte
	if fmtSource, err = format.Source(source); err != nil {
		return
	}

	err = ioutil.WriteFile(target, fmtSource, 0644)

	return
}
