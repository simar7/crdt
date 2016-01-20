package crdt

import (
	"fmt"
	"testing"
)

func Test_New(t *testing.T) {
	crdtA := New()
	crdtB := New("a", "b", "c")

	if crdtB.Size() != 3 {
		t.Error(fmt.Sprintf("crdtB.Size() returned %d, should be %d", crdtB.Size(), 3))
	}

	if crdtA.Size() != 0 {
		t.Error(fmt.Sprintf("crdtA.Size() returned %d, should be %d", crdtA.Size(), 0))
	}

	fmt.Println(crdtB.String())
	crdtB.Remove("a")
	fmt.Println(crdtB.String())

	if crdtB.Size() != 2 {
		t.Error(fmt.Sprintf("crdtB.Size() returned %d, should be %d", crdtB.Size(), 2))
	}

	if crdtB.Get("b") != 0 {
		t.Error(fmt.Sprintf("crdtB.Get() returned %d, should be %d", crdtB.Get("b"), 0))
	}

}
