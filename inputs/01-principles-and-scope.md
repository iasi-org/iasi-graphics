# 01 — Principles and Scope

## P1. Hide implementation complexity

Internal complexity is allowed.

Public complexity is not.

The author must not need to understand SVG, Go internals, layout algorithms, browser APIs or JavaScript.

## P2. Semantic input

The language expresses semantic structures:

- flow;
- comparison;
- ecosystem;
- split;
- emphasis;
- text;
- items.

It must not expose graphical primitives as the primary authoring model.

## P3. No coordinates in the DSL

The following is explicitly forbidden in v0.1:

```text
x: 120
y: 340
width: 500
height: 180
```

If a figure requires this level of control, it is outside the purpose of the first version.

## P4. Reproducibility

The same valid input, engine version and theme must produce deterministic output.

## P5. Text is the source of truth

The `.ig` source or the `{iasi-graphics}` code block is the source artifact.

Generated SVG is derived output.

## P6. SVG first

The first renderer produces SVG.

PNG, PDF and other outputs are future concerns.

## P7. Quarto is a consumer, not the core

The engine must work without Quarto.

Quarto integration is an adapter around the standalone compiler.

## P8. No JavaScript user workflow

The author must never need:

- Node.js;
- npm;
- JavaScript;
- TypeScript;
- a browser runtime;
- a `.js` intermediate file.

The engine is implemented in Go and distributed as a native executable.

## P9. No network requirement at render time

Rendering must work offline.

Themes, fonts assumptions, icons and other required assets must be local or embedded.

## P10. Presentation quality is a product feature

Whitespace, hierarchy, balance and typography are not decoration added later.

They are part of the renderer's correctness.

## MVP scope

v0.1 supports:

- `flow`
- `compare`
- `ecosystem`
- title
- short text
- items/steps
- split
- highlight
- optional named icons from a small built-in set
- one built-in theme: `iasi`
- SVG output
- standalone CLI
- Quarto integration after the standalone MVP

## Non-goals for v0.1

Do not implement yet:

- interactive editor;
- drag and drop;
- arbitrary SVG;
- animation;
- charts/data visualization;
- UML;
- sequence diagrams;
- automatic AI generation;
- multiple renderers;
- PNG renderer;
- PDF renderer;
- PPTX;
- web service;
- plugin marketplace;
- custom user-defined layout algorithms;
- arbitrary CSS;
- arbitrary coordinates;
- a complete general-purpose graphics language.
