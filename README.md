# iasi-graphics

Gráficos conceptuales declarativos compilados a SVG. La Fase 1 implementa el
primer vertical slice completo para `flow` en Go.

## Compilar y ejecutar

Requiere Go 1.26 o posterior.

```sh
cd src
go test ./...
go build -o ../bin/iasi-graphics ./cmd/iasi-graphics
../bin/iasi-graphics render examples/flow.ig -o examples/flow.svg
```

En Windows, usa `..\bin\iasi-graphics.exe` en el último comando.

El pipeline implementado es:

```text
.ig → lexer → parser → AST → validación/modelo semántico
    → layout de flow → escena neutral → renderer SVG → .svg
```

La Fase 1 excluye deliberadamente compare, ecosystem, stdin, Quarto,
empaquetado multiplataforma y temas extensibles.
