package main

import (
	"fmt"
	"os"
	"syscall"
)

type ngmodelsModule struct{}

func (m ngmodelsModule) Validate(input, target string) error {
	if len(input) == 0 {
		return fmt.Errorf("param -input is required")
	}
	if !m.isDirWR(target) {
		return fmt.Errorf("target '%s' isn't writable", target)
	}
	return nil
}

func (m ngmodelsModule) Compile(w moduleWriter, input, target string) (err error) {
	_, _ = input, target
	return
}

func (m ngmodelsModule) isDirWR(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		return false
	}
	if !fi.IsDir() {
		return false
	}
	if fi.Mode().Perm()&(1<<(uint(7))) == 0 {
		return false
	}
	var stat syscall.Stat_t
	if err = syscall.Stat(path, &stat); err != nil {
		return false
	}
	if uint32(os.Geteuid()) != stat.Uid {
		return false
	}
	return true
}
