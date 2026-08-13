package parser

import (
	"fmt"
	"strings"

	"github.com/iasi-org/iasi-graphics/internal/ast"
	"github.com/iasi-org/iasi-graphics/internal/lexer"
	"github.com/iasi-org/iasi-graphics/internal/source"
)

type Error struct {
	Name          string
	Pos           source.Position
	Message, Line string
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s:%d:%d: %s\n%s", e.Name, e.Pos.Line, e.Pos.Column, e.Message, e.Line)
}

func Parse(file source.File) (*ast.Document, error) {
	tokens, err := lexer.Lex(file)
	if err != nil {
		return nil, err
	}
	p := &parser{name: file.Name, input: file.Text, tokens: tokens}
	if !p.is("flow") {
		return nil, p.err(p.peek(), "expected 'flow'")
	}
	flow, err := p.flow()
	if err != nil {
		return nil, err
	}
	if p.peek().Kind != lexer.EOF {
		return nil, p.err(p.peek(), "expected end of file after flow")
	}
	return &ast.Document{Flow: flow}, nil
}

type parser struct {
	name, input string
	tokens      []lexer.Token
	at          int
}

func (p *parser) flow() (*ast.Flow, error) {
	start := p.take()
	title, err := p.require(lexer.String, "expected flow title")
	if err != nil {
		return nil, err
	}
	if _, err = p.require(lexer.LBrace, "expected '{' after flow title"); err != nil {
		return nil, err
	}
	f := &ast.Flow{Title: title.Text, Pos: start.Pos}
	for p.peek().Kind != lexer.RBrace {
		if p.peek().Kind == lexer.EOF {
			return nil, p.err(p.peek(), "expected '}' after flow block")
		}
		var item ast.FlowItem
		switch p.peek().Text {
		case "step":
			item, err = p.step()
		case "split":
			item, err = p.split()
		case "highlight":
			item, err = p.highlight()
		default:
			return nil, p.err(p.peek(), "expected 'step', 'split', or 'highlight'")
		}
		if err != nil {
			return nil, err
		}
		f.Items = append(f.Items, item)
	}
	p.take()
	return f, nil
}
func (p *parser) step() (*ast.Step, error) {
	start := p.take()
	id, err := p.require(lexer.Identifier, "expected step identifier")
	if err != nil {
		return nil, err
	}
	label, err := p.require(lexer.String, "expected step label")
	if err != nil {
		return nil, err
	}
	s := &ast.Step{ID: id.Text, Label: label.Text, Pos: start.Pos}
	if p.peek().Kind != lexer.LBrace {
		return s, nil
	}
	p.take()
	for p.peek().Kind != lexer.RBrace {
		if p.peek().Kind == lexer.EOF {
			return nil, p.err(p.peek(), "expected '}' after step block")
		}
		if !p.is("text") && !p.is("icon") {
			return nil, p.err(p.peek(), "expected 'text' or 'icon'")
		}
		kind := p.take()
		value, e := p.require(lexer.String, fmt.Sprintf("expected string after '%s'", kind.Text))
		if e != nil {
			return nil, e
		}
		s.Body = append(s.Body, ast.CommonStmt{Kind: kind.Text, Value: value.Text, Pos: kind.Pos})
	}
	p.take()
	return s, nil
}
func (p *parser) split() (*ast.Split, error) {
	start := p.take()
	if _, err := p.require(lexer.LBrace, "expected '{' after split"); err != nil {
		return nil, err
	}
	s := &ast.Split{Pos: start.Pos}
	for p.peek().Kind != lexer.RBrace {
		if p.peek().Kind == lexer.EOF {
			return nil, p.err(p.peek(), "expected '}' after split block")
		}
		if !p.is("step") {
			return nil, p.err(p.peek(), "expected 'step' inside split")
		}
		step, err := p.step()
		if err != nil {
			return nil, err
		}
		s.Steps = append(s.Steps, step)
	}
	p.take()
	return s, nil
}
func (p *parser) highlight() (*ast.Highlight, error) {
	start := p.take()
	value, err := p.require(lexer.String, "expected highlight text")
	if err != nil {
		return nil, err
	}
	return &ast.Highlight{Text: value.Text, Pos: start.Pos}, nil
}
func (p *parser) peek() lexer.Token { return p.tokens[p.at] }
func (p *parser) take() lexer.Token { t := p.peek(); p.at++; return t }
func (p *parser) is(text string) bool {
	return p.peek().Kind == lexer.Identifier && p.peek().Text == text
}
func (p *parser) require(kind lexer.Kind, msg string) (lexer.Token, error) {
	t := p.peek()
	if t.Kind != kind {
		return lexer.Token{}, p.err(t, msg)
	}
	p.at++
	return t, nil
}
func (p *parser) err(t lexer.Token, msg string) error {
	lines := strings.Split(p.input, "\n")
	line := ""
	if t.Pos.Line > 0 && t.Pos.Line <= len(lines) {
		line = lines[t.Pos.Line-1]
	}
	return &Error{Name: p.name, Pos: t.Pos, Message: msg, Line: line}
}
