package gin

import (
	"flag"
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

const (
	ReturnJSON             = "you must return a JSON object https://github.com/mfloriach/linters_test/tree/main/pkg/ginReturnSerializer"
	ReturnObjectSerializer = "must return and object from serializer See: https://github.com/mfloriach/linters_test/tree/main/pkg/ginReturnSerializer"
)

//nolint:gochecknoglobals
var flagSet flag.FlagSet

func NewGinReturnAnalyzer() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name:  "gin_rules",
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

				if _, ok := dlc.(*ast.FuncDecl).Body.List[0].(*ast.ExprStmt); !ok {
					continue
				}

				ginReturn := dlc.(*ast.FuncDecl).Body.List[0].(*ast.ExprStmt).X.(*ast.CallExpr)

				function := ginReturn.Fun.(*ast.SelectorExpr)

				if ginReturn.Fun.(*ast.SelectorExpr).Sel.Name != "JSON" {
					pass.Report(
						analysis.Diagnostic{
							Pos:     function.Pos(),
							Message: ReturnJSON,
						},
					)
				}

				_, ok := ginReturn.Args[1].(*ast.CallExpr)
				if !ok {
					pass.Report(
						analysis.Diagnostic{
							Pos:     function.Pos(),
							End:     function.End(),
							Message: ReturnObjectSerializer,
							SuggestedFixes: []analysis.SuggestedFix{
								{
									Message: "See: ",
									TextEdits: []analysis.TextEdit{{
										Pos:     function.Pos(),
										End:     function.End(),
										NewText: []byte("https://disaev.me/p/writing-useful-go-analysis-linter/"),
									}},
								},
							},
						},
					)
				}

				serializeFromModules := ginReturn.Args[1].(*ast.CompositeLit).Elts[0].(*ast.KeyValueExpr).Value.(*ast.CallExpr).Fun.(*ast.SelectorExpr).X.(*ast.SelectorExpr).X.(*ast.Ident).Obj.Decl.(*ast.Field).Names[0].Name

				if serializeFromModules != "serializer" {
					pass.Report(
						analysis.Diagnostic{
							Pos:     function.Pos(),
							End:     function.End(),
							Message: ReturnObjectSerializer,
							SuggestedFixes: []analysis.SuggestedFix{
								{
									Message: "See: ",
									TextEdits: []analysis.TextEdit{{
										Pos:     function.Pos(),
										End:     function.End(),
										NewText: []byte("https://disaev.me/p/writing-useful-go-analysis-linter/"),
									}},
								},
							},
						},
					)
				}
			}
		}
	}
	return nil, nil
}
