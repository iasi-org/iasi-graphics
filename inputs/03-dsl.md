# 03 — DSL Specification

## Status

This is the **v0.1 language contract**.

The language is intentionally small.

File extension:

```text
.ig
```

Encoding:

```text
UTF-8
```

## Design goals

The DSL must be:

- readable without tooling;
- easy to write inside Markdown/Quarto code blocks;
- deterministic;
- easy to parse in Go;
- semantic rather than geometric;
- compact enough for figures to live beside prose.

## Lexical rules

### Strings

Double-quoted UTF-8 strings:

```ig
"Conversación"
"Cómo llegamos"
```

Escapes required in v0.1:

```text
\"  quote
\\  backslash
\n  newline
```

### Identifiers

Identifiers use:

```text
[A-Za-z_][A-Za-z0-9_-]*
```

Examples:

```text
journey
edr
knowledge-flow
item_1
```

### Comments

Single-line comments begin with `#`:

```ig
# This is a comment
```

### Blocks

Blocks use braces:

```ig
flow "Title" {
  ...
}
```

Braces are chosen for v0.1 because they are explicit, stable inside QMD code blocks and simple to parse.

Indentation is stylistic, not semantic.

## Top-level forms

Exactly one top-level visual is allowed per source in v0.1:

```text
flow
compare
ecosystem
```

---

# Flow

A `flow` communicates progression.

Example:

```ig
flow "Del diálogo al artefacto" {
  step conversation "Conversación" {
    text "Exploración y discusión"
    icon "message"
  }

  split {
    step journey "Journey" {
      text "Cómo llegamos"
      icon "route"
    }

    step edr "EDR" {
      text "Qué decidimos"
      icon "decision"
    }
  }

  step knowledge "Conocimiento" {
    icon "book"
  }

  step artifacts "Artefactos" {
    icon "boxes"
  }

  highlight "La experiencia y la decisión se convierten en conocimiento reutilizable"
}
```

Allowed direct children of `flow`:

- `step`
- `split`
- `highlight`

`split` contains two or more `step` items in v0.1.

A flow is ordered by source order.

---

# Compare

A `compare` communicates contrast or complementarity.

Example:

```ig
compare "Journey y EDR" {
  side journey "Journey" {
    text "Conserva el proceso"
    text "Alternativas, errores y evolución"
    icon "route"
  }

  side edr "EDR" {
    text "Conserva la decisión"
    text "Qué se decidió y por qué"
    icon "decision"
  }

  highlight "Dos formas complementarias de conservar conocimiento"
}
```

v0.1 requires exactly two `side` blocks.

Allowed direct children:

- exactly two `side`
- zero or one `highlight`

---

# Ecosystem

An `ecosystem` communicates a central concept surrounded by related elements.

Example:

```ig
ecosystem "El ecosistema IASI" {
  center "Ingeniería"

  item books "Libros" {
    text "Conocimiento consolidado"
    icon "book"
  }

  item manuals "Manuales" {
    text "Cómo utilizarlo"
    icon "manual"
  }

  item journey "Journey" {
    text "Proceso"
    icon "route"
  }

  item edr "EDR" {
    text "Decisiones"
    icon "decision"
  }

  item artifacts "Artefactos" {
    text "Implementación"
    icon "boxes"
  }
}
```

v0.1 requires:

- exactly one `center`;
- between 3 and 8 `item` blocks.

---

# Common nested statements

## text

```ig
text "Supporting sentence"
```

A container may include multiple `text` statements.

Text should be treated as short presentation copy, not arbitrary long prose.

## icon

```ig
icon "book"
```

Icons are optional.

Unknown icon names must produce a diagnostic, not silently access the network.

## highlight

```ig
highlight "Key message"
```

A highlight is a layout-level conclusion/emphasis element.

## IDs

IDs are local to the visual and must be unique.

Example:

```ig
step journey "Journey"
```

Here:

- `journey` is the identifier;
- `"Journey"` is the visible label.

## No style syntax in v0.1

The DSL must not support:

```text
color
font-size
x
y
width
height
stroke
fill
css
class
```

Visual style belongs to the theme.

## Grammar sketch

This is descriptive EBNF, not a requirement to use a parser generator.

```text
document       = flow | compare | ecosystem ;

flow           = "flow" string "{" flow_item* "}" ;
flow_item      = step | split | highlight ;
step           = "step" identifier string block? ;
split          = "split" "{" step step step* "}" ;

compare        = "compare" string "{"
                   side side highlight?
                 "}" ;
side           = "side" identifier string block? ;

ecosystem      = "ecosystem" string "{"
                   center item item item item*
                 "}" ;
center         = "center" string ;
item           = "item" identifier string block? ;

block          = "{" common_stmt* "}" ;
common_stmt    = text | icon ;

text           = "text" string ;
icon           = "icon" string ;
highlight      = "highlight" string ;
```

Semantic cardinality constraints are validated after parsing.
