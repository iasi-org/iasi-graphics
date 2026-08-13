# 00 — Product Vision

## Problem

IASI uses technical diagrams such as PlantUML where precision, architecture and formal relationships matter.

A different class of figure is also needed: conceptual graphics with the visual language of a good presentation.

Examples include:

- conceptual flows;
- comparisons;
- ecosystem maps;
- explanatory figures;
- visual summaries;
- chapter-opening graphics;
- diagrams intended primarily to communicate an idea rather than formal architecture.

Existing diagramming DSLs can generate many of these structures, but they retain the visual grammar of technical diagrams.

Manual PowerPoint or Figma editing can produce attractive results, but it breaks reproducibility and moves complexity to the author.

Writing JavaScript or explicit SVG also breaks the IASI design principle because the author is forced to describe implementation details.

## Product goal

`iasi-graphics` must allow an author to write a compact declarative description of a visual idea and obtain a polished SVG.

Example:

```ig
flow "Del diálogo al artefacto" {
  step conversation "Conversación"

  split {
    step journey "Journey"
    step edr "EDR"
  }

  step knowledge "Conocimiento"
  step artifacts "Artefactos"
}
```

The author does not decide:

- coordinates;
- pixel dimensions;
- box sizes;
- connector geometry;
- typography metrics;
- SVG elements;
- PowerPoint shapes;
- JavaScript code.

Those decisions belong to the engine.

## Core promise

> The author describes **what the figure means**.  
> `iasi-graphics` decides **how to compose it visually**.

## Product identity

`iasi-graphics` is not:

- a vector editor;
- a presentation editor;
- a replacement for PlantUML;
- a general-purpose SVG programming language;
- a JavaScript drawing framework;
- a charting library.

It is a **declarative conceptual-graphics compiler**.

## Primary use case

A figure lives directly beside the text that explains it:

````markdown
```{iasi-graphics}
flow "Del diálogo al artefacto" {
  ...
}
```
````

The same language must also work standalone using `.ig` files.

## Success criterion for v0.1

The MVP succeeds if three source files can generate three presentation-quality SVG figures:

1. `flow`
2. `compare`
3. `ecosystem`

The generated graphics must feel editorial/presentation-oriented rather than like UML or generic graph-layout output.
