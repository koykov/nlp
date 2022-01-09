package main

type module interface {
	Validate(input, target string, reason *string) bool
}
