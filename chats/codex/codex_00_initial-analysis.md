# Análisis inicial de `iasi-graphics`

Este análisis se basa exclusivamente en los documentos incluidos en `inputs/` y se realizó antes de iniciar la implementación.

## 1. Producto que hay que construir

`iasi-graphics` es un compilador de gráficos conceptuales declarativos.

El usuario escribe una descripción semántica breve en un archivo `.ig` o, posteriormente, en un bloque Quarto `{iasi-graphics}`. El programa interpreta esa descripción, compone automáticamente la figura y genera un SVG con calidad editorial o de presentación.

La promesa esencial es:

> El autor describe qué significa la figura; el motor decide cómo representarla visualmente.

No es un editor vectorial, una biblioteca de dibujo, un lenguaje SVG, una herramienta UML, un sistema de gráficos estadísticos ni un sustituto de PlantUML.

El producto debe ser:

- Un ejecutable nativo escrito en Go.
- Reproducible y determinista.
- Utilizable completamente offline.
- Independiente de JavaScript, Node.js, navegador, PowerPoint y servicios web.
- Capaz de producir gráficos presentacionales, no simples diagramas técnicos con cajas.

El MVP completo v0.1 incluye tres tipos de composición: `flow`, `compare` y `ecosystem`, además de una integración Quarto posterior al compilador standalone.

## 2. Capacidades principales

### Lenguaje declarativo

El DSL `.ig` permite expresar:

- Un único gráfico por documento.
- Títulos.
- Textos breves.
- Iconos opcionales de un registro incorporado.
- Mensajes destacados mediante `highlight`.
- Identificadores locales únicos.
- Comentarios con `#`.
- Strings UTF-8 con los escapes `\"`, `\\` y `\n`.

No permite coordenadas, dimensiones, colores, CSS ni primitivas gráficas.

### Tipos de gráfico

#### `flow`

- Secuencia ordenada de pasos.
- Grupos paralelos mediante `split`.
- Reconvergencia visual después de un split.
- Mensaje destacado.
- Pasos con etiqueta, textos e icono.

#### `compare`

- Exactamente dos lados.
- Peso visual equilibrado.
- Textos e iconos por lado.
- Cero o un `highlight`.

#### `ecosystem`

- Exactamente un centro.
- Entre tres y ocho elementos periféricos.
- Distribución equilibrada alrededor del centro.
- Apariencia de ecosistema o constelación, no de grafo genérico.

### Compilación y salida

Pipeline previsto:

```text
source
  → lexer
  → parser
  → AST
  → validación y normalización semántica
  → modelo semántico
  → layout específico
  → escena con geometría
  → renderer SVG
```

El resultado debe proporcionar:

- SVG válido y escalable mediante `viewBox`.
- Diseño interno inicial de `1600 × 900`.
- Texto real y seleccionable.
- `<title>` y, cuando sea viable, `<desc>`.
- Ausencia de scripts y recursos de red.
- IDs internos estables.
- Resultado determinista.
- Iconos vectoriales incorporados localmente.
- Una única tematización inicial: `iasi`.

### Interfaz de línea de comandos

Interfaz principal:

```bash
iasi-graphics render input.ig
iasi-graphics render input.ig -o output.svg
```

También debe aceptar stdin:

```bash
iasi-graphics render - -o output.svg
```

Los errores deben ir a stderr, incluir diagnóstico útil con archivo, línea y columna y producir un código de salida distinto de cero.

`version` y `validate` son deseables, pero secundarios frente a `render`.

### Quarto

Después de estabilizar el compilador standalone:

- Un filtro Lua detectará bloques `{iasi-graphics}`.
- Invocará el mismo ejecutable.
- Generará un SVG en una ubicación de build determinista.
- Sustituirá el bloque por una figura Pandoc/Quarto.
- Dará prioridad a `fig-cap`, `fig-alt` y `label`.

El adaptador no duplicará el parser, la validación, el layout ni el renderer.

## 3. Arquitectura inicial propuesta en Go

Mantendría las etapas especificadas, pero evitaría una fragmentación excesiva durante el primer slice:

```text
CLI
  → compiler
      → source
      → lexer/parser → AST
      → semantic → Graphic
      → layout/flow → Scene
      → svg → []byte
```

Propondría un pequeño paquete orquestador interno:

```go
func Compile(name string, source []byte, options Options) ([]byte, error)
```

No sería necesariamente una API pública en v0.1. Su función sería conectar etapas sin introducir decisiones semánticas propias.

Responsabilidades:

