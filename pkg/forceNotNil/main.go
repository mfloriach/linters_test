package force

import (
	"flag"
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
)

const (
	InterfacesMusTBeChecked = "must allways check your dependecy injected are not nil See: https://github.com/mfloriach/linters_test/tree/main/pkg/forceNotNil"
)

//nolint:gochecknoglobals
var flagSet flag.FlagSet

func NewAnalyzer() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name:  "not_nil",
		Doc:   "injected dependencies must be checked",
		Run:   run,
		Flags: flagSet,
	}
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		if file.Name.String() == "services" || file.Name.String() == "interfaces" {
			for _, declaration := range file.Decls {
				if function, ok := declaration.(*ast.FuncDecl); ok {
					if strings.HasPrefix(function.Name.Name, "New") {
						for _, paramObj := range function.Type.Params.List {
							if _, ok := paramObj.Type.(*ast.SelectorExpr); ok {
								paramType := paramObj.Type.(*ast.SelectorExpr).Sel.Name
								if strings.HasSuffix(paramType, "Interface") {
									paramName := paramObj.Names[0].Name

									switch function.Body.List[0].(type) {
									case *ast.IfStmt:
										condition := function.Body.List[0].(*ast.IfStmt).Cond.(*ast.BinaryExpr)

										if condition.X.(*ast.Ident).Obj.Name != paramName ||
											condition.Op.String() != "==" ||
											condition.Y.(*ast.Ident).Name != "nil" {
											pass.Report(
												analysis.Diagnostic{
													Pos:     function.Pos(),
													End:     function.End(),
													Message: InterfacesMusTBeChecked,
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
									default:
										pass.Report(
											analysis.Diagnostic{
												Pos:     function.Pos(),
												End:     function.End(),
												Message: InterfacesMusTBeChecked,
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
					}
				}
			}
		}
	}
	return nil, nil
}
