# Clasificación de ambigüedades de `iasi-graphics`

Este documento clasifica las ambigüedades encontradas en la especificación de `inputs/` antes de iniciar la implementación.

## Categorías

- **A. Implementation detail:** puede decidirse durante la implementación utilizando la opción más simple, determinista y reversible.
- **B. Architecture decision:** debe decidirse y persistirse antes o durante la implementación, pero no modifica el contrato público del producto.
- **C. Specification / acceptance decision:** no debe decidirse unilateralmente porque cambia comportamiento observable, DSL, compatibilidad o criterios de aceptación.

“Bloquea el primer vertical slice” significa que impide completar el recorrido ejecutable `flow.ig → SVG`. Una decisión puede ser importante para el MVP completo sin bloquear ese primer flujo positivo.

## DSL y validación

| Ambigüedad | Categoría | Por qué | ¿Bloquea el primer slice? |
|---|---|---|---|
| Mínimo y máximo de elementos de un `flow` | C | Determina qué fuentes `.ig` son válidas y cambia el contrato observable del DSL. | No. El ejemplo puede ser válido sin decidir los límites generales. |
| Número máximo de `highlight` en `flow` | C | Aceptar varios, rechazar el segundo o aplicar una precedencia cambia la validez y el significado del DSL. | No. El slice puede usar exactamente uno y aplazar los casos múltiples. |
| Posición obligatoria de `highlight` dentro de `flow` | C | Obligar a colocarlo al final afecta compatibilidad sintáctica y comportamiento observable. | No. El ejemplo puede colocarlo al final. |
| Orden estricto o flexible de `center`, `side`, `item` y `highlight` | C | Determina qué documentos acepta el lenguaje. La EBNF sugiere un orden, pero el texto habla también de cardinalidades semánticas. | No para `flow`; tampoco hace falta resolverlo hasta `compare` y `ecosystem`. |
| Permitir varios `icon` dentro de un bloque | C | Cambia la cardinalidad y semántica pública de una construcción del DSL. | No. El slice puede utilizar cero o un icono sin definir el caso múltiple. |
| Límites de longitud para título, etiquetas, textos y highlight | C | Un límite convierte entradas antes válidas en errores o warnings y afecta aceptación. | No para una entrada breve; sí deberá resolverse antes de prometer validación de densidad. |
| Umbrales de “densidad excesiva” y si generan error o warning | C | Afecta códigos de salida, diagnósticos y aceptación de documentos. | No. El primer ejemplo puede mantenerse claramente dentro de una densidad razonable. |
| Posición que debe señalar un error semántico compuesto | A | Es una elección diagnóstica interna mientras siempre se proporcione una ubicación útil y determinista. | No. Se puede escoger el nodo causante o la apertura del contenedor. |
| Tratamiento de tokens después del gráfico raíz | C | Aceptarlos o rechazarlos modifica la gramática observable. “Exactamente un visual” sugiere rechazo, pero conviene fijarlo explícitamente. | No para el camino válido; sí para cerrar las pruebas de errores del parser. |

## Modelo semántico y escena

| Ambigüedad | Categoría | Por qué | ¿Bloquea el primer slice? |
|---|---|---|---|
| Esquema exacto del modelo semántico | B | Es una frontera arquitectónica entre sintaxis y layout, sin necesidad de exponerla públicamente. | Sí. Se necesita una representación concreta para conectar validación y layout. |
| Modelo normalizado común frente a tipos separados por layout | B | Afecta extensibilidad, acoplamiento y organización interna, pero no el DSL ni el SVG contractual. | Sí. Al menos debe adoptarse una forma inicial para `flow`. |
| Representación de relaciones, grupos y `split` | B | Define el contrato interno entre semantic model y layout. | Sí. El layout de `flow` necesita comprender orden y grupos paralelos. |
| Primitivas exactas de la escena | B | La escena es la frontera entre layout y renderer; su diseño afecta paquetes y evolución interna. | Sí. Hace falta un mínimo de texto, tarjeta, conector, grupo y highlight. |
| Orden de pintado de los elementos de la escena | A | Es una decisión renderer/layout local y reversible mientras produzca un resultado correcto y determinista. | Sí, pero puede resolverse durante la implementación. |
| Contrato exacto entre theme, layout y renderer | B | Determina qué capa resuelve colores, dimensiones, estilos y tokens; una mala frontera genera acoplamiento estructural. | Sí. Debe existir una división mínima antes de conectar layout y SVG. |
| Representación de clipping, sombras, bordes y conectores curvos | A | Son técnicas internas de representación, sustituibles sin cambiar el DSL. | No. El slice puede comenzar sin clipping ni sombras y con conectores simples. |

## Diseño visual