- `source`: nombre, contenido normalizado, posiciones y líneas para diagnósticos.
- `lexer`: tokens con rangos de origen.
- `parser`: AST fiel a la sintaxis.
- `ast`: estructuras sintácticas y posiciones.
- `semantic`: validación, resolución de iconos y normalización.
- `model`: modelo semántico sin geometría.
- `layout`: selección de estrategia por tipo de gráfico.
- `layout/flow`, `layout/compare`, `layout/ecosystem`: algoritmos independientes.
- `textmeasure`: medición heurística y wrapping deterministas mediante una interfaz aislada.
- `scene`: primitivas renderer-neutral con geometría.
- `theme`: tokens visuales resueltos del tema `iasi`.
- `icons`: registro pequeño de iconos y sus datos vectoriales locales.
- `render/svg`: serialización determinista de la escena.
- `diagnostic`: errores comunes con origen, línea, columna y mensaje.
- `compiler`: coordinación pura del pipeline.
- `cmd/iasi-graphics`: lectura de argumentos, archivos/stdin, escritura y códigos de salida.

Evitaría deliberadamente:

- Un motor universal de grafos.
- Interfaces abstractas para cada etapa antes de necesitarlas.
- Exponer paquetes públicos prematuramente.
- Hacer que el renderer conozca el AST.
- Hacer que el layout vuelva a validar estructuras inválidas.
- Incorporar una dependencia SVG si `encoding/xml` y escritura controlada son suficientes.

## 4. Mínimo vertical slice ejecutable

El primer slice debería ser un `flow` completo de referencia, no solamente un parser que “ya compila”.

Entrada:

```ig
flow "Del diálogo al artefacto" {
  step conversation "Conversación" {
    text "Exploración y discusión"
  }

  split {
    step journey "Journey" {
      text "Cómo llegamos"
    }

    step edr "EDR" {
      text "Qué decidimos"
    }
  }

  step knowledge "Conocimiento"
  step artifacts "Artefactos"

  highlight "La experiencia se convierte en conocimiento reutilizable"
}
```

Ejecución:

```bash
go run ./cmd/iasi-graphics render examples/flow.ig -o flow.svg
```

Debe atravesar realmente:

```text
archivo → lexer → parser → AST → validación → modelo semántico
        → FlowLayout → Scene → SVG → archivo
```

Este slice mínimo debe incluir:

- Lectura desde archivo.
- Tokens, strings, identificadores, comentarios y bloques.
- Parseo de `flow`, `step`, `split`, `text`, `icon` y `highlight`.
- Posiciones de origen.
- IDs duplicados y split con menos de dos pasos.
- Validación de iconos si aparecen.
- Layout vertical de pasos.
- Fila equilibrada para el split y reconvergencia.
- Título, tarjetas, conectores, texto y highlight.
- Canvas `1600 × 900` mediante `viewBox`.
- Theme `iasi` inicial.
- SVG determinista, accesible y sin recursos externos.
- Al menos pruebas de parser, validación, invariantes geométricas y validez XML.
- Inspección visual manual del SVG resultante.

No incluiría todavía `compare`, `ecosystem`, stdin, Quarto, empaquetado multiplataforma ni un sistema extensible de temas. Los iconos pueden limitarse a los necesarios para el ejemplo, aunque el registro debe quedar preparado para rechazar nombres desconocidos.

## 5. Estructura inicial propuesta

```text
iasi-graphics/
├── go.mod
├── cmd/
│   └── iasi-graphics/
│       └── main.go
├── internal/
│   ├── compiler/
│   ├── source/
│   ├── diagnostic/
│   ├── lexer/
│   ├── parser/
│   ├── ast/
│   ├── semantic/
│   ├── model/
│   ├── textmeasure/
│   ├── layout/
│   │   ├── layout.go
│   │   ├── flow/
│   │   ├── compare/
│   │   └── ecosystem/
│   ├── scene/
│   ├── theme/
│   ├── icons/
│   └── render/
│       └── svg/
├── examples/
│   ├── flow.ig
│   ├── compare.ig
│   └── ecosystem.ig
├── testdata/
│   ├── flow-basic.ig
│   ├── compare-basic.ig
│   ├── ecosystem-basic.ig
│   └── invalid/
├── docs/
│   └── implementation-notes/
└── quarto/
    └── _extensions/
        └── iasi-graphics/
```

Al principio solo crearía las carpetas y paquetes necesarios para el slice de `flow`. `compare`, `ecosystem` y `quarto` pueden añadirse cuando se alcancen sus fases; no necesitan paquetes vacíos.

## 6. Decisiones ambiguas o no especificadas

