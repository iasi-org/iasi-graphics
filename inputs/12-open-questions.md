# 12 — Open Questions

These questions are intentionally **not** part of the first implementation contract unless Codex needs a local placeholder.

Do not invent permanent answers while implementing v0.1.

## OQ-001 — Final language name

Product name is fixed provisionally as:

```text
iasi-graphics
```

The DSL itself does not yet need a separate marketing name.

## OQ-002 — Final icon set

The MVP may use a small bundled set.

The long-term icon strategy is not fixed.

## OQ-003 — Theme extensibility

v0.1 has one theme: `iasi`.

User-defined themes are future work.

## OQ-004 — Canvas/aspect ratios

v0.1 uses a presentation-friendly internal 16:9 design space.

Future layouts may need portrait or content-driven sizing.

Do not expose this complexity prematurely.

## OQ-005 — More layouts

Possible future layouts include:

```text
timeline
cycle
layers
matrix
process
before-after
hierarchy
roadmap
```

They are not part of v0.1.

## OQ-006 — AI generation

A future AI may translate natural-language intent into `.ig`.

That is outside the renderer and outside v0.1.

The important architectural property is that AI would generate **editable declarative source**, not opaque images.

## OQ-007 — Relationship with IASI Quarto tooling

The long-term home of the Quarto integration may interact with `iasi-quarto` or `iasi-lua`.

For the MVP, keep `iasi-graphics` independently usable and keep the adapter thin.

## OQ-008 — PDF/PNG outputs

Future output conversion may be useful.

Do not add these renderers until SVG is mature.

## OQ-009 — Exact typography

The initial theme should use robust font fallbacks.

A final IASI typography standard can be integrated later without changing the DSL.
