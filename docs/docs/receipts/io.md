
# <Badge style='font-size: 1.8rem; text-shadow: 1px 1px 2px rgba(0, 0, 0, 0.3); padding: 0.35rem 0.75rem 0.35rem 0;' type='info' text='🔖 Io' />


```go
import (
	"encoding/json"
	"os"
	"github.com/thuongtruong109/gouse"
)
```

## 1. Create path

Description: Create a path where file to be created.<br>Input params: (path string)<br>

```go
func CreatePath() {
	relativePath := "tmp/example.txt"

	if err := gouse.CreatePath(relativePath); err != nil {
		gouse.Println("Error creating file:", err)
		return
	}
	println("File created successfully.")
}
```

## 2. Read path

Description: Read the path where file to be read.<br>Input params: (path string)<br>

```go
func ReadPath() {
	relativePath := "tmp/example.txt"

	content, err := gouse.ReadPath(relativePath)
	if err != nil {
		gouse.Println("Error reading file:", err)
		return
	}
	gouse.Println("File content:", string(content))
}
```

## 3. Write path

Description: Write content to the path where file to be updated.<br>Input params: (path string, content []byte)<br>

```go
func WritePath() {
	relativePath := "tmp/example.txt"

	newContent := []byte("This is a new content")

	if err := gouse.WritePath(relativePath, newContent); err != nil {
		gouse.Println("Error writing to file:", err)
		return
	}
	println("File updated successfully.")
}
```

## 4. Create dir

Description: Create folder.<br>Input params: (folderName string)<br>

```go
func CreateDir() {
	err2 := gouse.CreateDir("tmp")
	if err2 != nil {
		println(err2.Error())
	}
	println("dir created")
}
```

## 5. Current dir

Description: Get current folder.<br>

```go
func CurrentDir() {
	data, err := gouse.CurrentDir()
	if err != nil {
		println(err.Error())
		return
	}

	println(data)
}
```

## 6. Hierarchy dir

Description: Get the folder hierarchy.<br>Input params: (position string)<br>

```go
func HierarchyDir() {
	data, err := gouse.HierarchyDir(".")
	if err != nil {
		println(err.Error())
		return
	}

	for _, v := range data {
		println(v)
	}
}
```

## 7. Check dir

Description: Check if the folder exists.<br>Input params: (dirName string)<br>

```go
func CheckDir() {
	isExist, err1 := gouse.IsExistDir("tmp")
	if err1 != nil {
		println(err1.Error())
	}
	if isExist {
		println("dir exist")
	} else {
		println("dir not exist")
	}
}
```

## 8. Ls dir

Description: List all files in the folder.<br>Input params: (position string)<br>

```go
func LsDir() {
	data, err := gouse.LsDir(".")
	if err != nil {
		println(err.Error())
		return
	}

	for _, v := range data {
		println(v)
	}
}
```

## 9. Remove dir

Description: Remove folder.<br>Input params: (dirName string)<br>

```go
func RemoveDir() {
	err3 := gouse.RemoveDir("tmp")
	if err3 != nil {
		println(err3.Error())
	}
	println("dir removed")
}
```

## 10. Append to file

Description: Append data to the file.<br>Input params: (fileName string, data []string)<br>

```go
func AppendToFile() {
	err := gouse.AppendFile("data.json", []string{"this is data 3", "this is data 4"})
	if err != nil {
		println(err.Error())
	}
	println("file appended")
}
```

## 11. Clean file

Description: Clean the file.<br>Input params: (fileName string)<br>

```go
func CleanFile() {
	err := gouse.CleanFile("data.json")
	if err != nil {
		println(err.Error())
	}

	// or using truncate with 0 bytes
	// err := helper.TruncateFile("data.json", 0)
	// if err != nil {
	// 	println(err.Error())
	// }

	println("file cleaned")
}
```

## 12. Copy file

Description: Copy content from one file to another.<br>Input params: (sourceName string, destinationName string)<br>

```go
func CopyFile() {
	err := gouse.CopyFile("data.json", "data2.json")
	if err != nil {
		println(err.Error())
	}
	println("file copied")
}
```

## 13. Create file

Description: Create a file.<br>Input params: (fileName string)<br>

```go
func CreateFile() {
	err := gouse.CreateFile("data.json")
	if err != nil {
		println(err.Error())
	}
	println("file created")
}
```

## 14. File info

Description: Get file info.<br>Input params: (fileName string)<br>

