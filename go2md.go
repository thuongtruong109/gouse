package gouse

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

type Function struct {
	Import string
	Name   string
	Body   string
}

func (f *Function) HighlightImport() string {
	return fmt.Sprintf("```go\nimport (\n%s)\n```\n", f.Import)
}

func (f *Function) HighlightName(orderNum string) string {
	return fmt.Sprintf("\n#### %s\n", orderNum+". "+f.Name)
}

func (f *Function) HighlightBody() string {
	return fmt.Sprintf("\n```go\n%s```\n", f.Body)
}

func extractImports(content []byte) string {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "string", string(content), parser.ParseComments)
	if err != nil {
		fmt.Println("Error parsing file:", err)
		return ""
	}

	var result []byte
	for _, decl := range node.Decls {
		if genDecl, ok := decl.(*ast.GenDecl); ok && genDecl.Tok == token.IMPORT {
			for _, spec := range genDecl.Specs {
				if importSpec, ok := spec.(*ast.ImportSpec); ok {
					result = append(result, fmt.Sprintf("\t%s", string(content[importSpec.Pos()-1:importSpec.End()]))...)
				}
			}
		}
	}

	return string(result)
}

func ExtractFunctions(code []byte) []Function {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "string", string(code), parser.ParseComments)
	if err != nil {
		fmt.Println("Error parsing file:", err)
		return nil
	}

	var functions []Function

	for _, decl := range file.Decls {
		if fn, ok := decl.(*ast.FuncDecl); ok {
			funcName := fn.Name.Name
			funcBody := string(code[fn.Pos()-1 : fn.End()])

			functions = append(functions, Function{
				Import: extractImports(code),
				Name:   funcName,
				Body:   funcBody,
			})
		}
	}

	return functions
}

func detectContent(content []byte) []byte {
	var result []string
	functions := ExtractFunctions(content)
	result = append(result, functions[0].HighlightImport())
	for i, function := range functions {
		tmpFunc := &Function{
			Import: function.Import,
			Name:   function.Name,
			Body:   function.Body,
		}
		result = append(result, tmpFunc.HighlightName(IntToString(i+1)))
		result = append(result, tmpFunc.HighlightBody())
	}

	return StringsToBytes(result)
}

func Go2Md(inputFilePath string, outputFilePath string) {
	content, err := ReadPath(inputFilePath)
	if err != nil {
		panic(err)
	}

	wrapper := detectContent(content)

	err = WritePath(outputFilePath, wrapper)
	if err != nil {
		panic(err)
	}
}
