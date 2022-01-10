package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

type languagesModule struct{}

type languagesTuple struct {
	Name    string `json:"name"`
	Native  string `json:"native"`
	Iso6391 string `json:"iso639_1"`
	Iso6393 string `json:"iso639_3"`
}

func (m languagesModule) Validate(input, _ string) error {
	if len(input) == 0 {
		return fmt.Errorf("param -input is required")
	}
	return nil
}

func (m languagesModule) Compile(input, target string) (err error) {
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

	for i := 0; i < len(tuples)-1; i++ {
		tuple := &tuples[i]
		if len(tuple.Native) == 0 {
			tuple.Native = tuple.Name
		}
	}

	return
}
