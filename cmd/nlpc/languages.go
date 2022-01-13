package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"go/format"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/koykov/entry"
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

func (m languagesModule) Compile(w moduleWriter, input, target string) (err error) {
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

	_, _ = w.WriteString("import \"github.com/koykov/entry\"\n\nconst (\n")
	for i := 0; i < len(tuples); i++ {
		tuple := &tuples[i]
		if len(tuple.Native) == 0 {
			tuple.Native = tuple.Name
		}
		name := tuple.Name
		name = strings.ReplaceAll(name, " ", "_")
		name = strings.ReplaceAll(name, "-", "_")
		_, _ = w.WriteString(name)
		if i == 0 {
			_, _ = w.WriteString(" Language = iota")
		}
		_ = w.WriteByte('\n')
	}
	_, _ = w.WriteString(")\n\n")

	var (
		lo, hi uint16
		ol     entry.Entry32
	)
	_, _ = w.WriteString("type lt struct {\n\tname, native, iso1, iso3 entry.Entry32\n}\n\nvar (\n")
	_, _ = w.WriteString("__lt_lst = []lt{\n")
	for i := 0; i < len(tuples); i++ {
		tuple := &tuples[i]
		hi = lo + uint16(len(tuple.Name))
		ol.Encode(lo, hi)
		lo = hi
		_, _ = w.WriteString("{name:0x")
		_, _ = w.WriteString(fmt.Sprintf("%08x", ol))

		hi = lo + uint16(len(tuple.Native))
		ol.Encode(lo, hi)
		lo = hi
		_, _ = w.WriteString(",native:0x")
		_, _ = w.WriteString(fmt.Sprintf("%08x", ol))

		hi = lo + uint16(len(tuple.Iso6391))
		ol.Encode(lo, hi)
		lo = hi
		_, _ = w.WriteString(",iso1:0x")
		_, _ = w.WriteString(fmt.Sprintf("%08x", ol))

		hi = lo + uint16(len(tuple.Iso6393))
		ol.Encode(lo, hi)
		lo = hi
		_, _ = w.WriteString(",iso3:0x")
		_, _ = w.WriteString(fmt.Sprintf("%08x", ol))
		_, _ = w.WriteString("},\n")
	}
	_, _ = w.WriteString("}\n")

	_, _ = w.WriteString("__lt_buf = []byte(\"")
	for i := 0; i < len(tuples); i++ {
		tuple := &tuples[i]
		_, _ = w.WriteString(tuple.Name)
		_, _ = w.WriteString(tuple.Native)
		_, _ = w.WriteString(tuple.Iso6391)
		_, _ = w.WriteString(tuple.Iso6393)
	}
	_, _ = w.WriteString("\")\n")

	_, _ = w.WriteString(")\n")

	source := w.Bytes()
	var fmtSource []byte
	if fmtSource, err = format.Source(source); err != nil {
		return
	}

	err = ioutil.WriteFile(target, fmtSource, 0644)

	return
}
