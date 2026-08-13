# 08 — CLI Specification

## Binary

```text
iasi-graphics
```

## MVP command

```bash
iasi-graphics render input.ig
```

Default output:

```text
input.svg
```

Explicit output:

```bash
iasi-graphics render input.ig -o output.svg
```

## stdin

Support stdin for adapters:

```bash
cat input.ig | iasi-graphics render - -o output.svg
```

This is especially useful for the Quarto Lua adapter.

## Theme

v0.1 has one theme and does not require the author to specify it.

An optional future-compatible flag may exist:

```bash
--theme iasi
```

but it must not complicate the MVP.

## Diagnostics

- errors go to stderr;
- successful SVG output goes to the specified file;
- exit code `0` on success;
- non-zero exit code on syntax, semantic or rendering failure.

## Helpful commands

v0.1 may also provide:

```bash
iasi-graphics version
iasi-graphics validate input.ig
```

`validate` is desirable but secondary to `render`.

## Explicit exclusions

Do not add in v0.1:

- server mode;
- watch mode;
- HTTP API;
- plugin installation;
- JavaScript output;
- PPTX output;
- PNG conversion;
- PDF conversion.

## Distribution goal

The final product should be distributable as a native Go binary for major platforms.

Packaging/release automation is not required for the first coding milestone.
