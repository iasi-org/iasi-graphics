# 06 — Layout Engine

## Product-critical component

The parser is necessary infrastructure.

The layout engine is where `iasi-graphics` becomes a product.

A technically correct SVG that looks like a generic graph diagram is not sufficient.

## Strategy

v0.1 uses **layout-specific algorithms**.

Implement separately:

- FlowLayout
- CompareLayout
- EcosystemLayout

Do not build a universal graph-layout engine.

## Default canvas

The initial internal design space should use a presentation-friendly 16:9 canvas:

```text
1600 × 900
```

The SVG should rely on `viewBox` so it scales responsively.

This is an internal default, not DSL syntax.

## Shared layout principles

### Whitespace

Whitespace is intentional.

Do not fill every available area.

### Hierarchy

The title, primary message, content cards and supporting text must have visibly different hierarchy.

### Minimum readable text

The engine must avoid shrinking text merely to make content fit.

If content exceeds sensible limits, return a diagnostic where possible.

### Balance

Parallel concepts should have balanced visual weight.

### Alignment

Use strong alignment lines and consistent spacing.

### Connectors

Connectors are secondary.

They must not dominate the composition.

Avoid crossings wherever possible.

### Cards

Cards are presentation components, not UML boxes.

Use generous padding and restrained borders.

### Density

A figure should communicate one main idea.

The engine is allowed to reject or warn about excessive density.

## Flow layout

Target visual grammar:

```text
        TITLE

      primary step
           ↓
      ┌─────────┐
      │  split  │
      └─────────┘
           ↓
      next step
           ↓
       outcome

     [highlight]
```

A split should form a balanced row of parallel cards.

After the split, the flow reconverges visually.

## Compare layout

Target visual grammar:

```text
             TITLE

       LEFT          RIGHT
     ┌───────┐     ┌───────┐
     │       │     │       │
     └───────┘     └───────┘

          [highlight]
```

Sides should be visually equal unless a future semantic feature explicitly changes emphasis.

## Ecosystem layout

Target visual grammar:

```text
           item

     item        item

         CENTER

     item        item

           item
```

Items should be distributed around a center with balanced angular/positional spacing.

The result should feel like an ecosystem or constellation, not a network graph.

## Text fitting

v0.1 may use deterministic heuristic text measurement if exact font metrics would introduce large dependencies.

However:

- wrapping rules must be deterministic;
- card sizes must remain consistent;
- layout tests must cover long labels;
- the implementation must isolate text measurement behind an interface so it can improve later.

## Scene output

Layout produces a renderer-neutral scene with concrete geometry.

Example concepts:

```text
Scene
├── Canvas
├── Text
├── Card
├── Connector
├── Icon
└── Highlight
```

Only after this stage are coordinates acceptable.
