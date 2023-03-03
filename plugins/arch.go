package main

import (
	"hoge/pkg/context"
	force "hoge/pkg/forceNotNil"
	gin "hoge/pkg/ginReturnSerializer"
	"hoge/pkg/imports"
	"hoge/pkg/swagoo"

	"golang.org/x/tools/go/analysis"
)

type analyzerPlugin struct{}

// This must be implemented
func (*analyzerPlugin) GetAnalyzers() []*analysis.Analyzer {
	return []*analysis.Analyzer{
		context.NewContextAnalyzer(),
		force.NewForceNotNilAnalyzer(),
		gin.NewGinReturnAnalyzer(),
		imports.NewImportsAnalyzer(),
		swagoo.NewSwagooAnalyzer(),
	}
}

// This must be defined and named 'AnalyzerPlugin'
var AnalyzerPlugin analyzerPlugin
