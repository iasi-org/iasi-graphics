package model

type Graphic struct {
	Title string
	Flow  *Flow
}
type Flow struct {
	Items     []FlowItem
	Highlight string
}
type FlowItem struct {
	Step  *Element
	Split []*Element
}
type Element struct {
	ID, Label string
	Body      []string
	Icon      string
}
