package gin

import (
	"flag"
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
)

const (
	ReturnObjectSerializer = "must return and object from serializer See: https://github.com/mfloriach/linters_test/tree/main/pkg/ginReturnSerializer"
)

//nolint:gochecknoglobals
var flagSet flag.FlagSet

func NewAnalyzer() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name:  "serializer",
		Doc:   "its must return and serialized objects",
		Run:   run,
		Flags: flagSet,
	}
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		if file.Name.String() == "interfaces" {
			for _, dlc := range file.Decls {
				if _, ok := dlc.(*ast.FuncDecl); !ok {
					continue
				}

				if strings.HasPrefix(dlc.(*ast.FuncDecl).Name.Name, "New") == true {
					continue
				}

				for _, stmt := range dlc.(*ast.FuncDecl).Body.List {
					if _, ok := stmt.(*ast.ExprStmt); !ok {
						continue
					}

					if stmt.(*ast.ExprStmt).X.(*ast.CallExpr).Fun.(*ast.SelectorExpr).Sel.Name != "JSON" {
						continue
					}

					if _, ok := stmt.(*ast.ExprStmt).X.(*ast.CallExpr).Args[1].(*ast.CompositeLit).Elts[0].(*ast.KeyValueExpr).Value.(*ast.CallExpr); !ok {
						pass.Report(
							analysis.Diagnostic{
								Pos:     stmt.(*ast.ExprStmt).X.(*ast.CallExpr).Args[1].(*ast.CompositeLit).Elts[0].(*ast.KeyValueExpr).Value.Pos(),
								Message: ReturnObjectSerializer,
							},
						)
					}
				}

			}
		}
	}
	return nil, nil
}
