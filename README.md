# swaggo-demo

Projeto mínimo em [Go](https://go.dev) para demonstrar geração de documentação utilizando [swaggo](https://github.com/swaggo/swag). 

## Pre-requisitos

- Go 1.21+
- swag CLI instalado:

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

Garanta que `$GOPATH/bin` (ou `$HOME/go/bin`) esteja no seu PATH para o comando `swag` funcionar.

## Passo a passo

```bash
# 1. baixar dependencias
go mod tidy

# 2. gerar a pasta docs/ a partir dos comentários no código
swag init

# 3. rodar a API
go run main.go
```

Depois, acesse:

- Swagger UI: http://localhost:8080/swagger/index.html
- JSON gerado: http://localhost:8080/swagger/doc.json


## Arquivos Importantes

- `main.go` - bloco de anotações gerais (`@title`, `@version`, `@host`, `@BasePath`) e montagem das rotas do Swagger UI.
- `handlers/task.go` - anotações por endpoint (`@Summary`, `@Param`, `@Success`, `@Failure`, `@Router`).
- `models/task.go` - structs usadas nos exemplos de request/response, com tags `example`.
- `docs/` - gerado pelo `swag init`.
