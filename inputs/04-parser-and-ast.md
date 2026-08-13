# 04 — Parser and AST

## Goal

Parse `.ig` source into a faithful structural AST without introducing layout or SVG concerns.

## Diagnostics

Every syntax error must include:

- source name;
- line;
- column;
- concise message;
- when practical, the offending source line.

Example:

```text
examples/flow.ig:8:3: expected '}' after split block
```

## AST requirements

The AST should preserve:

- top-level layout kind;
- title;
- identifiers;
- labels;
- nested statements;
- source positions.

Suggested conceptual types:

```go
type Document struct {
    Visual Visual
}

type Visual interface {
    visualNode()
}

type Flow struct {
    Title string
    Items []FlowItem
}

type Step struct {
    ID    string
    Label string
    Body  []CommonStmt
}

type Split struct {
    Steps []Step
}

type Compare struct {
    Title     string
    Sides     []Side
    Highlight *string
}

type Ecosystem struct {
    Title  string
    Center string
    Items  []Item
}
```

Exact Go naming is not normative.

## Parser strategy

A handwritten lexer and recursive-descent parser are preferred for v0.1.

Reasons:

- grammar is small;
- better control over diagnostics;
- avoids a parser-generator dependency;
- easy to evolve intentionally.

## Parser must not

The parser must not:

- calculate coordinates;
- select fonts;
- emit SVG;
- decide colors;
- infer presentation styles;
- perform layout.

## Semantic validation after parsing

Examples of semantic errors:

```text
duplicate id 'journey'
compare requires exactly 2 sides
ecosystem requires 3 to 8 items
split requires at least 2 steps
unknown icon 'my-company-logo'
```

These are not lexer/parser errors unless the syntax itself is invalid.
