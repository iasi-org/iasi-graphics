# Phase 1 architecture decisions

Status: accepted for Phase 1. Category: B (architecture).

## Compiler boundaries

The pipeline is split into packages for source normalization, lexer, parser/AST,
semantic validation/model, flow-specific layout, renderer-neutral scene, theme,
and SVG rendering. The compiler package only orchestrates these transformations.
This keeps syntax, meaning, geometry, and serialization independently replaceable.

## Flow and split representation

The AST preserves the source structure. The semantic model represents a flow as
an ordered sequence whose entries contain either one element or an ordered split
group. A split is not generalized into a graph: Phase 1 only needs parallel steps
that reconverge, and layout remains flow-specific.

## Scene model and theme boundary

Layout resolves geometry and visual tokens into a small renderer-neutral scene of
rectangles, text, paths, and named icons. The SVG renderer serializes that scene
and makes no semantic or layout decisions. The single built-in IASI theme owns
the palette and font fallback stack.

## Text measurement boundary

Layout depends on a small `Measurer` interface. Phase 1 uses a deterministic,
dependency-free heuristic implementation. A future metric implementation can be
substituted without changing the semantic model or renderer.
