package main

type module interface {
	Validate(input, target string) error
	Compile(input, target string) error
}
