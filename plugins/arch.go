package main

import (
	"hoge/pkg/architecture"
	force "hoge/pkg/forceNotNil"
	gin "hoge/pkg/ginReturnSerializer"
	hash "hoge/pkg/hashIds"
	"hoge/pkg/modularization"
	"hoge/pkg/swaggo"

	"hoge/pkg/traceavility"

	"golang.org/x/tools/go/analysis"
)

type analyzerPlugin struct{}

// This must be implemented
func (*analyzerPlugin) GetAnalyzers() []*analysis.Analyzer {
	return []*analysis.Analyzer{
		traceavility.NewAnalyzer(),
		force.NewAnalyzer(),
		gin.NewAnalyzer(),
		architecture.NewAnalyzer(),
		swaggo.NewAnalyzer(),
		modularization.NewAnalyzer(),
		hash.NewAnalyzer(),
	}
}

// This must be defined and named 'AnalyzerPlugin'
var AnalyzerPlugin analyzerPlugin
