package lexer

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/iasi-org/iasi-graphics/internal/source"
)

type Kind int

const (
	EOF Kind = iota
	Identifier
	String
	LBrace
	RBrace
)

type Token struct {
	Kind Kind
	Text string
	Pos  source.Position
}

type Error struct {
	Name          string
	Pos           source.Position
	Message, Line string
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s:%d:%d: %s\n%s", e.Name, e.Pos.Line, e.Pos.Column, e.Message, e.Line)
}

func Lex(file source.File) ([]Token, error) {
	l := scanner{name: file.Name, input: file.Text, line: 1, col: 1}
	var tokens []Token
	for {
		t, err := l.next()
		if err != nil {
			return nil, err
		}
		tokens = append(tokens, t)
		if t.Kind == EOF {
			return tokens, nil
		}
	}
}

type scanner struct {
	name, input       string
	offset, line, col int
}

func (s *scanner) next() (Token, error) {
	for s.offset < len(s.input) {
		r, _ := utf8.DecodeRuneInString(s.input[s.offset:])
		if r == ' ' || r == '\t' || r == '\n' {
			s.advance(r)
			continue
		}
		if r == '#' {
			for s.offset < len(s.input) {
				r, _ = utf8.DecodeRuneInString(s.input[s.offset:])
				s.advance(r)
				if r == '\n' {
					break
				}
			}
			continue
		}
		break
	}
	pos := source.Position{Line: s.line, Column: s.col}
	if s.offset >= len(s.input) {
		return Token{Kind: EOF, Pos: pos}, nil
	}
	r, _ := utf8.DecodeRuneInString(s.input[s.offset:])
	switch r {
	case '{':
		s.advance(r)
		return Token{Kind: LBrace, Text: "{", Pos: pos}, nil
	case '}':
		s.advance(r)
		return Token{Kind: RBrace, Text: "}", Pos: pos}, nil
	case '"':
		return s.scanString(pos)
	}
	if isIdentStart(r) {
		start := s.offset
		for s.offset < len(s.input) {
			r, _ = utf8.DecodeRuneInString(s.input[s.offset:])
			if !isIdentContinue(r) {
				break
			}
			s.advance(r)
		}
		return Token{Kind: Identifier, Text: s.input[start:s.offset], Pos: pos}, nil
	}
	return Token{}, s.err(pos, fmt.Sprintf("unexpected character %q", r))
}

func (s *scanner) scanString(pos source.Position) (Token, error) {
	s.advance('"')
	var b strings.Builder
	for s.offset < len(s.input) {
		r, _ := utf8.DecodeRuneInString(s.input[s.offset:])
		s.advance(r)
		if r == '"' {
			return Token{Kind: String, Text: b.String(), Pos: pos}, nil
		}
		if r == '\n' {
			return Token{}, s.err(pos, "unterminated string")
		}
		if r != '\\' {
			b.WriteRune(r)
			continue
		}
		if s.offset >= len(s.input) {
			break
		}
		esc, _ := utf8.DecodeRuneInString(s.input[s.offset:])
		s.advance(esc)
		switch esc {
		case '"':
			b.WriteRune('"')
		case '\\':
			b.WriteRune('\\')
		case 'n':
			b.WriteRune('\n')
		default:
			return Token{}, s.err(source.Position{Line: s.line, Column: s.col - 1}, fmt.Sprintf("unsupported escape \\%c", esc))
		}
	}
	return Token{}, s.err(pos, "unterminated string")
}

func (s *scanner) advance(r rune) {
	s.offset += utf8.RuneLen(r)
	if r == '\n' {
		s.line++
		s.col = 1
	} else {
		s.col++
	}
}
func isIdentStart(r rune) bool    { return r == '_' || r >= 'A' && r <= 'Z' || r >= 'a' && r <= 'z' }
func isIdentContinue(r rune) bool { return isIdentStart(r) || r >= '0' && r <= '9' || r == '-' }
func (s *scanner) err(pos source.Position, message string) error {
	lines := strings.Split(s.input, "\n")
	line := ""
	if pos.Line > 0 && pos.Line <= len(lines) {
		line = lines[pos.Line-1]
	}
	return &Error{Name: s.name, Pos: pos, Message: message, Line: line}
}
