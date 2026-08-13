package compiler

import (
	"github.com/iasi-org/iasi-graphics/internal/layout/flow"
	"github.com/iasi-org/iasi-graphics/internal/parser"
	"github.com/iasi-org/iasi-graphics/internal/render/svg"
	"github.com/iasi-org/iasi-graphics/internal/semantic"
	"github.com/iasi-org/iasi-graphics/internal/source"
)

func Compile(name string, data []byte) ([]byte, error) {
	doc, err := parser.Parse(source.New(name, data))
	if err != nil {
		return nil, err
	}
	graphic, err := semantic.Build(doc)
	if err != nil {
		return nil, err
	}
	scene, err := flow.New().Build(graphic)
	if err != nil {
		return nil, err
	}
	return svg.Render(scene)
}
