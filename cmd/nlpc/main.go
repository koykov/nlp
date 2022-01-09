package main

import (
	"flag"
	"log"
)

var (
	mod, in, trg string

	mods = map[string]module{
		"languages": modLanguages{},
	}
)

func init() {
	rf := func(v *string, names []string, value, usage string) {
		for i := range names {
			flag.StringVar(v, names[i], value, usage)
		}
	}
	rf(&mod, []string{"module", "mod", "m"}, "", "Module to compile: [languages]")
	rf(&in, []string{"input", "in", "i"}, "", "Path to source data file")
	rf(&trg, []string{"target", "t"}, "", "Target file or directory")
	flag.Parse()

	if len(mod) == 0 {
		log.Fatalln("param -module is required")
	}
	if _, ok := mods[mod]; !ok {
		log.Fatalf("unknown module: %s\n", mod)
	}
}

func main() {

}
