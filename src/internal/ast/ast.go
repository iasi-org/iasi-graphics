package ast

import "github.com/iasi-org/iasi-graphics/internal/source"

type Document struct{ Flow *Flow }
type Flow struct {
	Title string
	Items []FlowItem
	Pos   source.Position
}
type FlowItem interface {
	flowItem()
	Position() source.Position
}
type Step struct {
	ID, Label string
	Body      []CommonStmt
	Pos       source.Position
}
type Split struct {
	Steps []*Step
	Pos   source.Position
}
type Highlight struct {
	Text string
	Pos  source.Position
}
type CommonStmt struct {
	Kind, Value string
	Pos         source.Position
}

func (*Step) flowItem()                        {}
func (s *Step) Position() source.Position      { return s.Pos }
func (*Split) flowItem()                       {}
func (s *Split) Position() source.Position     { return s.Pos }
func (*Highlight) flowItem()                   {}
func (h *Highlight) Position() source.Position { return h.Pos }
