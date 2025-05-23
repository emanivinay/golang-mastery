package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Roger"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello(name)
	if err != nil || !want.MatchString(msg) {
		t.Errorf(`Hello("Roger") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if err == nil || msg != "" {
		t.Errorf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
