package crdtgo

import (
	"testing"
)

func Test_New(t *testing.T) {
	crdtA := New()
	crdtB := New("a", "b", "c")

	if crdtB.Size() != 3 {
		t.Error("crdtB.size() returned %d, should be %d", crdtB.Size(), 3)
	}

	if crdtA.Size() != 0 {
		t.Error("crdtA.size() returned %d, should be %d", crdtA.Size(), 0)
	}
}
