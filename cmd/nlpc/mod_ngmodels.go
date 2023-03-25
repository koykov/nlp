package main

import (
	"bufio"
	"fmt"
	"net/http"
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
	_, _ = w, target
	var resp *http.Response
	if resp, err = http.Get(input); err != nil {
		return
	}
	defer func() { _ = resp.Body.Close() }()

	scanner := bufio.NewScanner(resp.Body)

	for scanner.Scan() {
		line := scanner.Bytes()
		_ = line
		// todo remove macroses and parse
	}

	return
}
