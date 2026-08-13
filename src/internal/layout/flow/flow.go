package flow

import (
	"fmt"
	"github.com/iasi-org/iasi-graphics/internal/model"
	"github.com/iasi-org/iasi-graphics/internal/scene"
	"github.com/iasi-org/iasi-graphics/internal/textmeasure"
	"github.com/iasi-org/iasi-graphics/internal/theme"
	"strings"
)

type Layout struct {
	Measure textmeasure.Measurer
	Theme   theme.Theme
}

func New() Layout { return Layout{Measure: textmeasure.Heuristic{}, Theme: theme.IASI()} }

type box struct{ x, y, w, h float64 }

func (l Layout) Build(g *model.Graphic) (*scene.Scene, error) {
	if g == nil || g.Flow == nil {
		return nil, fmt.Errorf("flow layout requires a flow model")
	}
	s := &scene.Scene{Width: 1600, Height: 900, Title: g.Title, Description: "Conceptual flow: " + g.Title}
	s.Elements = append(s.Elements, scene.Rect{X: 0, Y: 0, Width: 1600, Height: 900, Fill: l.Theme.Background, Kind: "background"})
	titleLines := l.Measure.Wrap(g.Title, 1200, 52)
	s.Elements = append(s.Elements, scene.Text{X: 800, Y: 78, Size: 52, Anchor: "middle", Fill: l.Theme.Text, Weight: "700", Lines: titleLines, LineHeight: 62, Kind: "title"})
	contentTop := 170.0
	contentBottom := 800.0
	if g.Flow.Highlight != "" {
		contentBottom = 720
	}
	n := len(g.Flow.Items)
	gap := 28.0
	slot := (contentBottom - contentTop - gap*float64(max(0, n-1))) / float64(n)
	if slot > 154 {
		slot = 154
	}
	if slot < 96 {
		return nil, fmt.Errorf("flow contains too much content for the initial 1600x900 layout")
	}
	var previous []box
	for i, item := range g.Flow.Items {
		y := contentTop + float64(i)*(slot+gap)
		var current []box
		if item.Step != nil {
			b := box{x: 490, y: y, w: 620, h: slot}
			current = []box{b}
			l.connector(s, previous, current)
			l.card(s, b, item.Step)
		} else {
			count := len(item.Split)
			w := min(360.0, (1240-gap*float64(count-1))/float64(count))
			total := w*float64(count) + gap*float64(count-1)
			x := (1600 - total) / 2
			for j, e := range item.Split {
				b := box{x: x + float64(j)*(w+gap), y: y, w: w, h: slot}
				current = append(current, b)
				l.card(s, b, e)
			}
			l.connector(s, previous, current)
		}
		previous = current
	}
	if g.Flow.Highlight != "" {
		b := box{x: 300, y: 780, w: 1000, h: 82}
		l.connector(s, previous, []box{b})
		s.Elements = append(s.Elements, scene.Rect{X: b.x, Y: b.y, Width: b.w, Height: b.h, Radius: 20, Fill: l.Theme.AccentSoft, Stroke: l.Theme.Accent, StrokeWidth: 1.5, Kind: "highlight"})
		lines := l.Measure.Wrap(g.Flow.Highlight, b.w-80, 24)
		s.Elements = append(s.Elements, scene.Text{X: 800, Y: b.y + 34, Size: 24, Anchor: "middle", Fill: l.Theme.Text, Weight: "600", Lines: lines, LineHeight: 29, Kind: "highlight-text"})
	}
	return s, nil
}
func (l Layout) card(s *scene.Scene, b box, e *model.Element) {
	s.Elements = append(s.Elements, scene.Rect{X: b.x, Y: b.y, Width: b.w, Height: b.h, Radius: 22, Fill: l.Theme.Card, Stroke: l.Theme.CardStroke, StrokeWidth: 2, Kind: "card"})
	textX := b.x + b.w/2
	if e.Icon != "" {
		s.Elements = append(s.Elements, scene.Icon{Name: e.Icon, X: b.x + 38, Y: b.y + b.h/2 - 17, Size: 34, Stroke: l.Theme.Accent})
		textX += 20
	}
	label := l.Measure.Wrap(e.Label, b.w-110, 26)
	labelY := b.y + 42
	if len(e.Body) == 0 {
		labelY = b.y + b.h/2 - 4
	}
	s.Elements = append(s.Elements, scene.Text{X: textX, Y: labelY, Size: 26, Anchor: "middle", Fill: l.Theme.Text, Weight: "700", Lines: label, LineHeight: 31, Kind: "label"})
	if len(e.Body) > 0 {
		var body []string
		for _, v := range e.Body {
			body = append(body, l.Measure.Wrap(v, b.w-100, 18)...)
		}
		s.Elements = append(s.Elements, scene.Text{X: textX, Y: b.y + b.h - 39, Size: 18, Anchor: "middle", Fill: l.Theme.Muted, Weight: "400", Lines: body, LineHeight: 23, Kind: "body"})
	}
}
func (l Layout) connector(s *scene.Scene, from, to []box) {
	if len(from) == 0 {
		return
	}
	fy := from[0].y + from[0].h
	ty := to[0].y
	mid := (fy + ty) / 2
	var paths []string
	for _, b := range from {
		x := b.x + b.w/2
		paths = append(paths, fmt.Sprintf("M %.1f %.1f L %.1f %.1f", x, fy, x, mid))
	}
	if len(from) > 1 || len(to) > 1 {
		left, right := 800.0, 800.0
		for _, group := range [][]box{from, to} {
			for _, b := range group {
				x := b.x + b.w/2
				if x < left {
					left = x
				}
				if x > right {
					right = x
				}
			}
		}
		paths = append(paths, fmt.Sprintf("M %.1f %.1f L %.1f %.1f", left, mid, right, mid))
	}
	for _, b := range to {
		x := b.x + b.w/2
		paths = append(paths, fmt.Sprintf("M %.1f %.1f L %.1f %.1f", x, mid, x, ty))
	}
	s.Elements = append(s.Elements, scene.Path{Data: strings.Join(paths, " "), Stroke: l.Theme.Connector, StrokeWidth: 3, Fill: "none", Kind: "connector"})
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
