# Base API Product

Sistema de exemplo de como construir uma APi com Golang

> Requisitos do projeto:

- Go Lang >= 1.18

As demais dependências estão no arquivo go.mod e package.json

- https://go.dev/dl/

> Build do Back-End Go:
```bash
# Baixando as dependências
$ go mod tidy

# Compilar servidor HTTP
$ go build -o main cmd/product/main.go

# Ou compilar para outra plataforma ex: windows
$ GOOS=windows GOARCH=amd64 go build -o main64.exe cmd/product/main.go

# build modo production
$ go build -ldflags "-s -w" .
# Ou
$ go build -ldflags "-s -w" cmd/product/main.go
# Ou
$ go build -ldflags "-s -w" -o main cmd/product/main.go
```
## Opções de execução
- SRV_PORT (Porta padrão 8080)
- SRV_MODE (developer, homologation ou production / padrão production)

> Exemplo de Uso:
```bash
$ ./main.exe
# Ou
$ SRV_PORT=8080 SRV_MODE=developer ./main.exe
# Ou
$ SRV_PORT=9090 SRV_MODE=production ./main.exe
```

> Acesse:
- http://localhost:8080/api/v1/products
