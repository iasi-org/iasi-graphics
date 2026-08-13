package semantic

import (
	"github.com/iasi-org/iasi-graphics/internal/parser"
	"github.com/iasi-org/iasi-graphics/internal/source"
	"strings"
	"testing"
)

func build(t *testing.T, s string) error {
	t.Helper()
	d, e := parser.Parse(source.New("test.ig", []byte(s)))
	if e != nil {
		return e
	}
	_, e = Build(d)
	return e
}
func TestValidationErrors(t *testing.T) {
	tests := []struct{ name, src, want string }{
		{"duplicate", `flow "X" { step a "A" step a "B" }`, "duplicate id"},
		{"small split", `flow "X" { split { step a "A" } }`, "at least 2"},
		{"unknown icon", `flow "X" { step a "A" { icon "remote" } }`, "unknown icon"},
		{"empty", `flow "X" { highlight "H" }`, "at least one step"},
		{"highlights", `flow "X" { step a "A" highlight "H" highlight "I" }`, "at most one highlight"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := build(t, tt.src)
			if err == nil || !strings.Contains(err.Error(), tt.want) {
				t.Fatalf("got %v, want %q", err, tt.want)
			}
		})
	}
}
func TestBuildNormalizesSplit(t *testing.T) {
	src := `flow "X" { step a "A" split { step b "B" step c "C" } }`
	d, _ := parser.Parse(source.New("x.ig", []byte(src)))
	g, err := Build(d)
	if err != nil {
		t.Fatal(err)
	}
	if len(g.Flow.Items) != 2 || len(g.Flow.Items[1].Split) != 2 {
		t.Fatalf("unexpected model: %#v", g)
	}
}
