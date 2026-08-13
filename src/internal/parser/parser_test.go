package parser

import (
	"github.com/iasi-org/iasi-graphics/internal/ast"
	"github.com/iasi-org/iasi-graphics/internal/source"
	"strings"
	"testing"
)

func TestParseFlowStructure(t *testing.T) {
	src := `flow "Title" { step a "A" { text "Body" } split { step b "B" step c "C" } highlight "Done" }`
	doc, err := Parse(source.New("flow.ig", []byte(src)))
	if err != nil {
		t.Fatal(err)
	}
	if doc.Flow.Title != "Title" || len(doc.Flow.Items) != 3 {
		t.Fatalf("unexpected flow: %#v", doc.Flow)
	}
	split, ok := doc.Flow.Items[1].(*ast.Split)
	if !ok || len(split.Steps) != 2 {
		t.Fatalf("unexpected split: %#v", doc.Flow.Items[1])
	}
}
func TestParseReportsLocationForMissingBrace(t *testing.T) {
	_, err := Parse(source.New("broken.ig", []byte("flow \"X\" {\n step a \"A\"")))
	if err == nil || !strings.Contains(err.Error(), "broken.ig:2:12") || !strings.Contains(err.Error(), "expected '}'") {
		t.Fatalf("unexpected error: %v", err)
	}
}
func TestParseRejectsOtherTopLevelKinds(t *testing.T) {
	_, err := Parse(source.New("x.ig", []byte(`compare "X" {}`)))
	if err == nil {
		t.Fatal("expected error")
	}
}
