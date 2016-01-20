package crdt

import (
	"testing"
)

func TestAdd(t *testing.T) {
	expected := "Resp(0)"
	actual := Add(4, "mysql")
	if actual != expected {
		t.Error("test failed")
	}
}
