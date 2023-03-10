package traceavility

import (
	"flag"
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
)

const (
	ContextReport = "must always propagate the context See: https://github.com/mfloriach/linters_test/tree/main/pkg/context"
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

					// if file.Name.String() == "repositories" {
					// 	mysqlWithContext(pass, function)
					// }

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

func mysqlWithContext(pass *analysis.Pass, function *ast.FuncDecl) {
	for _, stmts := range function.Body.List {
		if stmt, ok := stmts.(*ast.ExprStmt); ok {

			// stmt.X.(*ast.CallExpr).Fun.(*ast.SelectorExpr).X.(*ast.CallExpr).Fun
			err := ast.Print(pass.Fset, stmt.X.(*ast.CallExpr).Fun.(*ast.SelectorExpr).X.(*ast.CallExpr).Fun.(*ast.SelectorExpr).Sel)
			if err != nil {
				panic(err)
			}

			// line := stmt.X.(*ast.CallExpr).Fun.(*ast.SelectorExpr).X.(*ast.SelectorExpr).X.(*ast.Ident) // desde aqui

			// object := line.Obj.Decl.(*ast.Field).Type.(*ast.Ident).Obj.Decl.(*ast.TypeSpec).Type.(*ast.StructType).Fields.List[0].Type.(*ast.SelectorExpr)
			// // using injected
			// if line.Name == "r" && object.Sel.Name == "DB" && object.X.(*ast.Ident).Name == "gorm" {
			// 	if stmt.X.(*ast.CallExpr).Fun.(*ast.SelectorExpr).Sel.Name == "WithContext" {
			// 		fmt.Println("sdfgsdfdsfds")
			// 	}

			// }

		}
	}
}
