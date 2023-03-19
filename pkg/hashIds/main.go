package hashids

import (
	"flag"
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
)

const (
	hashID = "must IDs must be string See: https://github.com/mfloriach/linters_test/tree/main/pkg/hashIds"
)

//nolint:gochecknoglobals
var flagSet flag.FlagSet

func NewAnalyzer() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name:  "hash_ids",
		Doc:   "struct ID must be string",
		Run:   run,
		Flags: flagSet,
	}
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		if file.Name.String() == "serializers" {

			for _, o := range file.Scope.Objects {
				if o.Kind.String() == "type" {
					for _, p := range o.Decl.(*ast.TypeSpec).Type.(*ast.StructType).Fields.List {
						if strings.Contains(p.Names[0].Name, "ID") && p.Type.(*ast.Ident).Name != "string" {
							pass.Report(
								analysis.Diagnostic{
									Pos:     p.Pos(),
									Message: hashID,
								},
							)
						}
					}
				}
			}

		}
	}
	return nil, nil
}
