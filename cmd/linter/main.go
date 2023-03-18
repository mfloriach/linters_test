package main

import (
	"hoge/pkg/traceavility"

	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(traceavility.NewAnalyzer())
}
