package scene

type Scene struct {
	Width, Height      float64
	Title, Description string
	Elements           []Element
}
type Element interface{ sceneElement() }
type Rect struct {
	X, Y, Width, Height, Radius float64
	Fill, Stroke                string
	StrokeWidth                 float64
	Kind                        string
}
type Text struct {
	X, Y, Size           float64
	Anchor, Fill, Weight string
	Lines                []string
	LineHeight           float64
	Kind                 string
}
type Path struct {
	Data, Stroke string
	StrokeWidth  float64
	Fill         string
	Kind         string
}
type Icon struct {
	Name       string
	X, Y, Size float64
	Stroke     string
}

func (Rect) sceneElement() {}
func (Text) sceneElement() {}
func (Path) sceneElement() {}
func (Icon) sceneElement() {}
