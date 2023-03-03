package imports

import (
	"flag"
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
)

const (
	NotAllow       = "forbidden to import this library in this module"
	NotAllowModule = "forbidden to import other internal modules to this module"
)

//nolint:gochecknoglobals
var flagSet flag.FlagSet

func NewImportsAnalyzer() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name:  "arch",
		Doc:   "check don't forget document endpoints",
		Run:   run,
		Flags: flagSet,
	}
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		checkNotInterModulesRelationship(file, pass)

		if file.Name.String() == "interfaces" {
			checkInterfacesImports(file, pass)
		}

		if file.Name.String() == "repositories" {
			checkRepositoryImports(file, pass)
		}

		if file.Name.String() == "services" {
			checkServiceImports(file, pass)
		}
	}

	return nil, nil
}

func checkServiceImports(file *ast.File, pass *analysis.Pass) {
	for _, im := range file.Imports {
		if im.Path.Value == "\"gorm.io/gorm\"" || im.Path.Value == "\"github.com/gin-gonic/gin\"" {
			pass.Report(
				analysis.Diagnostic{
					Pos:     im.Pos(),
					Message: NotAllow,
				},
			)
		}
	}
}

func checkRepositoryImports(file *ast.File, pass *analysis.Pass) {
	for _, im := range file.Imports {
		if im.Path.Value != "\"gorm.io/gorm\"" && im.Path.Value != "\"context\"" {
			pass.Report(
				analysis.Diagnostic{
					Pos:     im.Pos(),
					Message: NotAllow,
				},
			)
		}
	}
}

func checkInterfacesImports(file *ast.File, pass *analysis.Pass) {
	fileModule := checkModuleName(file, pass)

	for _, im := range file.Imports {
		if im.Path.Value != "\"github.com/gin-gonic/gin\"" && im.Path.Value != "\"net/http\"" {
			importPath := strings.Split(im.Path.Value, "/")
			if len(importPath) > 2 && importPath[0] == "\"hoge" && importPath[2] != fileModule {
				pass.Report(
					analysis.Diagnostic{
						Pos:     im.Pos(),
						Message: NotAllowModule,
					},
				)
			}
		}
	}
}

func checkModuleName(file *ast.File, pass *analysis.Pass) string {
	filePath := pass.Fset.Position(file.Package).Filename
	namespaces := strings.Split(filePath, "/")
	return namespaces[6]
}

func checkNotInterModulesRelationship(file *ast.File, pass *analysis.Pass) {
	filePath := pass.Fset.Position(file.Package).Filename

	if strings.Contains(filePath, "/backend_session/cmd/") {
		return
	}

	fileModule := checkModuleName(file, pass)

	for _, im := range file.Imports {
		importPath := strings.Split(im.Path.Value, "/")
		if len(importPath) > 2 && importPath[0] == "\"hoge" && importPath[2] != fileModule {
			pass.Report(
				analysis.Diagnostic{
					Pos:     im.Pos(),
					Message: NotAllowModule,
				},
			)
		}
	}
}
