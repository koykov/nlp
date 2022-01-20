package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
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
		target = strings.TrimSuffix(input, filepath.Ext(input)) + "_repo.go"
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

	// todo compile target

	return
}
