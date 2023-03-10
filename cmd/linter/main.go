package main

import (
	traceavility "hoge/pkg/modularization"

	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(traceavility.NewAnalyzer())
}
