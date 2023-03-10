package swaggo

import (
	"flag"
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
)

//nolint:gochecknoglobals
var flagSet flag.FlagSet

const (
	SwagooReport = "should allways document your endpoints See: https://github.com/mfloriach/linters_test/tree/main/pkg/swagoo"
)

// conditions to trigger the linter
func swagooBadDocumented(function *ast.FuncDecl) bool {
	return !strings.HasPrefix(function.Name.Name, "New") &&
		function.Doc.Text() == "" &&
		!strings.Contains(function.Doc.Text(), "@Router") &&
		!strings.Contains(function.Doc.Text(), "@Summary") &&
		!strings.Contains(function.Doc.Text(), "@Description") &&
		!strings.Contains(function.Doc.Text(), "@Success")
}

func NewAnalyzer() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name:  "documentation",
		Doc:   "check don't forget document endpoints",
		Run:   run,
		Flags: flagSet,
	}
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		if file.Name.String() == "interfaces" {
			for _, declaration := range file.Decls {
				if function, ok := declaration.(*ast.FuncDecl); ok {
					if swagooBadDocumented(function) {
						pass.Report(
							analysis.Diagnostic{
								Pos:     function.Pos(),
								Message: SwagooReport,
							},
						)
					}
				}
			}
		}
	}
	return nil, nil
}
