package main

import (
	"errors"
	"flag"
	"fmt"
)

type ngmodelsModule struct{}

func (m ngmodelsModule) Validate(input, target string) error {
	if len(input) == 0 {
		return fmt.Errorf("param -input is required")
	}
	fsrc := flag.Lookup("source")
	if fsrc == nil {
		return errors.New("no source directory provided")
	}
	if isDirWR(target) {
		return fmt.Errorf("target '%s' isn't writable", target)
	}
	return nil
}

func (m ngmodelsModule) Compile(w moduleWriter, input, target string) (err error) {
	fsrc := flag.Lookup("source")
	src := fsrc.Value.String()
	_, _, _, _ = w, input, src, target
	return
}
