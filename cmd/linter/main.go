package main

import (
	force "hoge/pkg/forceNotNil"

	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(force.NewForceNotNilAnalyzer())
}
