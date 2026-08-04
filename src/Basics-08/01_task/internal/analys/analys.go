package analys

import (
	"go/ast"
	"go/parser"
	"go/token"
)

type AnalysResult struct {
	DeclCount    int
	CallCount    int
	AssignCount  int
	ImportsCount int
}

func Analys(filepath string) (*AnalysResult, error) {
	fileSet := token.NewFileSet()
	node, err := parser.ParseFile(fileSet, filepath, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	result := &AnalysResult{}
	result.ImportsCount = len(node.Imports)

	ast.Inspect(node, func(n ast.Node) bool {
		switch n := n.(type) {
		case *ast.GenDecl:
			// Declarations: var, const
			if n.Tok == token.VAR || n.Tok == token.CONST {
				result.DeclCount++
			}
		case *ast.AssignStmt:
			// =, :=
			result.AssignCount++
		case *ast.CallExpr:
			// f(x)
			result.CallCount++
		}

		return true
	})

	return result, nil
}
