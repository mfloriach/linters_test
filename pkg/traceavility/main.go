package traceavility

import (
	"flag"
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
)

const (
	ContextReport      = "must always propagate the context See: https://github.com/mfloriach/linters_test/tree/main/pkg/traceavility"
	ContextMysqlReport = "third parties (gorm) must contain `WithContext` See: https://github.com/mfloriach/linters_test/tree/main/pkg/traceavility"
)

//nolint:gochecknoglobals
var flagSet flag.FlagSet

func NewAnalyzer() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name:  "traceavility",
		Doc:   "mandatory to pass context to trace the request workflow",
		Run:   run,
		Flags: flagSet,
	}
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		if file.Name.String() == "repositories" || file.Name.String() == "services" {
			for _, declaration := range file.Decls {
				if function, ok := declaration.(*ast.FuncDecl); ok {

					if file.Name.String() == "repositories" && !mysqlWithContext(pass, function) {
						pass.Report(
							analysis.Diagnostic{
								Pos:     function.Pos(),
								Message: ContextReport,
							},
						)
					}

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

func mysqlWithContext(pass *analysis.Pass, function *ast.FuncDecl) bool {
	for _, stmts := range function.Body.List {
		if stmt, ok := stmts.(*ast.ExprStmt); ok {

			if _, ok := stmt.X.(*ast.CallExpr).Fun.(*ast.SelectorExpr).X.(*ast.SelectorExpr); ok {
				if stmt.X.(*ast.CallExpr).Fun.(*ast.SelectorExpr).X.(*ast.SelectorExpr).Sel.Name == "db" {
					return false
				}
			}

			if _, ok := stmt.X.(*ast.CallExpr).Fun.(*ast.SelectorExpr).X.(*ast.Ident); ok {
				continue
			}

			if stmt.X.(*ast.CallExpr).Fun.(*ast.SelectorExpr).X.(*ast.CallExpr).Fun.(*ast.SelectorExpr).X.(*ast.SelectorExpr).Sel.Name == "db" && stmt.X.(*ast.CallExpr).Fun.(*ast.SelectorExpr).X.(*ast.CallExpr).Fun.(*ast.SelectorExpr).Sel.Name != "WithContext" {
				return false
			}
		}
	}

	return true
}
