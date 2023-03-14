package metrics

import (
	"encoding/json"
	"fmt"
	"math"
	"strings"
)

type Spm struct {
	Packages []Package
}

type Package struct {
	Name             string
	Path             string
	AfferentCoupling int `json:"afferent_coupling"`
	EfferentCoupling int `json:"efferent_coupling"`
}

func Coupling() {
	var result Spm
	json.Unmarshal([]byte(couplingP), &result)

	average := 0
	totalPkgs := 0

	for _, pkg := range result.Packages {
		if strings.Contains(pkg.Path, "modules/") {
			totalPkgs += 1
			average += pkg.EfferentCoupling
		}
	}

	fmt.Println("Average Component Dependency: ", math.Ceil(float64(average)/float64(totalPkgs)*100)/100)
}

var couplingP = `
{
    "packages": [
        {
            "name": "main",
            "path": "hoge/cmd",
            "files": [
                "main.go"
            ],
            "files_count": 1,
            "dependencies": {
                "internals": [
                    "hoge/modules/comments",
                    "hoge/modules/posts",
                    "hoge/pkg"
                ],
                "externals": [
                    "github.com/gin-gonic/gin",
                    "gorm.io/gorm"
                ],
                "internals_count": 3,
                "externals_count": 2,
                "count": 5
            },
            "afferent_coupling": 0,
            "efferent_coupling": 3,
            "instability": 1,
            "abstractness_details": {
                "functions": 3
            },
            "abstractions_count": 0,
            "implementations_count": 3,
            "abstractness": 0,
            "distance": 0
        },
        {
            "name": "main",
            "path": "hoge/cmd/linter",
            "files": [
                "main.go"
            ],
            "files_count": 1,
            "dependencies": {
                "internals": [
                    "hoge/pkg/modularization"
                ],
                "externals": [
                    "golang.org/x/tools/go/analysis/singlechecker"
                ],
                "internals_count": 1,
                "externals_count": 1,
                "count": 2
            },
            "afferent_coupling": 0,
            "efferent_coupling": 1,
            "instability": 1,
            "abstractness_details": {
                "functions": 1
            },
            "abstractions_count": 0,
            "implementations_count": 1,
            "abstractness": 0,
            "distance": 0
        },
        {
            "name": "main",
            "path": "hoge/cmd/metrics",
            "files": [
                "main.go"
            ],
            "files_count": 1,
            "dependencies": {
                "standard": [
                    "encoding/json",
                    "fmt",
                    "math",
                    "strings"
                ],
                "standard_count": 4,
                "count": 4
            },
            "afferent_coupling": 0,
            "efferent_coupling": 0,
            "instability": 1,
            "abstractness_details": {
                "functions": 2,
                "structs": 3
            },
            "abstractions_count": 3,
            "implementations_count": 2,
            "abstractness": 0.6,
            "distance": 0.6
        },
        {
            "name": "comments",
            "path": "hoge/modules/comments",
            "files": [
                "main.go"
            ],
            "files_count": 1,
            "dependencies": {
                "internals": [
                    "hoge/modules/comments/interfaces",
                    "hoge/modules/comments/repositories",
                    "hoge/modules/comments/services"
                ],
                "externals": [
                    "github.com/gin-gonic/gin",
                    "gorm.io/gorm"
                ],
                "internals_count": 3,
                "externals_count": 2,
                "count": 5
            },
            "dependants": [
                "hoge/modules/comments"
            ],
            "afferent_coupling": 1,
            "efferent_coupling": 3,
            "instability": 0.75,
            "abstractness_details": {
                "functions": 3
            },
            "abstractions_count": 0,
            "implementations_count": 3,
            "abstractness": 0,
            "distance": 0.25
        },
        {
            "name": "interfaces",
            "path": "hoge/modules/comments/interfaces",
            "files": [
                "gin.go"
            ],
            "files_count": 1,
            "dependencies": {
                "standard": [
                    "errors",
                    "net/http"
                ],
                "internals": [
                    "hoge/modules/comments/serializers",
                    "hoge/modules/comments/services"
                ],
                "externals": [
                    "github.com/gin-gonic/gin"
                ],
                "standard_count": 2,
                "internals_count": 2,
                "externals_count": 1,
                "count": 5
            },
            "dependants": [
                "hoge/modules/comments",
                "hoge/modules/comments/interfaces"
            ],
            "afferent_coupling": 2,
            "efferent_coupling": 2,
            "instability": 0.5,
            "abstractness_details": {
                "methods": 1,
                "functions": 1,
                "structs": 1
            },
            "abstractions_count": 1,
            "implementations_count": 2,
            "abstractness": 0.33,
            "distance": 0.17
        },
        {
            "name": "repositories",
            "path": "hoge/modules/comments/repositories",
            "files": [
                "comments.go",
                "comments_mysql.go"
            ],
            "files_count": 2,
            "dependencies": {
                "standard": [
                    "context"
                ],
                "externals": [
                    "gorm.io/gorm"
                ],
                "standard_count": 1,
                "externals_count": 1,
                "count": 2
            },
            "dependants": [
                "hoge/modules/comments",
                "hoge/modules/comments/interfaces",
                "hoge/modules/comments/services",
                "hoge/modules/comments/repositories"
            ],
            "afferent_coupling": 4,
            "efferent_coupling": 0,
            "instability": 0,
            "abstractness_details": {
                "methods": 1,
                "functions": 1,
                "interfaces": 1,
                "structs": 1
            },
            "abstractions_count": 2,
            "implementations_count": 2,
            "abstractness": 0.5,
            "distance": 0.5
        },
        {
            "name": "serializers",
            "path": "hoge/modules/comments/serializers",
            "files": [
                "comment_serializer.go"
            ],
            "files_count": 1,
            "dependencies": {
                "standard": [
                    "time"
                ],
                "standard_count": 1,
                "count": 1
            },
            "dependants": [
                "hoge/modules/comments",
                "hoge/modules/comments/interfaces",
                "hoge/modules/comments/serializers"
            ],
            "afferent_coupling": 3,
            "efferent_coupling": 0,
            "instability": 0,
            "abstractness_details": {
                "functions": 1,
                "structs": 1
            },
            "abstractions_count": 1,
            "implementations_count": 1,
            "abstractness": 0.5,
            "distance": 0.5
        },
        {
            "name": "services",
            "path": "hoge/modules/comments/services",
            "files": [
                "comment_service.go"
            ],
            "files_count": 1,
            "dependencies": {
                "standard": [
                    "context"
                ],
                "internals": [
                    "hoge/modules/comments/repositories"
                ],
                "standard_count": 1,
                "internals_count": 1,
                "count": 2
            },
            "dependants": [
                "hoge/modules/comments",
                "hoge/modules/comments/interfaces",
                "hoge/modules/comments/services"
            ],
            "afferent_coupling": 3,
            "efferent_coupling": 1,
            "instability": 0.25,
            "abstractness_details": {
                "methods": 1,
                "functions": 1,
                "interfaces": 1,
                "structs": 1
            },
            "abstractions_count": 2,
            "implementations_count": 2,
            "abstractness": 0.5,
            "distance": 0.25
        },
        {
            "name": "posts",
            "path": "hoge/modules/posts",
            "files": [
                "main.go"
            ],
            "files_count": 1,
            "dependencies": {
                "internals": [
                    "hoge/modules/posts/interfaces",
                    "hoge/modules/posts/repositories",
                    "hoge/modules/posts/services"
                ],
                "externals": [
                    "github.com/gin-gonic/gin",
                    "gorm.io/gorm"
                ],
                "internals_count": 3,
                "externals_count": 2,
                "count": 5
            },
            "dependants": [
                "hoge/modules/posts"
            ],
            "afferent_coupling": 1,
            "efferent_coupling": 3,
            "instability": 0.75,
            "abstractness_details": {
                "functions": 3
            },
            "abstractions_count": 0,
            "implementations_count": 3,
            "abstractness": 0,
            "distance": 0.25
        },
        {
            "name": "interfaces",
            "path": "hoge/modules/posts/interfaces",
            "files": [
                "gin.go"
            ],
            "files_count": 1,
            "dependencies": {
                "standard": [
                    "net/http"
                ],
                "internals": [
                    "hoge/modules/posts/services"
                ],
                "externals": [
                    "github.com/gin-gonic/gin"
                ],
                "standard_count": 1,
                "internals_count": 1,
                "externals_count": 1,
                "count": 3
            },
            "dependants": [
                "hoge/modules/posts",
                "hoge/modules/posts/interfaces"
            ],
            "afferent_coupling": 2,
            "efferent_coupling": 1,
            "instability": 0.33,
            "abstractness_details": {
                "methods": 1,
                "functions": 1,
                "structs": 1
            },
            "abstractions_count": 1,
            "implementations_count": 2,
            "abstractness": 0.33,
            "distance": 0.34
        },
        {
            "name": "repositories",
            "path": "hoge/modules/posts/repositories",
            "files": [
                "post.go",
                "post_mysql.go"
            ],
            "files_count": 2,
            "dependencies": {
                "standard": [
                    "context"
                ],
                "externals": [
                    "gorm.io/gorm"
                ],
                "standard_count": 1,
                "externals_count": 1,
                "count": 2
            },
            "dependants": [
                "hoge/modules/posts",
                "hoge/modules/posts/interfaces",
                "hoge/modules/posts/services",
                "hoge/modules/posts/repositories"
            ],
            "afferent_coupling": 4,
            "efferent_coupling": 0,
            "instability": 0,
            "abstractness_details": {
                "methods": 1,
                "functions": 1,
                "interfaces": 1,
                "structs": 1
            },
            "abstractions_count": 2,
            "implementations_count": 2,
            "abstractness": 0.5,
            "distance": 0.5
        },
        {
            "name": "services",
            "path": "hoge/modules/posts/services",
            "files": [
                "posts_service.go"
            ],
            "files_count": 1,
            "dependencies": {
                "standard": [
                    "context"
                ],
                "internals": [
                    "hoge/modules/posts/repositories"
                ],
                "standard_count": 1,
                "internals_count": 1,
                "count": 2
            },
            "dependants": [
                "hoge/modules/posts",
                "hoge/modules/posts/interfaces",
                "hoge/modules/posts/services"
            ],
            "afferent_coupling": 3,
            "efferent_coupling": 1,
            "instability": 0.25,
            "abstractness_details": {
                "methods": 1,
                "functions": 1,
                "interfaces": 1,
                "structs": 1
            },
            "abstractions_count": 2,
            "implementations_count": 2,
            "abstractness": 0.5,
            "distance": 0.25
        },
        {
            "name": "pkg",
            "path": "hoge/pkg",
            "files": [
                "gorm_sqlite.go"
            ],
            "files_count": 1,
            "dependencies": {
                "externals": [
                    "gorm.io/driver/sqlite",
                    "gorm.io/gorm"
                ],
                "externals_count": 2,
                "count": 2
            },
            "dependants": [
                "hoge/pkg"
            ],
            "afferent_coupling": 1,
            "efferent_coupling": 0,
            "instability": 0,
            "abstractness_details": {
                "functions": 1
            },
            "abstractions_count": 0,
            "implementations_count": 1,
            "abstractness": 0,
            "distance": 1
        },
        {
            "name": "architecture",
            "path": "hoge/pkg/architecture",
            "files": [
                "main.go"
            ],
            "files_count": 1,
            "dependencies": {
                "standard": [
                    "flag",
                    "go/ast",
                    "strings"
                ],
                "externals": [
                    "golang.org/x/tools/go/analysis"
                ],
                "standard_count": 3,
                "externals_count": 1,
                "count": 4
            },
            "afferent_coupling": 0,
            "efferent_coupling": 0,
            "instability": 1,
            "abstractness_details": {
                "functions": 7,
                "interfaces": 1
            },
            "abstractions_count": 1,
            "implementations_count": 7,
            "abstractness": 0.13,
            "distance": 0.13
        },
        {
            "name": "force",
            "path": "hoge/pkg/forceNotNil",
            "files": [
                "main.go"
            ],
            "files_count": 1,
            "dependencies": {
                "standard": [
                    "flag",
                    "go/ast",
                    "strings"
                ],
                "externals": [
                    "golang.org/x/tools/go/analysis"
                ],
                "standard_count": 3,
                "externals_count": 1,
                "count": 4
            },
            "dependants": [
                "hoge/pkg/forceNotNil"
            ],
            "afferent_coupling": 1,
            "efferent_coupling": 0,
            "instability": 0,
            "abstractness_details": {
                "functions": 2,
                "interfaces": 1
            },
            "abstractions_count": 1,
            "implementations_count": 2,
            "abstractness": 0.33,
            "distance": 0.67
        },
        {
            "name": "gin",
            "path": "hoge/pkg/ginReturnSerializer",
            "files": [
                "main.go"
            ],
            "files_count": 1,
            "dependencies": {
                "standard": [
                    "flag",
                    "go/ast"
                ],
                "externals": [
                    "golang.org/x/tools/go/analysis"
                ],
                "standard_count": 2,
                "externals_count": 1,
                "count": 3
            },
            "dependants": [
                "hoge/pkg/ginReturnSerializer"
            ],
            "afferent_coupling": 1,
            "efferent_coupling": 0,
            "instability": 0,
            "abstractness_details": {
                "functions": 2,
                "interfaces": 1
            },
            "abstractions_count": 1,
            "implementations_count": 2,
            "abstractness": 0.33,
            "distance": 0.67
        },
        {
            "name": "modularization",
            "path": "hoge/pkg/modularization",
            "files": [
                "main.go"
            ],
            "files_count": 1,
            "dependencies": {
                "standard": [
                    "flag",
                    "strings"
                ],
                "externals": [
                    "golang.org/x/tools/go/analysis"
                ],
                "standard_count": 2,
                "externals_count": 1,
                "count": 3
            },
            "dependants": [
                "hoge/pkg/modularization"
            ],
            "afferent_coupling": 1,
            "efferent_coupling": 0,
            "instability": 0,
            "abstractness_details": {
                "functions": 2,
                "interfaces": 1
            },
            "abstractions_count": 1,
            "implementations_count": 2,
            "abstractness": 0.33,
            "distance": 0.67
        },
        {
            "name": "swaggo",
            "path": "hoge/pkg/swaggo",
            "files": [
                "main.go"
            ],
            "files_count": 1,
            "dependencies": {
                "standard": [
                    "flag",
                    "go/ast",
                    "strings"
                ],
                "externals": [
                    "golang.org/x/tools/go/analysis"
                ],
                "standard_count": 3,
                "externals_count": 1,
                "count": 4
            },
            "dependants": [
                "hoge/pkg/swaggo"
            ],
            "afferent_coupling": 1,
            "efferent_coupling": 0,
            "instability": 0,
            "abstractness_details": {
                "functions": 3,
                "interfaces": 1
            },
            "abstractions_count": 1,
            "implementations_count": 3,
            "abstractness": 0.25,
            "distance": 0.75
        },
        {
            "name": "traceavility",
            "path": "hoge/pkg/traceavility",
            "files": [
                "main.go"
            ],
            "files_count": 1,
            "dependencies": {
                "standard": [
                    "flag",
                    "go/ast",
                    "strings"
                ],
                "externals": [
                    "golang.org/x/tools/go/analysis"
                ],
                "standard_count": 3,
                "externals_count": 1,
                "count": 4
            },
            "dependants": [
                "hoge/pkg/traceavility"
            ],
            "afferent_coupling": 1,
            "efferent_coupling": 0,
            "instability": 0,
            "abstractness_details": {
                "functions": 6,
                "interfaces": 1
            },
            "abstractions_count": 1,
            "implementations_count": 6,
            "abstractness": 0.14,
            "distance": 0.86
        },
        {
            "name": "main",
            "path": "hoge/plugins",
            "files": [
                "arch.go"
            ],
            "files_count": 1,
            "dependencies": {
                "internals": [
                    "hoge/pkg/forceNotNil",
                    "hoge/pkg/ginReturnSerializer",
                    "hoge/pkg/modularization",
                    "hoge/pkg/swaggo",
                    "hoge/pkg/traceavility"
                ],
                "externals": [
                    "golang.org/x/tools/go/analysis"
                ],
                "internals_count": 5,
                "externals_count": 1,
                "count": 6
            },
            "afferent_coupling": 0,
            "efferent_coupling": 5,
            "instability": 1,
            "abstractness_details": {
                "methods": 1,
                "structs": 1
            },
            "abstractions_count": 1,
            "implementations_count": 1,
            "abstractness": 0.5,
            "distance": 0.5
        }
    ]
}`
