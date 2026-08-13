# 13 — Instructions for Codex

## Mission

Implement the first MVP of `iasi-graphics` from these specifications.

Treat the documents in this pack as the product contract.

## Most important rule

Do not redesign the product while implementing it.

If a specification is ambiguous, choose the smallest implementation consistent with the stated principles and document the ambiguity.

Do not silently invent major features.

## Mandatory technology decisions

- Go for the engine.
- SVG as the first renderer.
- Lua only for the later Quarto adapter.
- No JavaScript.
- No Node.js.
- No npm.
- No PowerPoint/PPTX pipeline.
- No browser dependency.
- No runtime network dependency.

## Mandatory public interface decisions

- repository/product: `iasi-graphics`;
- source extension: `.ig`;
- CLI binary: `iasi-graphics`;
- QMD block: `{iasi-graphics}`.

## Implementation style

Prefer:

- small packages;
- explicit data structures;
- pure transformations where practical;
- deterministic behavior;
- table-driven Go tests;
- standard library first;
- readable error messages;
- minimal dependencies.

Avoid:

- premature frameworks;
- dependency-heavy SVG packages unless clearly justified;
- a universal graph engine;
- exposing geometry to the DSL;
- hidden mutation across compiler stages.

## Work order

Follow `11-implementation-plan.md`.

The first meaningful milestone is not "parser completed".

It is:

```text
examples/flow.ig
      ↓
iasi-graphics render
      ↓
flow.svg
```

with a visually convincing result.

## Required implementation notes

When Codex makes a non-trivial implementation choice not explicitly fixed by the specs, record it in a Markdown note under:

```text
docs/implementation-notes/
```

This is temporary lightweight decision persistence.

Do not introduce OpenSpec yet.

## Before claiming completion

Run at least:

```bash
go test ./...
go build ./cmd/iasi-graphics
```

Then render the three reference examples.

If Quarto integration has been reached, also render a QMD example containing an `{iasi-graphics}` block.

## Report back

At each completed phase, report:

- files created/changed;
- behavior implemented;
- tests added;
- commands run;
- remaining gaps;
- any spec ambiguity or deviation.

## Collaboration with VS Code + GitHub Copilot

Codex is not the only implementation tool.

After or between larger Codex work packages, the author may use VS Code + GitHub Copilot for interactive debugging, focused tests, small refactors and layout adjustments.

Both tools must follow the same specification pack.

See `14-development-workflow.md`.
