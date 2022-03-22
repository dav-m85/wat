package internal

import (
	"os"
	"testing"
)

func TestWalk(t *testing.T) {
	ok := false
	ef := func(e *Entry, err error) {
		ok = true
		if err != nil {
			t.Errorf("Should not give error: %s", err)
			return
		}
		expected := `test in tar`
		if e.Excerpt != expected {
			t.Errorf(`Expected "%s", got "%s"`, expected, e.Excerpt)
		}
	}
	Walk(os.DirFS("test"), 0, ef)
	if !ok {
		t.Error("Should have found an entry")
	}
}
