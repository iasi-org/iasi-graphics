# Principio IASI: integraciones reemplazables

Fecha: 2026-08-13

## Contexto

Tras completar la Fase 1 de `iasi-graphics`, surgió la cuestión de utilizar
OpenSpec en el proceso de ingeniería. La conversación aclaró que el problema no
era decidir si IASI debía aceptar o rechazar OpenSpec, sino distinguir la
filosofía de IASI de las herramientas empleadas para aplicarla.

La analogía inicial fue `iasi-quarto` y Quarto:

- IASI utiliza Quarto como tecnología base.
- `iasi-quarto` añade controles, convenciones y garantías propias.
- Si Quarto desapareciera o dejara de ser adecuado, sería posible construir una
  integración diferente sin cambiar la filosofía de IASI.

El mismo razonamiento se aplica a OpenSpec:

- OpenSpec puede utilizarse como mecanismo estándar.
- IASI puede exigir controles adicionales mediante una capa superior.
- El proceso de IASI no debe quedar definido ni cautivo por OpenSpec.
- Si se adopta otra tecnología en el futuro, debe cambiar la integración, no los
  principios ni las garantías de IASI.

## Distinción conceptual

```text
Principios y garantías IASI
            ↓
Requisitos y controles propios
            ↓
Capa de integración reemplazable
            ↓
Herramienta externa concreta
```

La capa IASI puede expandir una herramienta externa cuando falten capacidades y
restringirla cuando IASI requiera más seguridad. Por ejemplo, en el proceso de
ingeniería podrían conservarse independientemente de OpenSpec:

- clasificación de decisiones A/B/C;
- distinción entre comportamiento provisional y contractual;
- autoridad humana y de agentes;
- trazabilidad del origen de las decisiones;
- implementación mediante vertical slices;
- evidencia de verificación;
- reglas de cierre;
- prevención de contratos accidentales.

## Decisión conceptual

> IASI adopta herramientas externas mediante capas de integración reemplazables.
> Los principios, garantías y contratos propios permanecen independientes de la
> tecnología integrada.

Esto evita confundir una herramienta con el método y evita que sus limitaciones
definan permanentemente cómo trabaja IASI.

## Persistencia como principio transversal

La decisión se incorporó a los principios generales del workspace en:

- `C:\iasi-org\README_es.md`
- `C:\iasi-org\README_en.md`

Formulación en español:

> Los principios y garantías de IASI deben ser independientes de las herramientas
> concretas. Las tecnologías externas se adoptan mediante capas de integración
> reemplazables, de modo que puedan evolucionar o sustituirse sin alterar la
> filosofía de IASI.

## Trabajo futuro

La conversación no define todavía una capa IASI para OpenSpec ni su nombre. Solo
establece el principio que deberá guiar su posible diseño. Antes de construirla,
habrá que distinguir qué resuelve ya OpenSpec, qué controles adicionales requiere
IASI y qué partes pertenecen al proceso general en lugar de a la integración.
