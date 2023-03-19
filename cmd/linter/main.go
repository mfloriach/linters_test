package main

import (
	gin "hoge/pkg/ginReturnSerializer"

	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(gin.NewAnalyzer())
}
