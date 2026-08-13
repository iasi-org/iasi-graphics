package svg

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"github.com/iasi-org/iasi-graphics/internal/scene"
	"github.com/iasi-org/iasi-graphics/internal/theme"
	"strings"
)

func Render(s *scene.Scene) ([]byte, error) {
	if s == nil {
		return nil, fmt.Errorf("cannot render nil scene")
	}
	t := theme.IASI()
	var b bytes.Buffer
	fmt.Fprintf(&b, "<svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 %.0f %.0f\" role=\"img\" aria-labelledby=\"title desc\">\n", s.Width, s.Height)
	tag(&b, "title", "id=\"title\"", s.Title)
	tag(&b, "desc", "id=\"desc\"", s.Description)
	fmt.Fprintf(&b, "  <g font-family=\"%s\">\n", attr(t.FontFamily))
	for _, el := range s.Elements {
		switch e := el.(type) {
		case scene.Rect:
			fmt.Fprintf(&b, "    <rect x=\"%.1f\" y=\"%.1f\" width=\"%.1f\" height=\"%.1f\" rx=\"%.1f\" fill=\"%s\" stroke=\"%s\" stroke-width=\"%.1f\"/>\n", e.X, e.Y, e.Width, e.Height, e.Radius, attr(e.Fill), attr(e.Stroke), e.StrokeWidth)
		case scene.Path:
			fmt.Fprintf(&b, "    <path d=\"%s\" fill=\"%s\" stroke=\"%s\" stroke-width=\"%.1f\" stroke-linecap=\"round\" stroke-linejoin=\"round\"/>\n", attr(e.Data), attr(e.Fill), attr(e.Stroke), e.StrokeWidth)
		case scene.Text:
			renderText(&b, e)
		case scene.Icon:
			renderIcon(&b, e)
		}
	}
	b.WriteString("  </g>\n</svg>\n")
	return b.Bytes(), nil
}
func renderText(b *bytes.Buffer, e scene.Text) {
	fmt.Fprintf(b, "    <text x=\"%.1f\" y=\"%.1f\" text-anchor=\"%s\" fill=\"%s\" font-size=\"%.1f\" font-weight=\"%s\">", e.X, e.Y, attr(e.Anchor), attr(e.Fill), e.Size, attr(e.Weight))
	for i, line := range e.Lines {
		dy := 0.0
		if i > 0 {
			dy = e.LineHeight
		}
		fmt.Fprintf(b, "<tspan x=\"%.1f\" dy=\"%.1f\">%s</tspan>", e.X, dy, text(line))
	}
	b.WriteString("</text>\n")
}
func renderIcon(b *bytes.Buffer, e scene.Icon) {
	x, y, z := e.X, e.Y, e.Size
	fmt.Fprintf(b, "    <g fill=\"none\" stroke=\"%s\" stroke-width=\"3\" stroke-linecap=\"round\" stroke-linejoin=\"round\">", attr(e.Stroke))
	switch e.Name {
	case "message":
		fmt.Fprintf(b, "<rect x=\"%.1f\" y=\"%.1f\" width=\"%.1f\" height=\"%.1f\" rx=\"6\"/><path d=\"M %.1f %.1f l -6 8 v -8\"/>", x, y, z, z*.72, x+z*.3, y+z*.72)
	case "route":
		fmt.Fprintf(b, "<path d=\"M %.1f %.1f C %.1f %.1f %.1f %.1f %.1f %.1f\"/><circle cx=\"%.1f\" cy=\"%.1f\" r=\"3\"/><circle cx=\"%.1f\" cy=\"%.1f\" r=\"3\"/>", x, y+z, x, y, x+z, y+z, x+z, y, x, y+z, x+z, y)
	case "decision":
		fmt.Fprintf(b, "<path d=\"M %.1f %.1f l %.1f %.1f l %.1f %.1f l %.1f %.1f Z\"/><path d=\"M %.1f %.1f v %.1f\"/>", x+z/2, y, z/2, z/2, -z/2, z/2, -z/2, -z/2, x+z/2, y+z, 8.0)
	case "book":
		fmt.Fprintf(b, "<path d=\"M %.1f %.1f h %.1f q 8 0 8 8 v %.1f q -8 -8 -8 -8 h %.1f Z\"/><path d=\"M %.1f %.1f v %.1f\"/>", x, y, z/2-2, z*.75, -z/2+2, x+z/2, y+8, z*.75)
	default:
		fmt.Fprintf(b, "<rect x=\"%.1f\" y=\"%.1f\" width=\"%.1f\" height=\"%.1f\" rx=\"4\"/><path d=\"M %.1f %.1f h %.1f M %.1f %.1f h %.1f\"/>", x, y, z, z, x+5, y+z*.35, z-10, x+5, y+z*.65, z-10)
	}
	b.WriteString("</g>\n")
}
func tag(b *bytes.Buffer, name, attrs, value string) {
	fmt.Fprintf(b, "  <%s %s>%s</%s>\n", name, attrs, text(value), name)
}
func text(v string) string { var b strings.Builder; xml.EscapeText(&b, []byte(v)); return b.String() }
func attr(v string) string { return text(v) }
