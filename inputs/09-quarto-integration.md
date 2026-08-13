# 09 — Quarto Integration

## Goal

Allow an author to embed `iasi-graphics` directly in a `.qmd` file using standard Quarto executable-block syntax.

Canonical form:

````markdown
```{iasi-graphics}
flow "Del diálogo al artefacto" {
  step conversation "Conversación"
  step knowledge "Conocimiento"
  step artifacts "Artefactos"
}
```
````

## Design principle

The Quarto integration must be a thin adapter.

It must not duplicate:

- parser;
- semantic validation;
- layout;
- SVG rendering.

All those belong to the Go engine.

## Adapter technology

Use a Quarto/Pandoc **Lua filter**.

The Lua filter identifies code blocks whose engine/class is `iasi-graphics`, invokes the `iasi-graphics` executable and replaces the code block with an image/figure node.

## Conceptual pipeline

```text
QMD CodeBlock
    ↓
Lua filter
    ↓
source text
    ↓
iasi-graphics render
    ↓
SVG file
    ↓
Pandoc Image/Figure
```

## Figure options

The first integration should preserve or map common Quarto figure metadata where practical.

Priority:

```text
fig-cap
fig-alt
label
```

Example:

````markdown
```{iasi-graphics}
#| fig-cap: "Del diálogo al artefacto"
#| fig-alt: "Flujo conceptual desde conversación hasta artefactos"

flow "Del diálogo al artefacto" {
  ...
}
```
````

The exact parsing mechanism for `#|` options should follow normal Quarto conventions.

## Generated file location

Generated SVG assets must be written to a deterministic build location.

Requirements:

- do not pollute the source directory unnecessarily;
- filenames must be stable for the same block;
- multiple blocks must not collide;
- rebuilds should be safe.

A content hash or document/block-derived stable identifier is acceptable.

## Quarto formats

The adapter's responsibility ends at producing a valid SVG figure reference.

HTML should consume SVG directly.

PDF conversion, when required by Quarto/Pandoc/LaTeX, is outside the core renderer and should use the normal Quarto publication pipeline.

Do not add a second PDF graphics renderer merely for Quarto.

## Missing executable

If `iasi-graphics` is not available, the filter must produce a clear error explaining what is missing.

Do not silently leave the original code block unchanged.

## Implementation order

Do not implement this integration until:

```text
.ig → CLI → SVG
```

works reliably.
