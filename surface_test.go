package plainrouter

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"reflect"
	"sort"
	"strings"
	"testing"
)

func TestRootPackageExportedSurfaceRemainsCurated(t *testing.T) {
	packages, err := parser.ParseDir(token.NewFileSet(), ".", func(info fs.FileInfo) bool {
		return !strings.HasSuffix(info.Name(), "_test.go")
	}, 0)
	if err != nil {
		t.Fatal(err)
	}

	root := packages["plainrouter"]
	if root == nil {
		t.Fatal("plainrouter package not found")
	}

	var exports []string
	for _, file := range root.Files {
		for _, declaration := range file.Decls {
			switch declaration := declaration.(type) {
			case *ast.FuncDecl:
				if !ast.IsExported(declaration.Name.Name) {
					continue
				}
				name := declaration.Name.Name
				if declaration.Recv != nil {
					name = receiverName(declaration.Recv.List[0].Type) + "." + name
				}
				exports = append(exports, name)
			case *ast.GenDecl:
				for _, specification := range declaration.Specs {
					switch specification := specification.(type) {
					case *ast.TypeSpec:
						if ast.IsExported(specification.Name.Name) {
							exports = append(exports, specification.Name.Name)
						}
					case *ast.ValueSpec:
						for _, name := range specification.Names {
							if ast.IsExported(name.Name) {
								exports = append(exports, name.Name)
							}
						}
					}
				}
			}
		}
	}

	sort.Strings(exports)
	want := []string{
		"Client",
		"Client.OpenAPI",
		"ContextAccessToken",
		"New",
		"NewAPIClient",
		"NewConfiguration",
		"Option",
		"WithBaseURL",
		"WithHTTPClient",
		"WithUserAgent",
	}
	if !reflect.DeepEqual(exports, want) {
		t.Fatalf("root package exports changed:\n got: %v\nwant: %v", exports, want)
	}
}

func receiverName(expression ast.Expr) string {
	switch expression := expression.(type) {
	case *ast.Ident:
		return expression.Name
	case *ast.StarExpr:
		return receiverName(expression.X)
	default:
		return "unknown"
	}
}
