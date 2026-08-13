package semantic

import (
	"fmt"
	"github.com/iasi-org/iasi-graphics/internal/ast"
	"github.com/iasi-org/iasi-graphics/internal/model"
)

var icons = map[string]bool{"message": true, "route": true, "decision": true, "book": true, "boxes": true}

func Build(doc *ast.Document) (*model.Graphic, error) {
	if doc == nil || doc.Flow == nil {
		return nil, fmt.Errorf("semantic: flow is required")
	}
	out := &model.Graphic{Title: doc.Flow.Title, Flow: &model.Flow{}}
	ids := map[string]bool{}
	highlights := 0
	steps := 0
	convert := func(s *ast.Step) (*model.Element, error) {
		if ids[s.ID] {
			return nil, fmt.Errorf("%d:%d: duplicate id %q", s.Pos.Line, s.Pos.Column, s.ID)
		}
		ids[s.ID] = true
		e := &model.Element{ID: s.ID, Label: s.Label}
		for _, stmt := range s.Body {
			if stmt.Kind == "text" {
				e.Body = append(e.Body, stmt.Value)
			} else {
				if !icons[stmt.Value] {
					return nil, fmt.Errorf("%d:%d: unknown icon %q", stmt.Pos.Line, stmt.Pos.Column, stmt.Value)
				}
				if e.Icon != "" {
					return nil, fmt.Errorf("%d:%d: step %q has more than one icon", stmt.Pos.Line, stmt.Pos.Column, s.ID)
				}
				e.Icon = stmt.Value
			}
		}
		return e, nil
	}
	for _, item := range doc.Flow.Items {
		switch n := item.(type) {
		case *ast.Step:
			e, err := convert(n)
			if err != nil {
				return nil, err
			}
			steps++
			out.Flow.Items = append(out.Flow.Items, model.FlowItem{Step: e})
		case *ast.Split:
			if len(n.Steps) < 2 {
				return nil, fmt.Errorf("%d:%d: split requires at least 2 steps; found %d", n.Pos.Line, n.Pos.Column, len(n.Steps))
			}
			group := model.FlowItem{}
			for _, s := range n.Steps {
				e, err := convert(s)
				if err != nil {
					return nil, err
				}
				steps++
				group.Split = append(group.Split, e)
			}
			out.Flow.Items = append(out.Flow.Items, group)
		case *ast.Highlight:
			highlights++
			if highlights > 1 {
				return nil, fmt.Errorf("%d:%d: flow allows at most one highlight", n.Pos.Line, n.Pos.Column)
			}
			out.Flow.Highlight = n.Text
		}
	}
	if steps == 0 {
		return nil, fmt.Errorf("%d:%d: flow requires at least one step", doc.Flow.Pos.Line, doc.Flow.Pos.Column)
	}
	return out, nil
}
