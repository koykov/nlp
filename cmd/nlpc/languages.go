package main

type modLanguages struct{}

func (m modLanguages) Validate(input, _ string, reason *string) bool {
	if len(input) == 0 {
		*reason = "param -input is required"
		return false
	}
	return true
}
