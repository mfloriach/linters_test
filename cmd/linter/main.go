package main

import (
	"hoge/pkg/architecture"

	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(architecture.NewAnalyzer())
}
