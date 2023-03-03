package main

import (
	"hoge/pkg/context"
	gin "hoge/pkg/ginReturnSerializer"
	"hoge/pkg/imports"
	"hoge/pkg/swagoo"

	"golang.org/x/tools/go/analysis"
)

type analyzerPlugin struct{}

// This must be implemented
func (*analyzerPlugin) GetAnalyzers() []*analysis.Analyzer {
	return []*analysis.Analyzer{
		imports.NewImportsAnalyzer(),
		swagoo.NewSwagooAnalyzer(),
		context.NewContextAnalyzer(),
		gin.NewGinReturnAnalyzer(),
	}
}

// This must be defined and named 'AnalyzerPlugin'
var AnalyzerPlugin analyzerPlugin
