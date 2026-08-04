package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	fileSet := token.NewFileSet()
	node, err := parser.ParseFile(fileSet, "another_file.go", nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	fmt.Println("List of imported packages")
	for _, c := range node.Imports {
		fmt.Println(c.Path, c.Name)
	}
	// &{25 STRING "fmt"} <nil>
	// &{52 STRING "unsafe"} customNameForUnsafe

	fmt.Println("\nExported functions list")
	for _, decl := range node.Decls {
		f, ok := decl.(*ast.FuncDecl)
		if !ok {
			continue
		}
		if !f.Name.IsExported() {
			continue
		}

		fmt.Println(" - " + f.Name.Name)
	}

	fmt.Println("\nUnexported functions list")
	for _, decl := range node.Decls {
		f, ok := decl.(*ast.FuncDecl)
		if !ok {
			continue
		}
		if f.Name.IsExported() {
			continue
		}

		fmt.Println(" - " + f.Name.Name)
	}

	fmt.Println("\nAll return expressions")
	ast.Inspect(node, func(n ast.Node) bool {
		ret, ok := n.(*ast.ReturnStmt)
		if ok {
			fmt.Printf(" - return statement found on line %d:\n", fileSet.Position(ret.Pos()).Line)
		}
		return true
	})
	// return statement found on line 151:
}
