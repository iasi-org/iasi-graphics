package lexer

import (
	"github.com/iasi-org/iasi-graphics/internal/source"
	"testing"
)

func TestLexStringsEscapesCommentsAndPositions(t *testing.T) {
	tokens, err := Lex(source.New("flow.ig", []byte("# note\nflow \"A\\nB\\\"C\\\\D\" { step one \"One\" }")))
	if err != nil {
		t.Fatal(err)
	}
	if tokens[0].Text != "flow" || tokens[0].Pos.Line != 2 || tokens[0].Pos.Column != 1 {
		t.Fatalf("unexpected first token: %+v", tokens[0])
	}
	if tokens[1].Text != "A\nB\"C\\D" {
		t.Fatalf("unexpected string: %q", tokens[1].Text)
	}
}
func TestLexRejectsInvalidCharacter(t *testing.T) {
	_, err := Lex(source.New("bad.ig", []byte("flow @")))
	if err == nil {
		t.Fatal("expected error")
	}
}
