package swagoo

import (
	"flag"
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
)

//nolint:gochecknoglobals
var flagSet flag.FlagSet

const (
	SwagooReport = "should be swagger documented https://pkg.go.dev/golang.org/x/tools@v0.6.0/go/analysis#RelatedInformation.Message"
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

func NewSwagooAnalyzer() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name:  "swaggo",
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
