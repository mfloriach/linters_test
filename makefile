install: 
	docker build -t tools .

callvis: 
	docker run -it --rm -v $(pwd):/usr/src/app -p 7878:7878 --privileged tools go-callvis -skipbrowser cmd/main.go

custom:
	go build -buildmode=plugin -o plugins plugins/arch.go
	golangci-lint cache clean

coupling:
	docker run --rm -v $(pwd):/usr/src/app tools spm-go all --html cmd/main.go

arch:
	docker run --rm -v $(pwd):/usr/src/app tools arch-go --html

linter:
	docker run --rm -v $(pwd):/usr/src/app tools golangci-lint run

