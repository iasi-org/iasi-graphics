package compiler

import (
	"bytes"
	"os"
	"testing"
)

func TestFlowVerticalSlice(t *testing.T) {
	data, err := os.ReadFile("../../testdata/flow-basic.ig")
	if err != nil {
		t.Fatal(err)
	}
	a, err := Compile("flow-basic.ig", data)
	if err != nil {
		t.Fatal(err)
	}
	b, err := Compile("flow-basic.ig", data)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(a, b) {
		t.Fatal("compile is not deterministic")
	}
	if !bytes.Contains(a, []byte("Del diálogo al artefacto")) {
		t.Fatal("title missing")
	}
}
