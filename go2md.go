package gouse

import (
	"go/ast"
	"go/parser"
	"go/token"
)

type Function struct {
	Import string
	Order  string
	Name   string
	Desc   string
	Body   string
}

func (f *Function) HighlightImport() string {
	return Sprintf("```go\nimport (\n%s)\n```\n", f.Import)
}

func (f *Function) HighlightName() string {
	return Sprintf("\n## %s. %s\n", f.Order, FirstUpperCase(SpaceCase(f.Name)))
}

func (f *Function) HighlightDesc() string {
	newDesc := Replace(f.Desc, "\n", "<br>")
	return Sprintf("\n%s\n", newDesc)
}

func (f *Function) HighlightBody() string {
	return Sprintf("\n```go\n%s```\n", f.Body)
}

func _extractImports(content []byte) string {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "string", string(content), parser.ParseComments)
	if err != nil {
		Println("Error parsing file:", err)
		return ""
	}

	var result []byte
	for _, decl := range node.Decls {
		if genDecl, ok := decl.(*ast.GenDecl); ok && genDecl.Tok == token.IMPORT {
			for _, spec := range genDecl.Specs {
				if importSpec, ok := spec.(*ast.ImportSpec); ok {
					result = append(result, Sprintf("\t%s", string(content[importSpec.Pos()-1:importSpec.End()]))...)
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
		Println("Error parsing file:", err)
		return nil
	}

	var functions []Function

	for _, decl := range file.Decls {
		if fn, ok := decl.(*ast.FuncDecl); ok {
			funcName := fn.Name.Name
			funcBody := string(code[fn.Pos()-1 : fn.End()])

			var funcComment string
			if fn.Doc != nil {
				funcComment = fn.Doc.Text()
			}

			functions = append(functions, Function{
				Import: _extractImports(code),
				Order:  Int2Str(len(functions) + 1),
				Name:   funcName,
				Desc:   funcComment,
				Body:   funcBody,
			})
		}
	}

	return functions
}

func _detectContent(content []byte) []byte {
	var result []string
	functions := ExtractFunctions(content)
	result = append(result, functions[0].HighlightImport())
	for _, function := range functions {
		tmpFunc := &Function{
			Import: function.Import,
			Order:  function.Order,
			Name:   function.Name,
			Desc:   function.Desc,
			Body:   function.Body,
		}
		result = append(result, tmpFunc.HighlightName())
		result = append(result, tmpFunc.HighlightDesc())
		result = append(result, tmpFunc.HighlightBody())
	}

	return Strs2Bytes(result)
}

func Go2Md(inputFilePath string, outputFilePath string) {
	content, err := ReadPath(inputFilePath)
	if err != nil {
		panic(err)
	}

	wrapper := _detectContent(content)

	err = WritePath(outputFilePath, wrapper)
	if err != nil {
		panic(err)
	}
}
