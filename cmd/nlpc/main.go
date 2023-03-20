package main

import (
	"bytes"
	"flag"
	"log"
	"os"
)

type moduleWriter interface {
	Write(p []byte) (int, error)
	WriteString(s string) (int, error)
	WriteByte(b byte) error
	Bytes() []byte
	String() string
}

var (
	fmod, fin, ftrg string

	mods = map[string]module{
		"languages": languagesModule{},
		"scripts":   scriptsModule{},
		"ngmodels":  ngmodelsModule{},
	}
)

func init() {
	rf := func(v *string, names []string, value, usage string) {
		for i := range names {
			flag.StringVar(v, names[i], value, usage)
		}
	}
	rf(&fmod, []string{"module", "mod", "m"}, "", "Module to compile: [languages, scripts, ngmodels]")
	rf(&fin, []string{"input", "in", "i"}, "", "Path to source data file")
	rf(&ftrg, []string{"target", "t"}, "", "Target file or directory")
	rf(&ftrg, []string{"source", "src", "s"}, "", "Data source")
	flag.Parse()

	if len(fmod) == 0 {
		log.Fatalln("param -module is required")
	}
	var (
		mod module
		ok  bool
		err error
	)
	if mod, ok = mods[fmod]; !ok {
		log.Fatalf("unknown module: %s\n", fmod)
	}
	if err = mod.Validate(fin, ftrg); err != nil {
		log.Fatalf("module validation failed: %s\n", err.Error())
	}
}

func main() {
	var (
		err error
		mod module
		w   bytes.Buffer
	)

	_, _ = w.WriteString("// Code generated by \"")
	for i := 0; i < len(os.Args); i++ {
		if i > 0 {
			w.WriteByte(' ')
		}
		w.WriteString(os.Args[i])
	}
	_, _ = w.WriteString("\". DO NOT EDIT.\n\n")
	_, _ = w.WriteString("package nlp\n\n")

	mod = mods[fmod]
	log.Printf("%s compilation started\n", fmod)
	if err = mod.Compile(&w, fin, ftrg); err != nil {
		log.Fatalf("%s compilation failed: %s\n", fmod, err.Error())
	}
	log.Printf("%s compilation done\n", fmod)
}
