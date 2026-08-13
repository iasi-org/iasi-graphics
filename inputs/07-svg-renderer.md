# 07 — SVG Renderer

## Responsibility

Convert a fully laid-out Scene Model to SVG.

The renderer must not interpret `.ig` syntax and must not perform semantic layout.

## Output requirements

- valid SVG 1.1/modern browser-compatible SVG;
- UTF-8;
- deterministic output;
- scalable using `viewBox`;
- no external network assets;
- no JavaScript;
- no embedded executable content;
- no browser runtime assumptions.

## Recommended root

Conceptually:

```xml
<svg
  xmlns="http://www.w3.org/2000/svg"
  viewBox="0 0 1600 900"
  role="img"
  aria-labelledby="title desc">
```

## Accessibility

Where feasible, include:

- `<title>`;
- `<desc>`;
- semantic text as text rather than paths.

The visible title should remain searchable/selectable SVG text.

## Styling

The renderer receives resolved theme tokens.

It must not hard-code semantic color decisions across unrelated code paths.

## SVG primitives expected in v0.1

Likely sufficient:

- `<svg>`
- `<g>`
- `<rect>`
- `<line>` / `<path>`
- `<text>`
- `<tspan>`
- `<circle>` where useful
- `<defs>` for reusable safe definitions
- simple filters only if deterministic and broadly compatible

## Fonts

Do not embed font files in v0.1.

Use a documented fallback stack from the built-in theme.

SVG output must remain usable when the preferred font is unavailable.

## Icons

Icons must be rendered from bundled/local vector path data.

No remote icon fetching at render time.

v0.1 should contain a deliberately small icon registry sufficient for examples and tests.

Suggested initial names:

```text
message
route
decision
book
manual
boxes
gear
users
sparkles
```

Names are part of the public DSL once released, so keep the first set small.

## Determinism

Avoid random IDs.

If SVG definitions require IDs, derive stable IDs from document structure.

Golden tests should be possible byte-for-byte after normalization.
