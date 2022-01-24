package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"go/format"
	"io/ioutil"
	"strings"
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
		target = "scripts_repo1.go"
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
		tuples[i].Name = strings.ReplaceAll(tuples[i].Name, " ", "_")
		_, _ = w.WriteString("Script")
		_, _ = w.WriteString(tuples[i].Name)
		if i == 0 {
			_, _ = w.WriteString(" Script = iota")
		}
		_ = w.WriteByte('\n')
	}
	_, _ = w.WriteString(")\n\n")

	// todo implement SRE

	source := w.Bytes()
	var fmtSource []byte
	if fmtSource, err = format.Source(source); err != nil {
		return
	}

	err = ioutil.WriteFile(target, fmtSource, 0644)

	return
}
