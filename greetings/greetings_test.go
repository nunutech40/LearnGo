package greetings

import (
	"testing"
	"regexp"
)

// Test function Hello dengan sebuah nama, dan memeriksa return yang valid
func TestHelloName(t *testing.T) {
	name := "Dorothy"
	want := regexp.MustCompile(`\b` +name+ `\b`)
	msg, err := Hello("Dorothy")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Dorothy") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

// Test function Hello dengan mengisi empty string, dan memeriksa jika ada error
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}