# 11 — Implementation Plan

Codex should work in vertical slices.

Do not implement every abstraction before producing the first SVG.

## Phase 0 — Repository foundation

Create:

```text
go.mod
cmd/iasi-graphics/main.go
internal/...
examples/
testdata/
docs/
```

Add:

```bash
go test ./...
go build ./cmd/iasi-graphics
```

No Quarto work yet.

## Phase 1 — Flow vertical slice

Implement the smallest complete path:

```text
flow.ig
  ↓
parse
  ↓
validate
  ↓
semantic model
  ↓
flow layout
  ↓
scene
  ↓
SVG
```

Use one reference example.

Do not implement `compare` or `ecosystem` until flow renders end-to-end.

## Phase 2 — Presentation quality for flow

Improve:

- spacing;
- title hierarchy;
- cards;
- split layout;
- connectors;
- wrapping;
- highlight;
- icons if needed.

The goal is not merely passing XML tests.

The goal is the first convincing figure.

## Phase 3 — Compare

Add syntax, semantic validation and CompareLayout.

Reuse scene and theme primitives.

## Phase 4 — Ecosystem

Add syntax, semantic validation and EcosystemLayout.

Again reuse existing primitives.

## Phase 5 — Diagnostics and CLI polish

Add:

- line/column errors;
- stdin;
- `-o`;
- version;
- validation polish.

## Phase 6 — Golden/reference suite

Freeze three reference examples once visual quality is acceptable.

Add deterministic SVG regression tests.

## Phase 7 — Quarto adapter

Only now implement:

```text
quarto/_extensions/iasi-graphics/
```

with the Lua filter.

Create an end-to-end sample `.qmd`.

## Phase 8 — Packaging

After the product works:

- Windows binary;
- Linux binary;
- macOS binary if desired;
- release automation later.

## Definition of MVP complete

The MVP is complete when:

1. all three layouts render;
2. acceptance tests pass;
3. visual review passes;
4. CLI is usable;
5. Quarto block works;
6. no Node/JavaScript dependency exists;
7. source is fully reproducible.
