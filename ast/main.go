package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	var path = "./ast/ack.pb.go"
	result := GetStructNameByFile(path)
	fmt.Println(result)
}

func GetStructNameByFile(path string) []string {
	fset := token.NewFileSet()
	// 这里的参数基本是原封不动的传给了scanner的Init函数
	node, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}
	var attr *ast.GenDecl
	var ts *ast.TypeSpec
	var ok bool
	var nameMap = make(map[string]struct{})
	for _, decl := range node.Decls {
		attr, ok = decl.(*ast.GenDecl)
		if !ok {
			continue
		}
		for _, spec := range attr.Specs {
			ts, ok = spec.(*ast.TypeSpec)
			if !ok {
				continue
			}
			if _, ok = ts.Type.(*ast.StructType); !ok {
				continue
			}
			nameMap[ts.Name.Name] = struct{}{}
		}
	}
	var result = make([]string, len(nameMap))
	var index int
	for k, _ := range nameMap {
		result[index] = k
		index++
	}
	return result
}