| Ambigüedad | Categoría | Por qué | ¿Bloquea el primer slice? |
|---|---|---|---|
| Paleta, tamaños tipográficos, espaciado, radios y demás tokens del tema `iasi` | C | El tema incorporado forma parte directa del resultado observable y de la aceptación visual del producto. | Sí para declarar el slice “visualmente convincente”. No para producir un primer SVG técnico. |
| Familia y stack tipográfico inicial | C | Afecta apariencia, métricas, wrapping y consistencia entre plataformas. Además, la especificación deja pendiente el estándar tipográfico IASI. | Sí para aceptación visual; no para una prueba técnica provisional. |
| Definición objetiva de “presentation-quality” | C | Es un criterio explícito de aceptación, actualmente dependiente de revisión humana. Codificar un criterio propio cambiaría el umbral de éxito. | Sí para afirmar que la fase visual está completada. |
| Dimensiones de tarjetas, márgenes y áreas de título/highlight | A | Son parámetros internos ajustables y reversibles dentro de un canvas ya fijado. | Sí, pero pueden elegirse durante el layout inicial y refinarse. |
| Heurística de medición y wrapping | B | Debe aislarse detrás de una interfaz y afecta de forma transversal todos los layouts. El algoritmo concreto puede evolucionar. | Sí. Hace falta una primera estrategia determinista para colocar texto. |
| Cuándo envolver, ajustar o rechazar contenido | C | Rechazar o transformar texto afecta entradas aceptadas y resultado observable. | No para textos breves; sí antes de aceptar casos largos como comportamiento estable. |
| Distribución de 3 a 8 elementos de `ecosystem` | A | Es un algoritmo específico de layout y puede cambiar sin alterar el DSL. | No para el slice de `flow`. |
| Geometría de ramificación y reconvergencia de `split` | A | La semántica ya está fijada; la ruta y forma concreta de los conectores es interna. | Sí, pero se puede decidir y ajustar durante la implementación. |
| Conjunto definitivo de iconos | C | Los nombres publicados pasan a formar parte del DSL y afectan compatibilidad futura. | No. Se puede omitir el icono en el primer ejemplo o usar únicamente nombres ya sugeridos sin declarar definitivo el registro completo. |
| Diseño vectorial, procedencia y licencia de los iconos | B | Afecta activos incorporados, distribución y riesgos legales, aunque no la sintaxis si los nombres permanecen iguales. | No si el primer slice no usa iconos; sí antes de distribuir iconos incorporados. |

En la heurística de texto hay dos decisiones: la existencia y forma de la abstracción es B; los coeficientes concretos usados inicialmente para estimar anchura son A.

## Renderer SVG

| Ambigüedad | Categoría | Por qué | ¿Bloquea el primer slice? |
|---|---|---|---|
| Incluir `width` y `height` en el `<svg>` además de `viewBox` | C | Cambia el comportamiento de incrustación y tamaño por defecto del SVG en consumidores. | No para generar SVG escalable; sí conviene decidirlo antes de congelar golden files. |
| Contenido exacto de `<desc>` | A | Puede derivarse de manera determinista del contenido sin introducir nueva semántica. | No. Puede usarse una descripción mínima o posponerse, pues se pide “cuando sea viable”. |
| Política de escaping XML | A | Debe ser correcta, pero es una decisión de serialización interna; lo natural es usar capacidades estándar de Go. | Sí, pero no requiere una decisión de producto. |
| Precisión decimal de coordenadas | A | Es interna, determinista y reversible mientras no cause defectos visibles. | Sí, pero puede fijarse localmente. |
| Formato y normalización del XML | B | Determina estabilidad de snapshots y reproducibilidad entre versiones del renderer. | No para el primer render; sí antes de establecer golden tests byte a byte. |
| Esquema para generar IDs internos estables | A | Los IDs son internos mientras permanezcan deterministas, válidos y no colisionen. | No si el SVG inicial no necesita definiciones con IDs. |
| Uso de filtros visuales | A | La especificación los hace opcionales. Puede comenzarse sin filtros. | No. La opción más simple es no utilizarlos inicialmente. |

## CLI y distribución

| Ambigüedad | Categoría | Por qué | ¿Bloquea el primer slice? |
|---|---|---|---|
| Path del módulo Go | B | Es identidad técnica persistida en `go.mod` e imports internos; no altera el DSL ni la CLI. | Sí. Es necesario para inicializar y construir el módulo. |
| Sobrescritura de un archivo de salida existente | C | Es comportamiento observable y puede implicar pérdida de datos o exigir confirmación. | No si la prueba escribe a un destino nuevo. |
| Soporte de `-o -` para stdout | C | Amplía la interfaz pública de la CLI y afecta cómo los adaptadores consumen la salida. | No. La especificación solo exige stdin con salida a archivo. |
| Salida por defecto al usar stdin sin `-o` | C | No existe un nombre de entrada del que derivar el SVG; elegir error, stdout o nombre fijo cambia la CLI. | No. El slice inicial lee un archivo y la integración prevista puede proporcionar `-o`. |
| Rechazar entradas cuya extensión no sea `.ig` | C | Afecta qué invocaciones acepta la CLI. | No. El ejemplo puede usar `.ig`. |
| Formato estable de diagnósticos | C | El requisito mínimo está fijado, pero convertir el formato completo en contrato afecta herramientas y compatibilidad. | Parcialmente: sí hace falta `archivo:línea:columna: mensaje`; no hace falta fijar todavía más estructura. |
| Códigos de salida diferenciados por tipo de error | C | Es una interfaz observable para scripts. La especificación solo exige cero/no cero. | No. Basta inicialmente con `0` y un único valor no cero. |
| Semántica de `validate` | C | Define una orden pública todavía marcada como secundaria. | No. Debe posponerse. |
| Formato de `version` | C | Es salida pública consumible por usuarios y scripts. | No. Debe posponerse. |
| Inclusión de `--theme iasi` en v0.1 | C | Es una opción pública y la especificación dice que puede existir, no que deba hacerlo. | No. La opción mínima es no exponerla todavía. |

