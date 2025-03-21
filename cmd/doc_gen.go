package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/thuongtruong109/gouse"
)

func clearOutputPath(outputPath, newName string) {
	err := os.RemoveAll(filepath.Join(outputPath, newName))
	if err != nil {
		fmt.Println("Error removing output path:", err)
		return
	}
}

func processPath(path, outputPath, newName string) {
	_, err := gouse.IsExistDir(path)
	if err != nil {
		println("Please provide a path to a directory")
		return
	}

	stat, err := os.Stat(path)
	if err != nil {
		fmt.Println("Error getting file stat:", err)
		return
	}

	if !stat.IsDir() {
		fmt.Println("Please provide a path to a directory")
		return
	}

	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	processFiles(files, path, outputPath, newName)
}

func processFiles(files []os.DirEntry, path, outputPath, newName string) {
	for _, file := range files {
		if !gouse.EndsWith(file.Name(), ".go") || gouse.EndsWith(file.Name(), "_test.go") {
			continue
		}

		processGoFile(file, path, outputPath, newName)
	}
}

func processGoFile(file os.DirEntry, path, outputPath, newName string) {
	content, err := os.ReadFile(filepath.Join(path, file.Name()))
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	functions := gouse.ExtractFunctions(content)
	generateMarkdownFiles(file, functions, path, outputPath, newName)
}

/* Working use Vitepress style */

func highlightName(text string) string {
	return fmt.Sprintf("\n# <Badge style='font-size: 1.8rem; text-shadow: 1px 1px 2px rgba(0, 0, 0, 0.3); padding: 0.35rem 0.75rem 0.35rem 0;' type='info' text='🔖 %s' />\n", text)
}

func generateMarkdownFiles(file os.DirEntry, functions []gouse.Function, path, outputPath, newName string) {
	var result []byte

	result = append(result, []byte(highlightName(gouse.Capitalize(gouse.TrimSuf(file.Name(), ".go")))+"\n\n")...)

	result = append(result, functions[0].HighlightImport()...)

	for _, function := range functions {
		result = append(result, function.HighlightName()...)
		result = append(result, function.HighlightDesc()...)
		result = append(result, function.HighlightBody()...)
	}

	fileName := gouse.Replace(file.Name(), ".go", ".md")

	// replace (./oldName/*.go -> ./newName/*.go)
	detectOld := gouse.Split(path, "/")[1]
	renamedPath := gouse.Replace(path, detectOld, newName)

	subPath := filepath.Join(outputPath, renamedPath, filepath.Dir(file.Name()))

	createFilePath(subPath, fileName, result)
}

func createFilePath(subPath, fileName string, result []byte) {
	err := gouse.CreatePath(filepath.Join(subPath, fileName))
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}

	err = os.WriteFile(filepath.Join(subPath, fileName), result, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
}

func GenerateDocument(outputPath, newName string) {
	if len(os.Args) < 2 {
		println("Please provide at least one path to a file or directory")
		return
	}

	clearOutputPath(outputPath, newName)

	for i := 1; i < len(os.Args); i++ {
		processPath(os.Args[i], outputPath, newName)
	}
}

func main() {
	GenerateDocument("docs", "docs/receipts")
}