### DSL y validación

- No se establece si un `flow` debe contener un mínimo o máximo de pasos.
- No se fija cuántos `highlight` admite un `flow`; la gramática permite varios.
- No se define si `highlight` debe aparecer obligatoriamente al final.
- No se especifica si `center`, los dos `side` y los `item` deben respetar estrictamente el orden de la EBNF o si el parser acepta cualquier orden y la fase semántica valida cardinalidades.
- No se determina si un contenedor puede contener varios `icon`; el bloque común permite repetirlos conceptualmente, pero el modelo sugiere un solo icono.
- No se fijan límites concretos de longitud para título, etiqueta, texto o highlight.
- “Densidad excesiva” puede producir rechazo o advertencia, pero no se definen umbrales ni un mecanismo formal de warnings.
- No se aclara qué posiciones deben asociarse a errores semánticos compuestos.
- No se define si deben rechazarse tokens o contenido después del único gráfico raíz, aunque “exactamente uno” sugiere que sí.

### Modelo semántico y escena

- El esquema exacto del modelo semántico no es normativo.
- No se decide si usar un modelo normalizado común o tipos separados para cada layout.
- No se define formalmente la representación de relaciones, grupos y splits.
- Las primitivas y el orden de pintado de la escena no están cerrados.
- No se fija el contrato exacto entre theme, layout y renderer.
- No se especifica cómo representar clipping, sombras, bordes o conectores curvos.

### Diseño visual

- No están definidos los colores, tipografía, tamaños, espaciados, radios ni demás tokens concretos del tema `iasi`.
- La familia tipográfica final queda abierta.
- No existe una definición objetiva de “presentation-quality”; requiere revisión visual humana.
- No se concretan dimensiones de tarjetas, márgenes, áreas de título o highlight.
- No se define la heurística de medición y wrapping.
- No se decide cuándo el contenido debe ajustarse, envolverse o rechazarse.
- No se especifican reglas exactas para distribuir de tres a ocho elementos en `ecosystem`.
- No se define la geometría exacta de ramificación y reconvergencia de `split`.
- El conjunto inicial y el diseño vectorial de iconos no son definitivos; tampoco se especifican su procedencia o licencia.

### SVG

- No se fija si el root debe llevar `width` y `height` además de `viewBox`.
- No se define el contenido exacto de `<desc>`.
- No se establece una política concreta de escaping, precisión decimal o normalización del XML.
- “Byte-for-byte after normalization” no define qué normalización se aplicará.
- No se concreta el esquema para generar IDs internos estables.
- Los filtros visuales son opcionales, pero no hay criterio para decidir si usarlos.

### CLI

- No se especifica el path del módulo de Go.
- No se define el comportamiento si el archivo de salida ya existe.
- No se establece si `-o -` debe escribir el SVG en stdout.
- No se precisa cómo derivar el nombre de salida para stdin cuando falta `-o`.
- No se define si las extensiones distintas de `.ig` se rechazan.
- No se fija el formato estable de errores ni códigos de salida diferenciados.
- No se concreta la semántica de `validate` ni el formato de `version`.
- No queda cerrado si `--theme iasi` se incluye o se pospone.

### Quarto

- No está definida la estructura exacta de la extensión.
- No se decide la ubicación concreta de los SVG generados.
- No se fija el algoritmo de nombres o hashes.
- No se define cómo localizar el ejecutable.
- No se concreta cómo extraer las opciones `#|`.
- No se especifica el tratamiento de caché, limpieza o ejecuciones concurrentes.
- No se decide la relación final con `iasi-quarto` o `iasi-lua`.
- Tampoco se define el comportamiento exacto en HTML frente a formatos PDF cuando Quarto necesite convertir el SVG.

### Alcance y evolución

- El nombre del producto es provisionalmente fijo, pero no el nombre comercial del DSL.
- Temas personalizados, formatos alternativos, nuevos layouts e IA están explícitamente aplazados.
- No se define todavía una API Go pública para uso embebido.
- No se fija la estrategia de releases ni plataformas obligatorias.
- El documento de pruebas dice “cuatro niveles”, pero enumera cinco; la intención evidente es incluir el adaptador Quarto como quinto nivel posterior.

## Recomendación

Durante el primer slice deberían resolverse únicamente las ambigüedades locales necesarias, usando la opción más pequeña y reversible, y documentarlas en `docs/implementation-notes/`.

Las cuestiones que cambien el DSL público, el contrato del renderer, la integración Quarto o los criterios de aceptación deberían volver antes a la capa de especificación.
