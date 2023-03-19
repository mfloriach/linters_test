package modularization

import (
	"flag"
	"strings"

	"golang.org/x/tools/go/analysis"
)

//nolint:gochecknoglobals
var flagSet flag.FlagSet

const (
	Package      = "main"
	PathModules  = "\"hoge/modules/"
	FacadeReport = "should use facade design pattern for modules See: https://github.com/mfloriach/linters_test/tree/main/pkg/modularization"
)

func NewAnalyzer() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name:  "modularization",
		Doc:   "should use facade pattern",
		Run:   run,
		Flags: flagSet,
	}
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		if file.Name.String() == Package {
			for _, imp := range file.Imports {
				if strings.Contains(imp.Path.Value, PathModules) && len(strings.Split(imp.Path.Value, "/")) > 3 {
					pass.Report(
						analysis.Diagnostic{
							Pos:     imp.Pos(),
							Message: FacadeReport,
						},
					)
				}
				// err := ast.Print(pass.Fset, imp.Path.Value)
				// if err != nil {
				// 	panic(err)
				// }
			}
		}
	}
	return nil, nil
}
