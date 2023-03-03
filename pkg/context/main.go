package context

import (
	"flag"
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
)

const (
	ContextReport = "first parameter should be context See: https://github.com/mfloriach/linters_test/tree/main/pkg/context"
)

//nolint:gochecknoglobals
var flagSet flag.FlagSet

func NewContextAnalyzer() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name:  "context",
		Doc:   "want pass allways the context",
		Run:   run,
		Flags: flagSet,
	}
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		if file.Name.String() == "repositories" || file.Name.String() == "services" {
			for _, declaration := range file.Decls {
				if function, ok := declaration.(*ast.FuncDecl); ok {
					if !checkContextIsPropagate(function) {
						pass.Report(
							analysis.Diagnostic{
								Pos:     function.Pos(),
								Message: ContextReport,
							},
						)
					}
				}
			}
		}
	}
	return nil, nil
}

// check context is propagated around the workflow
func checkContextIsPropagate(fd *ast.FuncDecl) bool {
	// ignore constructor
	if strings.HasPrefix(fd.Name.Name, "New") {
		return true
	}

	if len(fd.Type.Params.List) == 0 {
		return false
	}

	field := fd.Type.Params.List[0]

	return isFirstFunctionParamNameCorrect(field, "ctx") && isFunctionParamTypeCorrect(field, "Context")
}

func isFirstFunctionParamNameCorrect(field *ast.Field, expectedName string) bool {
	return field.Names[0].Name == expectedName
}

func isFunctionParamTypeCorrect(field *ast.Field, expectedName string) bool {
	if _, ok := field.Type.(*ast.SelectorExpr); ok {
		if field.Type.(*ast.SelectorExpr).Sel.Name == expectedName {
			return true
		}
	}

	return false
}
