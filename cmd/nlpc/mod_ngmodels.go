package main

import (
	"fmt"
)

type ngmodelsModule struct{}

func (m ngmodelsModule) Validate(input, target string) error {
	if len(input) == 0 {
		return fmt.Errorf("param -input is required")
	}
	if isDirWR(target) {
		return fmt.Errorf("target '%s' isn't writable", target)
	}
	return nil
}

func (m ngmodelsModule) Compile(w moduleWriter, input, target string) (err error) {
	_, _, _ = w, input, target
	return
}
