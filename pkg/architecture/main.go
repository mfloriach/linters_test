package architecture

import (
	"flag"
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
)

const (
	RepositoryPattern = "should follow the repository pattern See: https://github.com/mfloriach/linters_test/tree/development/pkg/architecture"
	NotAllow          = "forbidden to import this library in this module See: https://github.com/mfloriach/linters_test/tree/development/pkg/architecture"
	NotAllowModule    = "forbidden to import other internal modules to this module See: https://github.com/mfloriach/linters_test/tree/development/pkg/architecture"
)

//nolint:gochecknoglobals
var flagSet flag.FlagSet

func NewAnalyzer() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name:  "architecture",
		Doc:   "should follows architecture complaince",
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
		if im.Path.Value == "\"gorm.io/gorm\"" {
			pass.Report(
				analysis.Diagnostic{
					Pos:     im.Pos(),
					Message: RepositoryPattern,
				},
			)
		}

		if im.Path.Value == "\"github.com/gin-gonic/gin\"" {
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
		if im.Path.Value != "\"gorm.io/gorm\"" && im.Path.Value != "\"context\"" && !strings.Contains(im.Path.Value, "/shared") {
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
		if im.Path.Value == "\"gorm.io/gorm\"" {
			pass.Report(
				analysis.Diagnostic{
					Pos:     im.Pos(),
					Message: RepositoryPattern,
				},
			)
		}

		if im.Path.Value != "\"github.com/gin-gonic/gin\"" && im.Path.Value != "\"net/http\"" && !strings.Contains(im.Path.Value, "/shared") {
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

	if strings.Contains(filePath, "/usr/src/app") {
		return namespaces[5]
	}

	return namespaces[6]
}

func checkNotInterModulesRelationship(file *ast.File, pass *analysis.Pass) {
	filePath := pass.Fset.Position(file.Package).Filename

	if strings.Contains(filePath, "/cmd/") {
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