## Integración Quarto

| Ambigüedad | Categoría | Por qué | ¿Bloquea el primer slice? |
|---|---|---|---|
| Estructura exacta de la extensión Quarto | B | Es organización interna del adaptador mientras preserve el bloque público `{iasi-graphics}`. | No. Quarto se implementa después del core. |
| Ubicación de los SVG generados | C | Es comportamiento observable del build y afecta limpieza, publicación y tooling. | No. |
| Algoritmo de nombres o hashes | B | Es una decisión persistente del adaptador para determinismo, caché y colisiones. | No. |
| Localización del ejecutable `iasi-graphics` | C | Afecta instalación, PATH, configuración y mensajes de error visibles. | No. |
| Extracción de opciones `#|` | B | Debe seguir las convenciones de Quarto; la técnica concreta pertenece al adaptador. | No. |
| Política de caché | C | Cambia cuándo se recompila y qué percibe el usuario. | No. |
| Limpieza de artefactos generados | C | Afecta archivos observables y posibles eliminaciones. | No. |
| Ejecuciones concurrentes | B | Requiere una estrategia interna segura para nombres y escritura atómica. | No. |
| Relación con `iasi-quarto` o `iasi-lua` | B | Define propiedad, empaquetado y dependencia entre repositorios sin cambiar necesariamente el bloque público. | No. |
| Comportamiento HTML frente a PDF | C | Afecta formatos soportados y criterios de aceptación de la integración. | No. |

## Alcance, evolución y pruebas

| Ambigüedad | Categoría | Por qué | ¿Bloquea el primer slice? |
|---|---|---|---|
| Nombre comercial del DSL | C | Es identidad pública y compatibilidad documental. | No. El producto y la extensión `.ig` ya bastan. |
| Temas personalizados | C | Introducen nueva sintaxis o configuración pública. Están explícitamente fuera de v0.1. | No. |
| Formatos PNG/PDF y otros renderers | C | Amplían el contrato de salida y están fuera de v0.1. | No. |
| Nuevos layouts | C | Amplían el lenguaje público. | No. |
| Generación mediante IA | C | Amplía alcance y workflow del producto; está fuera del renderer y del MVP. | No. |
| Existencia de una API Go pública | C | Una API exportada crea un contrato de compatibilidad externo. | No. El orquestador puede permanecer en `internal/`. |
| Estrategia de releases | B | Afecta distribución y mantenimiento, no el comportamiento del DSL. | No. |
| Plataformas obligatorias | C | Define compatibilidad y aceptación del producto distribuido. | No para el slice local. |
| “Cuatro niveles” de pruebas aunque se enumeran cinco | A | Es una inconsistencia editorial evidente: los cinco niveles enumerados pueden conservarse, con Quarto como fase posterior. | No. |

## Resumen de bloqueos reales

Para comenzar y terminar técnicamente el primer vertical slice hay que tomar las siguientes decisiones.

### B — Deben persistirse

- Path del módulo Go.
- Forma inicial del modelo semántico.
- Representación de `split` y orden del flujo.
- Primitivas mínimas de la escena.
- Contrato entre theme, layout y renderer.
- Abstracción de medición de texto.

### A — Se pueden resolver durante la implementación

- Orden de pintado.
- Geometría de tarjetas y conectores.
- Márgenes y dimensiones iniciales.
- Precisión numérica.
- Escaping XML.
- Heurística concreta de medición.
- Representación SVG simple, sin filtros.

### C — Impiden declarar el resultado visual como aceptado

- Tokens concretos del tema `iasi`.
- Stack tipográfico.
- Umbral de “presentation-quality”.

Las demás decisiones C no bloquean el slice positivo si se evita convertir los casos ambiguos en comportamiento estable. Se pueden posponer manteniendo pruebas únicamente para lo expresamente definido.

## Criterio operativo propuesto

Implementar el camino válido especificado, no añadir aceptación ni rechazo explícito para casos ambiguos salvo lo indispensable, y no congelar golden SVGs hasta resolver las decisiones visuales de categoría C.
