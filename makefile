install: 
	docker build -t tools .

# start a server to inspect the packages compling 
graph: 
	docker run -it --rm -v $(pwd):/usr/src/app -p 7878:7878 --privileged tools go-callvis -skipbrowser cmd/main.go

# only works locally it requeri to install/compile golangci-lint locally
local-linter:
	go build -buildmode=plugin -o plugins plugins/arch.go
	golangci-lint cache clean

# generate architecture metrics (inestability, abstraction and distance)
metrics:
	docker run --rm -v $(PWD):/usr/src/app tools spm-go all --html cmd/main.go

# generate architecture compliance
arch:
	docker run --rm -v $(PWD):/usr/src/app tools arch-go --html

# runs golangci-lint inside docker
linter:
	docker run --rm -v $(pwd):/usr/src/app tools golangci-lint run

# usefull for testing new linters (better debugging experience)
test:
	docker run --rm -v $(pwd):/usr/src/app tools go run cmd/linter/main.go ./modules/...

test-coverage:
	go test -v ./... -coverprofile coverage/coverage.out
	go tool cover -html coverage/coverage.out -o coverage/coverage.html
	open coverage/coverage.html

report:
	go run cmd/metrics/main.go 
	open report.html

