package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Hector"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Hector")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Hector) = %q, %v, quiere coincidencia para %#q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, quiere "", error`, msg, err)
	}

}
