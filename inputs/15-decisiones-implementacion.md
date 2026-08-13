# Decisiones de implementación

## Propósito

Este documento define cómo deben gestionarse las decisiones de implementación al construir `iasi-graphics` a partir de sus especificaciones.

Su objetivo es mantener una frontera clara entre:

- detalles de implementación;
- decisiones de arquitectura;
- decisiones de especificación;
- criterios de aceptación.

La implementación no debe convertir silenciosamente una cuestión aún no resuelta en un contrato permanente del producto.

---

## Categorías de decisión

Toda ambigüedad relevante descubierta durante la implementación debe clasificarse en una de las siguientes categorías.

### A. Detalle de implementación

Una decisión que puede tomarse localmente durante la implementación.

Debe utilizarse la solución más sencilla que sea:

- determinista;
- reversible;
- fácil de sustituir;
- compatible con la especificación existente.

Ejemplos:

- geometría interna;
- orden de pintado;
- precisión numérica;
- estrategia de escape XML;
- heurísticas iniciales de layout;
- representación interna del SVG.

Estas decisiones no necesitan volver previamente a especificación mientras no se conviertan en comportamiento contractual o sean difíciles de revertir.

---

### B. Decisión de arquitectura

Una decisión que establece una frontera estructural interna del sistema y que, por tanto, debe tomarse deliberadamente y persistirse.

Ejemplos:

- estructura del modelo semántico;
- representación de `flow` y `split`;
- modelo de escena;
- fronteras entre tema, layout y renderer;
- abstracción para la medición de texto;
- estructura de módulos y paquetes.

Las decisiones arquitectónicas pueden evolucionar, pero no deben permanecer como conocimiento accidental contenido únicamente en el código.

Las decisiones arquitectónicas relevantes deben persistirse en la base de conocimiento del proyecto.

---

### C. Decisión de especificación o aceptación

Una decisión que afecta al comportamiento observable desde el exterior.

Incluye decisiones que modifican:

- el DSL;
- qué entradas son válidas o inválidas;
- el comportamiento de la línea de comandos;
- la compatibilidad;
- la salida generada considerada contractual;
- APIs públicas;
- criterios de aceptación visual;
- plataformas o formatos soportados.

Un agente de implementación no debe tomar unilateralmente estas decisiones.

Deben volver a la capa de especificación antes de convertirse en comportamiento contractual.

---

## Decisiones provisionales

Una ambigüedad no tiene por qué bloquear necesariamente la implementación.

Cuando una cuestión todavía no resuelta permita continuar trabajando, la implementación puede adoptar una decisión provisional.

Toda decisión provisional debe ser:

- mínima;
- determinista;
- reversible;
- explícitamente identificable como provisional.

El objetivo es permitir que la implementación avance sin ampliar prematuramente la especificación del producto.

---

## El comportamiento provisional no es contrato del producto

Toda implementación presenta necesariamente algún comportamiento, incluso en aquellas áreas que todavía no han sido completamente especificadas.

Por tanto, la regla no es:

> No implementar ningún comportamiento cuando exista una ambigüedad.

La regla es:

> **No convertir un comportamiento provisional en contrato del producto.**

Una decisión provisional de implementación no se convierte en contractual simplemente porque el código actual se comporte de esa manera.

Mientras no haya sido decidida explícitamente y persistida en la especificación, ese comportamiento no debe considerarse una garantía de compatibilidad.

---

## Congelación del contrato

Una decisión provisional no debe quedar congelada mediante mecanismos que la conviertan implícitamente en requisito del producto.

En particular, un comportamiento provisional no debe estabilizarse mediante:

- acceptance tests;
- golden files;
- documentación pública;
- APIs públicas;
- compromisos de compatibilidad;
- ejemplos normativos;
- acoplamiento arquitectónico difícilmente reversible.

Pueden existir tests unitarios que protejan la corrección interna de una implementación provisional cuando sean necesarios.

Pero esos tests no deben transformar una decisión de producto aún no resuelta en un criterio de aceptación.

---

## Acceptance Tests

Los acceptance tests representan el contrato del producto.

Responden a la pregunta:

> ¿Qué comportamiento observable debe seguir siendo cierto?

Por tanto, un acceptance test solo debe introducirse cuando el comportamiento correspondiente esté explícitamente especificado.

Los detalles internos de implementación pueden cambiar mientras los acceptance tests continúan siendo válidos.

Si modificar un detalle interno obliga a modificar un acceptance test, debe revisarse si ese test está verificando realmente comportamiento del producto o simplemente detalles de implementación.

---

## Regla de escalado

Siempre que una decisión de implementación afecte a:

- sintaxis o semántica del DSL;
- entradas aceptadas o rechazadas;
- comportamiento observable de la CLI;
- compatibilidad;
- contrato del renderer;
- comportamiento público de una integración;
- aceptación visual;
- criterios de aceptación;

la decisión debe volver a la capa de especificación antes de considerarse estable.

El agente de implementación debe identificar la ambigüedad, explicar sus consecuencias y continuar con trabajo reversible siempre que sea posible.

---

## Principio operativo

La implementación debe avanzar mediante el menor vertical slice completo permitido por la especificación existente.

Durante ese trabajo:

1. Las decisiones **A** pueden resolverse localmente utilizando la opción más sencilla, determinista y reversible.
2. Las decisiones **B** deben tomarse deliberadamente y persistirse como conocimiento arquitectónico.
3. Las decisiones **C** deben volver a especificación antes de convertirse en comportamiento contractual.
4. Las cuestiones aún no resueltas que no bloqueen el trabajo pueden utilizar implementaciones provisionales.
5. Las implementaciones provisionales nunca deben convertirse accidentalmente en contrato del producto.

**La implementación puede descubrir decisiones.**

**No puede definir silenciosamente el producto.**
