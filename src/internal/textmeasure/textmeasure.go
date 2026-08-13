package textmeasure

import "strings"

type Measurer interface {
	Wrap(text string, maxWidth, fontSize float64) []string
}
type Heuristic struct{}

func (Heuristic) Wrap(text string, maxWidth, fontSize float64) []string {
	if text == "" {
		return nil
	}
	max := int(maxWidth / (fontSize * 0.54))
	if max < 1 {
		max = 1
	}
	var lines []string
	var current string
	for _, word := range strings.Fields(text) {
		next := word
		if current != "" {
			next = current + " " + word
		}
		if len([]rune(next)) <= max || current == "" {
			current = next
		} else {
			lines = append(lines, current)
			current = word
		}
	}
	if current != "" {
		lines = append(lines, current)
	}
	return lines
}