```go
func FileInfo() {
	data, err := gouse.FileInfo("main.go")
	if err != nil {
		println(err.Error())
	}
	gouse.Printf("File info: %+v\n", data.All)
	gouse.Println("File info (with name):", data.Name)
	gouse.Printf("File info (with size): %d bytes\n", data.Size)
	gouse.Println("File info (with permissions):", data.Mode)
	gouse.Println("File info (with last modified):", data.ModTime)
	gouse.Println("File info (with directory check): ", data.IsDir)
	gouse.Printf("File info (with system process): %+v\n", data.Sys)
}
```

## 15. Check file

Description: Check if the file exists.<br>Input params: (fileName string)<br>

```go
func CheckFile() {
	isExist, err := gouse.IsExistFile("data.json")
	if err != nil {
		println(err.Error())
	}
	if isExist {
		println("file exist")
	} else {
		println("file not exist")
	}
}
```

## 16. File obj

Description: Write and update object to the file.<br>Input params: (fileName string, data interface{})<br>

```go
func FileObj() {
	type User struct {
		Name string
		Age  int
	}

	exampleFile := "data.json"

	// read file
	// data, err := .ReadFileObj[User](exampleFile)
	// if err != nil {
	// 	println(err.Error())
	// }
	// gouse.Printf("data: %+v\n", data)

	// // write file
	// updateData := append(data, User{
	// 	Name: gouse.Sprintf("name %d", len(data)+1),
	// 	Age:  len(data) + 1,
	// })

	data, err := os.ReadFile(exampleFile)
	if err != nil {
		println(err.Error())
		return
	}
	gouse.Printf("data: %+v\n", data)

	// unmarshal data into a slice of User
	var users []User
	if err := json.Unmarshal(data, &users); err != nil {
		println(err.Error())
		return
	}

	// create a new user
	newUser := User{
		Name: gouse.Sprintf("name %d", len(users)+1),
		Age:  len(users) + 1,
	}

	// append the new user to the slice
	users = append(users, newUser)

	// marshal the updated slice back to JSON
	updatedData, err := json.Marshal(users)
	if err != nil {
		println(err.Error())
		return
	}

	// write the updated data back to the file
	// if err := util.WriteFile(exampleFile, updatedData, 0644); err != nil {
	// 	println(err.Error())
	// 	return
	// }

	err2 := gouse.WriteFileObj(exampleFile, updatedData)
	if err2 != nil {
		println(err2.Error())
	}
	println("data written")
}
```

## 17. Read file by line

Description: Read file by line.<br>Input params: (fileName string)<br>

```go
func ReadFileByLine() {
	data, err := gouse.ReadFileByLine("main.go")
	if err != nil {
		println(err.Error())
	}
	for _, v := range data {
		println(v)
	}
}
```

## 18. Remove file

Description: Remove file.<br>Input params: (fileName string)<br>

```go
func RemoveFile() {
	err := gouse.RemoveFile("data.json")
	if err != nil {
		println(err.Error())
	}
	println("file removed")
}
```

## 19. Rename file

Description: Rename file.<br>Input params: (oldName string, newName string)<br>

```go
func RenameFile() {
	err := gouse.RenameFile("data.json", "data2.json")
	if err != nil {
		println(err.Error())
	}
	println("file renamed")
}
```

## 20. Truncate file

Description: Truncate data of file.<br>Input params: (fileName string, size int64)<br>

```go
func TruncateFile() {
	err := gouse.TruncateFile("data.json", 10)
	if err != nil {
		println(err.Error())
	}
	println("file truncated to 10 bytes")
}
```

## 21. Write to file

Description: Write data to file.<br>Input params: (fileName string, data []string)<br>

```go
func WriteToFile() {
	err := gouse.WriteFile("data.json", []string{"this is data 1", "this is data 2"})
	if err != nil {
		println(err.Error())
	}
	println("file written")
}
```

## 22. Zip

Description: Zip files.<br>Input params: (zipFileName string, filesToZip []string)<br>

```go
func Zip() {
	filesToZip := []string{"file1.txt", "file2.txt"}
	zipFileName := "archive.zip"
	err := gouse.Zip(zipFileName, filesToZip)
	if err != nil {
		gouse.Println("Error zipping files:", err)
	}

	println("Files zipped successfully:", zipFileName)
}
```

## 23. Unzip

Description: Unzip files.<br>Input params: (zipFileName string, destFolder string)<br>

```go
func Unzip() {
	destFolder := "unzipped"
	zipFileName := "archive.zip"
	err := gouse.Unzip(zipFileName, destFolder)
	if err != nil {
		gouse.Println("Error unzipping files:", err)
	}

	println("Files unzipped successfully to:", destFolder)
}
```
