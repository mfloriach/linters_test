linters:
	go build -buildmode=plugin -o plugins plugins/arch.go
	golangci-lint cache clean