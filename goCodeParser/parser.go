package main

// import (
// 	"fmt"
// 	"go/ast"
// )

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	fset := token.NewFileSet() // positions are relative to fset

	src := `package foo

import (
	"fmt"
	"time"
)
var Count = 0
func bar() {
	fmt.Println(time.Now())
}

func Foo() {
	fmt.Println(time.Now())
}`

	// Parse src but stop after processing the imports.
	f, err := parser.ParseFile(fset, "", src, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println(f.Decls)
	funcs := []*ast.FuncDecl{}
	// Print the imports from the file's AST.
	for _, d := range f.Decls {
		if fn, isFn := d.(*ast.FuncDecl); isFn {
			funcs = append(funcs, fn)
		}
	}

	for index := 0; index < len(funcs); index++ {
		fmt.Println("FUNC NAME:", funcs[index].Name)
	}

}

// const subPackage = "../corelab.mkcl.org/MKCLOS/coredevelopmentplatform/corepkgv2/filemdl"

// func main() {
// 	set := token.NewFileSet()

// 	packs, err := parser.ParseDir(set, subPackage, nil, 0)
// 	if err != nil {
// 		fmt.Println("Failed to parse package:", err)
// 		os.Exit(1)
// 	}

// 	funcs := []*ast.FuncDecl{}
// 	for _, pack := range packs {
// 		for _, f := range pack.Files {
// 			for _, d := range f.Decls {
// 				// if fn, isFn := d.(); isFn {
// 				// funcs = append(funcs, fn)
// 				// fmt.Println("CALL EXP:", fn)
// 				// }
// 				fmt.Println("Declaration:", d)
// 				if fn, isFn := d.(*ast.FuncDecl); isFn {
// 					funcs = append(funcs, fn)
// 				}
// 			}
// 		}
// 	}

// 	for _, function := range funcs {
// 		extractFuncCallInFunc(function.Body.List)
// 	}
// 	// fmt.Printf("all funcs: %+v\n", funcs)

// 	for index := 0; index < len(funcs); index++ {
// 		// ba, _ := ffjson.Marshal(funcs[index])
// 		fmt.Println("FUNC NAME:", funcs[index].Name)
// 		// filemdl.WriteFile("./result.json", ba, false, false)
// 	}
// }

func extractFuncCallInFunc(stmts []ast.Stmt) {
	for _, stmt := range stmts {
		if exprStmt, ok := stmt.(*ast.ExprStmt); ok {
			if call, ok := exprStmt.X.(*ast.CallExpr); ok {
				if fun, ok := call.Fun.(*ast.SelectorExpr); ok {
					funcName := fun.Sel.Name
					fmt.Println("FUNCNCALL NAME:", funcName)
				}
			}
		}
	}
}
