package flow

import (
	"github.com/iasi-org/iasi-graphics/internal/model"
	"github.com/iasi-org/iasi-graphics/internal/scene"
	"reflect"
	"testing"
)

func fixture() *model.Graphic {
	return &model.Graphic{Title: "Flow", Flow: &model.Flow{Items: []model.FlowItem{{Step: &model.Element{ID: "a", Label: "A"}}, {Split: []*model.Element{{ID: "b", Label: "B"}, {ID: "c", Label: "C"}}}, {Step: &model.Element{ID: "d", Label: "D"}}}, Highlight: "Done"}}
}
func TestLayoutIsDeterministicAndWithinCanvas(t *testing.T) {
	l := New()
	a, err := l.Build(fixture())
	if err != nil {
		t.Fatal(err)
	}
	b, _ := l.Build(fixture())
	if !reflect.DeepEqual(a, b) {
		t.Fatal("layout is not deterministic")
	}
	for _, el := range a.Elements {
		if r, ok := el.(scene.Rect); ok {
			if r.Width < 0 || r.Height < 0 || r.X < 0 || r.Y < 0 || r.X+r.Width > a.Width || r.Y+r.Height > a.Height {
				t.Fatalf("rect outside canvas: %+v", r)
			}
		}
	}
}
func TestSplitCardsAreAlignedAndDoNotOverlap(t *testing.T) {
	s, err := New().Build(fixture())
	if err != nil {
		t.Fatal(err)
	}
	var cards []scene.Rect
	for _, el := range s.Elements {
		if r, ok := el.(scene.Rect); ok && r.Kind == "card" {
			cards = append(cards, r)
		}
	}
	if len(cards) != 4 {
		t.Fatalf("got %d cards", len(cards))
	}
	left, right := cards[1], cards[2]
	if left.Y != right.Y || left.Height != right.Height {
		t.Fatal("split cards are not aligned")
	}
	if left.X+left.Width > right.X {
		t.Fatal("split cards overlap")
	}
}
