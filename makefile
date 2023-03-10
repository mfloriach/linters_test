linters:
	go build -buildmode=plugin -o plugins plugins/arch.go
	golangci-lint cache clean

structure:
	docker compose run --rm web go-callvis -ignore hoge/pkg,github.com/gin-gonic/gin cmd/main.go

coupling:
	docker compose run --rm web spm-go all --html cmd/main.go

arch:
	docker compose run --rm web arch-go --html