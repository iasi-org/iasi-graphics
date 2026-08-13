# 02 — Architecture

## Overview

The engine is a compiler with a deliberately small public surface.

```text
Source (.ig or Quarto block)
          ↓
       Lexer
          ↓
       Parser
          ↓
         AST
          ↓
 Semantic validation
          ↓
  Semantic Model
          ↓
   Layout Engine
          ↓
    Scene Model
          ↓
    SVG Renderer
          ↓
         SVG
```

## Components

### 1. Source reader

Inputs:

- file;
- stdin;
- in-memory string for library/adaptor use.

Responsibilities:

- preserve source location;
- normalize line endings;
- provide diagnostics context.

### 2. Lexer/parser

Responsibilities:

- tokenize the language;
- parse layout-specific constructs;
- produce an AST;
- report syntax errors with line and column.

The parser must not calculate geometry.

### 3. Semantic validator

Responsibilities:

- enforce layout rules;
- reject duplicate identifiers;
- reject invalid child types;
- validate required structures;
- resolve optional icon names;
- produce meaningful user-facing diagnostics.

### 4. Semantic model

Renderer-independent description of meaning.

It must not contain SVG tags or final coordinates.

### 5. Layout engine

Responsibilities:

- choose canvas geometry;
- measure/estimate text;
- assign positions and dimensions;
- preserve hierarchy and whitespace;
- route connectors;
- avoid overlap;
- produce a Scene Model.

Each top-level layout has its own layout strategy.

v0.1 must not attempt to create a universal graph-layout algorithm.

### 6. Scene model

Renderer-neutral drawing scene.

Primitive scene elements may include:

- group;
- text;
- rounded card;
- line/connector;
- icon;
- background;
- emphasis block.

Coordinates are allowed here because this is an **internal model**.

### 7. SVG renderer

Transforms the Scene Model into deterministic SVG.

The SVG renderer must not make semantic decisions.

### 8. Theme

A theme provides visual tokens:

- typography sizes;
- font families/fallbacks;
- spacing scale;
- corner radii;
- stroke widths;
- palette;
- background treatment;
- connector treatment;
- card treatment;
- icon sizing.

v0.1 contains exactly one built-in theme: `iasi`.

## Suggested Go package structure

```text
iasi-graphics/
├── cmd/
│   └── iasi-graphics/
│       └── main.go
├── internal/
│   ├── source/
│   ├── lexer/
│   ├── parser/
│   ├── ast/
│   ├── semantic/
│   ├── model/
│   ├── layout/
│   │   ├── flow/
│   │   ├── compare/
│   │   └── ecosystem/
│   ├── scene/
│   ├── render/
│   │   └── svg/
│   ├── theme/
│   └── icons/
├── quarto/
│   └── _extensions/
│       └── iasi-graphics/
├── examples/
├── testdata/
└── docs/
```

This is a recommendation, not a requirement. Codex may improve package boundaries while preserving the architectural separations.

## Dependency policy

Prefer the Go standard library.

A third-party dependency is acceptable only if it eliminates substantial complexity and does not introduce a runtime dependency for users.

No Node or browser dependency is allowed.
