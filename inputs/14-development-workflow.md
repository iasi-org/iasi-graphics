# 14 — Development Workflow

## Purpose

This document defines how implementation work is distributed between human design work and AI coding assistants.

The specification is the source of truth.

No coding assistant owns the architecture.

## Roles

### ChatGPT + author

Responsibilities:

- discuss the problem;
- define product intent;
- challenge alternatives;
- make engineering decisions;
- write and refine specifications;
- define acceptance tests;
- review deviations from the specification.

Outputs:

- Markdown specifications;
- acceptance criteria;
- engineering decisions;
- implementation guidance.

### Codex

Primary role:

- implement larger, well-scoped work packages from the specifications;
- work across multiple files;
- create complete vertical slices;
- run tests and builds;
- report deviations, ambiguities and remaining gaps.

Typical tasks:

```text
Implement Phase 1 from 11-implementation-plan.md
Implement the flow parser + semantic model + layout + SVG vertical slice
Add the compare layout according to 03-dsl.md and 06-layout-engine.md
```

Codex must treat the specification pack as the contract.

### VS Code + GitHub Copilot

Primary role:

- interactive local development;
- fine-grained implementation work;
- debugging;
- completing or adjusting functions;
- writing focused tests;
- small refactors;
- investigating failures;
- making visual/layout tweaks while the author is inspecting the code.

Typical tasks:

```text
Fix this failing Go test
Refactor this function without changing behavior
Add tests for duplicate IDs
Adjust this spacing calculation
Explain why this SVG text is overflowing
```

Copilot must also follow the same specifications.

## Workflow

Recommended flow:

```text
ChatGPT + author
        ↓
specifications / decisions
        ↓
Codex
        ↓
major implementation
        ↓
VS Code + Copilot
        ↓
interactive refinement / debugging
        ↓
tests
        ↓
GitHub
```

This is not a rigid pipeline.

A task may be assigned directly to Copilot when it is small or highly interactive.

A task may return to ChatGPT + author when implementation exposes an architectural ambiguity.

## Core rule

> The implementation agent may change. The specification must survive the agent.

Codex and Copilot are tools that execute against the same persisted engineering knowledge.

## Escalation rule

If either Codex or Copilot encounters a question that changes:

- public DSL;
- architecture;
- product scope;
- renderer contract;
- Quarto integration contract;
- acceptance criteria;

it must not silently invent the answer.

The issue must be brought back to the specification/decision layer.

## GitHub

GitHub is the final shared state of the implementation.

Changes made through Codex or Copilot must ultimately be visible as ordinary repository changes, tests and commits.

## No tool coupling

`iasi-graphics` must not depend on Codex, Copilot or ChatGPT at runtime.

These tools are part of the engineering workflow, not part of the product architecture.
