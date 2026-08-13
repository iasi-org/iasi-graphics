package svg

import (
	"bytes"
	"encoding/xml"
	"github.com/iasi-org/iasi-graphics/internal/scene"
	"io"
	"strings"
	"testing"
)

func TestRenderProducesSafeDeterministicXML(t *testing.T) {
	s := &scene.Scene{Width: 1600, Height: 900, Title: "A & B", Description: "Flow", Elements: []scene.Element{scene.Text{X: 10, Y: 20, Size: 12, Anchor: "start", Fill: "#000", Weight: "400", Lines: []string{"Visible <text>"}, LineHeight: 14}}}
	a, err := Render(s)
	if err != nil {
		t.Fatal(err)
	}
	b, _ := Render(s)
	if !bytes.Equal(a, b) {
		t.Fatal("render is not deterministic")
	}
	d := xml.NewDecoder(bytes.NewReader(a))
	for {
		_, e := d.Token()
		if e == io.EOF {
			break
		}
		if e != nil {
			t.Fatal(e)
		}
	}
	out := string(a)
	for _, bad := range []string{"<script", `href="http://`, `href="https://`} {
		if strings.Contains(out, bad) {
			t.Fatalf("contains %q", bad)
		}
	}
	if !strings.Contains(out, `viewBox="0 0 1600 900"`) || !strings.Contains(out, "Visible &lt;text&gt;") {
		t.Fatalf("unexpected SVG: %s", out)
	}
}
