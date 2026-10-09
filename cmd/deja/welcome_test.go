package main

import (
	"bytes"
	"strings"
	"testing"

	"github.com/vshulcz/deja-vu/internal/index"
)

// A bare `deja` into a pipe or a script used to print the full usage block,
// which starts with commands only harnesses call. Issue #4621 wants a short
// welcome instead: what is indexed and one thing to try, with the full list
// kept behind `deja help`. This test pins the welcome shape.
func TestBareDejaPrintsWelcomeNotFullUsage(t *testing.T) {
	hermeticEnv(t)
	var out bytes.Buffer
	if err := printWelcome(index.DefaultDir(), &out); err != nil {
		t.Fatal(err)
	}
	s := out.String()
	if strings.Contains(s, "Usage:") {
		t.Fatalf("welcome fell back to the full usage block:\n%s", s)
	}
	if !strings.Contains(s, "deja help") {
		t.Fatalf("welcome does not point at the full list behind deja help:\n%s", s)
	}
	if !strings.Contains(s, "deja ") {
		t.Fatalf("welcome does not suggest one thing to try:\n%s", s)
	}
}
