package metrics

import (
	"encoding/json"
	"math"
	"os"
	"strings"

	"github.com/jedib0t/go-pretty/table"
)

var metrics = `{"files":[{"code":1282,"comment":0,"blank":336,"name":"spm-go/output.html","language":"HTML"},{"code":458,"comment":0,"blank":9,"name":".arch-go/report.css","language":"CSS"},{"code":228,"comment":0,"blank":44,"name":".arch-go/report.html","language":"HTML"},{"code":98,"comment":1,"blank":19,"name":"pkg/architecture/main.go","language":"Go"},{"code":84,"comment":1,"blank":13,"name":"pkg/ginReturnSerializer/main.go","language":"Go"},{"code":83,"comment":1,"blank":8,"name":"pkg/forceNotNil/main.go","language":"Go"},{"code":69,"comment":15,"blank":21,"name":"pkg/traceavility/main.go","language":"Go"},{"code":54,"comment":41,"blank":9,"name":".golangci.yml","language":"YAML"},{"code":46,"comment":2,"blank":7,"name":"pkg/swaggo/main.go","language":"Go"},{"code":37,"comment":0,"blank":5,"name":".air.toml","language":"TOML"},{"code":37,"comment":5,"blank":6,"name":"pkg/modularization/main.go","language":"Go"},{"code":30,"comment":13,"blank":10,"name":"modules/comments/interfaces/gin.go","language":"Go"},{"code":30,"comment":0,"blank":10,"name":"pkg/modularization/README.md","language":"Markdown"},{"code":28,"comment":0,"blank":10,"name":"cmd/main.go","language":"Go"},{"code":27,"comment":0,"blank":6,"name":"pkg/swaggo/README.md","language":"Markdown"},{"code":24,"comment":12,"blank":8,"name":"modules/posts/interfaces/gin.go","language":"Go"},{"code":21,"comment":0,"blank":3,"name":"arch-go.yml","language":"YAML"},{"code":20,"comment":2,"blank":5,"name":"modules/comments/main.go","language":"Go"},{"code":20,"comment":0,"blank":3,"name":"modules/comments/serializers/comment_serializer.go","language":"Go"},{"code":20,"comment":0,"blank":6,"name":"modules/comments/services/comment_service.go","language":"Go"},{"code":20,"comment":0,"blank":6,"name":"modules/posts/services/posts_service.go","language":"Go"},{"code":20,"comment":4,"blank":6,"name":"plugins/arch.go","language":"Go"},{"code":19,"comment":2,"blank":6,"name":"modules/posts/main.go","language":"Go"},{"code":19,"comment":0,"blank":10,"name":"pkg/forceNotNil/README.md","language":"Markdown"},{"code":17,"comment":0,"blank":7,"name":"pkg/ginReturnSerializer/README.md","language":"Markdown"},{"code":16,"comment":0,"blank":1,"name":".vscode/settings.json","language":"JSON"},{"code":16,"comment":0,"blank":8,"name":"pkg/traceavility/README.md","language":"Markdown"},{"code":15,"comment":0,"blank":5,"name":"modules/comments/repositories/comments_mysql.go","language":"Go"},{"code":14,"comment":0,"blank":5,"name":"modules/posts/repositories/post_mysql.go","language":"Go"},{"code":13,"comment":0,"blank":6,"name":"makefile","language":"Makefile"},{"code":12,"comment":0,"blank":3,"name":"pkg/gorm_sqlite.go","language":"Go"},{"code":9,"comment":42,"blank":0,"name":".reviewdog.yml","language":"YAML"},{"code":9,"comment":0,"blank":1,"name":"docker-compose.yml","language":"YAML"},{"code":8,"comment":0,"blank":3,"name":"README.md","language":"Markdown"},{"code":8,"comment":0,"blank":3,"name":"cmd/linter/main.go","language":"Go"},{"code":5,"comment":0,"blank":2,"name":"modules/comments/repositories/comments.go","language":"Go"},{"code":5,"comment":0,"blank":2,"name":"modules/posts/repositories/post.go","language":"Go"}],"total":{"files":37,"code":2921,"comment":141,"blank":612}}`

type Metrics struct {
	Files []File
}

type File struct {
	Code     int
	Name     string
	Language string
}

type Module struct {
	Files      int
	Code       int
	Percentage float64
}

func Code() {
	var result Metrics
	json.Unmarshal([]byte(metrics), &result)

	numFiles := 0
	numLines := 0

	var modules = make(map[string]*Module)

	for _, file := range result.Files {
		if file.Language != "Go" {
			continue
		}

		if strings.Contains(file.Name, "modules/") {
			modulesName := strings.Split(file.Name, "/")[1]
			if _, ok := modules[modulesName]; !ok {
				modules[modulesName] = &Module{
					Files: 0,
					Code:  0,
				}
			}

			modules[modulesName] = &Module{
				Files: modules[modulesName].Files + 1,
				Code:  modules[modulesName].Code + file.Code,
			}

			numFiles += 1
			numLines += file.Code
		}
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Modules", "Code (%)"})

	for key, m := range modules {
		t.AppendRows([]table.Row{
			{key, math.Ceil(float64(m.Code)/float64(numLines)*10000) / 100},
		})
	}

	t.Render()

}
