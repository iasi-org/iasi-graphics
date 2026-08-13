# Phase 1 provisional behavior

These are implementation choices, not product contracts or acceptance criteria.

- Exact card sizes, gaps, connector paths, wrapping thresholds, typography, and
  icon artwork are provisional presentation choices (category A where internal;
  category C if proposed as stable visual acceptance behavior).
- The initial icon registry contains only the five names needed by the reference
  flow (`message`, `route`, `decision`, `book`, `boxes`). The final icon set is an
  unresolved category C decision (OQ-002).
- Phase 1 reports excessive vertical density instead of shrinking text. Exact
  content limits and diagnostic wording remain provisional category C behavior.
- The CLI validates the `.ig` suffix and supports only file input. Stdin, version,
  validate, and broader CLI compatibility are deliberately outside this phase;
  details not fixed by the existing CLI specification remain category C.

No golden SVG or visual acceptance snapshot freezes these provisional choices.
