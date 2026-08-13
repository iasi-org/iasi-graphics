# 10 — Testing and Acceptance

## Testing philosophy

The product must be tested at four levels:

1. parser;
2. semantic model;
3. layout/scene;
4. rendered output;
5. Quarto adapter after the core engine is stable.

## Unit tests

### Lexer/parser

Cover:

- valid strings;
- escapes;
- comments;
- nested blocks;
- malformed braces;
- invalid identifiers;
- unexpected tokens;
- source positions.

### Semantic validation

Cover:

- duplicate IDs;
- split with fewer than two steps;
- compare with != 2 sides;
- ecosystem with fewer than 3 or more than 8 items;
- missing ecosystem center;
- unknown icon.

### Layout

Test invariant properties:

- no negative dimensions;
- elements remain within canvas bounds;
- no card overlap for valid MVP fixtures;
- split children are aligned;
- compare sides are balanced;
- ecosystem center remains central;
- deterministic geometry.

### SVG

Test:

- XML parses;
- expected `viewBox`;
- title present;
- no `<script>`;
- no external `http://` or `https://` asset dependencies;
- visible text remains SVG text;
- stable output.

## Golden fixtures

Create at least:

```text
testdata/
├── flow-basic.ig
├── compare-basic.ig
├── ecosystem-basic.ig
└── invalid/
```

Golden SVG snapshots are appropriate once the visual design stabilizes.

## MVP acceptance tests

### AT-001 — Standalone flow

Given a valid flow `.ig` file:

```bash
iasi-graphics render testdata/flow-basic.ig -o flow.svg
```

must succeed and produce valid SVG.

### AT-002 — Compare

A valid two-sided compare source must render without overlap.

### AT-003 — Ecosystem

A valid ecosystem with five items must render around a central concept with balanced spacing.

### AT-004 — No coordinates in source

The public DSL must not require or document coordinate properties.

### AT-005 — No JavaScript

Rendering must not generate or require `.js` files, Node.js or npm.

### AT-006 — Determinism

Two renders of the same source with the same engine version must produce identical normalized SVG.

### AT-007 — Useful errors

Malformed input must fail with line/column diagnostics.

### AT-008 — Offline

Rendering the built-in examples must require no network connection.

### AT-009 — Quarto block

After Quarto integration is implemented, a QMD block:

````markdown
```{iasi-graphics}
...
```
````

must be replaced by a generated figure rather than displayed as source code.

### AT-010 — Source remains the truth

The author can delete generated SVG and regenerate it from `.ig` or the QMD block with no manual graphical editing.

## Visual acceptance checklist

Automated tests cannot fully judge presentation quality.

Before declaring v0.1 visually successful, manually inspect the three reference outputs and confirm:

- clear visual hierarchy;
- generous whitespace;
- no accidental UML appearance;
- no text collisions;
- no cramped cards;
- balanced composition;
- connectors are subordinate to content;
- readable at normal document width;
- figure communicates one main idea;
- output looks intentional rather than algorithmically dumped.

This human visual review is a release criterion.
