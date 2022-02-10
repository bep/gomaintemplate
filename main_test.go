package main

import (
	"testing"

	qt "github.com/frankban/quicktest"
)

func TestFoo(t *testing.T) {
	t.Parallel()

	c := qt.New(t)

	c.Assert(true, qt.IsTrue)
}
