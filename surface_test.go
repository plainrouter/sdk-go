package plainrouter

import (
	"encoding/json"
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
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
					receiver := receiverName(declaration.Recv.List[0].Type)
					if !ast.IsExported(receiver) {
						continue
					}
					name = receiver + "." + name
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
	// Compare the full regenerated request-builder operation set with the signed
	// contract, rather than a stale hard-coded count.
	raw, err := os.ReadFile("spec/openapi.json")
	if err != nil {
		t.Fatal(err)
	}
	var spec struct {
		Paths map[string]map[string]struct {
			OperationID string `json:"operationId"`
		} `json:"paths"`
	}
	if err := json.Unmarshal(raw, &spec); err != nil {
		t.Fatal(err)
	}
	var operations []string
	separator := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	for _, item := range spec.Paths {
		for _, operation := range item {
			if operation.OperationID == "" {
				continue
			}
			parts := separator.Split(operation.OperationID, -1)
			for i, part := range parts {
				parts[i] = strings.ToUpper(part[:1]) + part[1:]
			}
			operations = append(operations, strings.Join(parts, ""))
		}
	}
	files, err := filepath.Glob("openapi/api_*.go")
	if err != nil {
		t.Fatal(err)
	}
	var builders []string
	for _, name := range files {
		file, err := parser.ParseFile(token.NewFileSet(), name, nil, 0)
		if err != nil {
			t.Fatal(err)
		}
		for _, declaration := range file.Decls {
			method, ok := declaration.(*ast.FuncDecl)
			if !ok || method.Recv == nil || method.Type.Results == nil || len(method.Type.Results.List) != 1 {
				continue
			}
			if !strings.HasSuffix(receiverName(method.Recv.List[0].Type), "Service") {
				continue
			}
			if strings.HasSuffix(receiverName(method.Type.Results.List[0].Type), "Request") {
				builders = append(builders, method.Name.Name)
			}
		}
	}
	sort.Strings(operations)
	sort.Strings(builders)
	if !reflect.DeepEqual(operations, builders) {
		t.Fatalf("operation set mismatch: contract=%v generated=%v", operations, builders)
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
