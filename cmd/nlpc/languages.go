package main

import "fmt"

type modLanguages struct{}

func (m modLanguages) Validate(input, _ string) error {
	if len(input) == 0 {
		return fmt.Errorf("param -input is required")
	}
	return nil
}

func (m modLanguages) Compile(input, target string) error {
	//
	return nil
}
