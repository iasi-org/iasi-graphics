# 05 — Semantic Model

## Purpose

The semantic model is the stable boundary between language syntax and visual layout.

Different future syntaxes should be able to produce the same semantic model.

Different future renderers should be able to consume layout results without knowing the source syntax.

## Model principles

The semantic model contains meaning, not geometry.

Allowed concepts include:

- document title;
- layout kind;
- semantic items;
- ordered flow structure;
- split groups;
- compare sides;
- ecosystem center/items;
- text;
- icon references;
- highlight.

It must not contain:

- `x`;
- `y`;
- SVG path data;
- font metrics;
- concrete colors;
- final pixel sizes.

## Suggested normalized model

Conceptually:

```text
Graphic
├── Kind: Flow | Compare | Ecosystem
├── Title
├── Elements
├── Relationships
└── Highlight
```

A normalized element may include:

```text
ID
Label
Body[]
Icon
Role
```

`Role` is internal semantic metadata such as:

```text
flow-step
compare-side
ecosystem-center
ecosystem-item
```

## Normalization

The semantic stage should normalize syntax-specific structures where useful.

Example:

```ig
split {
  step journey "Journey"
  step edr "EDR"
}
```

can become a semantic `Group` with:

```text
kind = split
children = [journey, edr]
```

## Validation

Semantic validation must run before layout.

No layout package should need to defend against malformed semantic structures.

## Error philosophy

Prefer a specific error over automatic repair.

Bad:

```text
ecosystem only has 2 items → silently invent spacing or duplicate an item
```

Good:

```text
ecosystem requires at least 3 items; found 2
```
