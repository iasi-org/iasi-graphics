# IASI Graphics — Specification Pack v0.1

This directory is the implementation contract for the first MVP of **`iasi-graphics`**.

The purpose of `iasi-graphics` is to generate presentation-quality conceptual graphics from a small declarative textual language.

The author expresses **meaning and structure**, not coordinates, SVG primitives, PowerPoint instructions, JavaScript or layout details.

## Decisions already fixed

- Product/repository name: `iasi-graphics`
- Standalone source extension: `.ig`
- Engine implementation language: **Go**
- Primary output: **SVG**
- Quarto block syntax: ```` ```{iasi-graphics} ````
- Quarto adapter: **Lua filter**
- No JavaScript, Node.js, npm, browser runtime or PPTX in the user workflow
- No coordinates in the public DSL
- PlantUML remains the tool for technical diagrams
- `iasi-graphics` targets conceptual, editorial and presentation-style graphics

## Architectural pipeline

```text
.ig source
   ↓
parser
   ↓
AST
   ↓
semantic model
   ↓
layout engine
   ↓
scene model
   ↓
SVG renderer
   ↓
.svg
```

Quarto integration adds only an adapter around the same engine:

```text
.qmd
 ↓
```{iasi-graphics}
...
```
 ↓
Lua filter
 ↓
iasi-graphics engine
 ↓
SVG
 ↓
Quarto/Pandoc figure
```

## Documents

1. `00-product-vision.md`
2. `01-principles-and-scope.md`
3. `02-architecture.md`
4. `03-dsl.md`
5. `04-parser-and-ast.md`
6. `05-semantic-model.md`
7. `06-layout-engine.md`
8. `07-svg-renderer.md`
9. `08-cli.md`
10. `09-quarto-integration.md`
11. `10-testing-and-acceptance.md`
12. `11-implementation-plan.md`
13. `12-open-questions.md`
14. `13-codex-instructions.md`
15. `14-development-workflow.md`

## Priority

Codex should implement the standalone engine first.

Do **not** start with Quarto integration.

The first proof is:

```text
example.ig → iasi-graphics → example.svg
```

Only after the standalone engine is stable should the Lua/Quarto adapter be implemented.
