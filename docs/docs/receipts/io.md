
# <Badge style='font-size: 1.8rem; text-shadow: 1px 1px 2px rgba(0, 0, 0, 0.3); padding: 0.35rem 0.75rem 0.35rem 0;' type='info' text='🔖 Io' />


```go
import (
	"encoding/json"
	"fmt"
	"os"
	"github.com/thuongtruong109/gouse"
)
```

## 1. Io create path

Description: Create a path where file to be created.<br>Input params: (path string)<br>

```go
func IoCreatePath() {
	relativePath := "tmp/example.txt"

	if err := gouse.CreatePath(relativePath); err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	println("File created successfully.")
}
```

## 2. Io read path

Description: Read the path where file to be read.<br>Input params: (path string)<br>

```go
func IoReadPath() {
	relativePath := "tmp/example.txt"

	content, err := gouse.ReadPath(relativePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println("File content:", string(content))
}
```

## 3. Io write path

Description: Write content to the path where file to be updated.<br>Input params: (path string, content []byte)<br>

```go
func IoWritePath() {
	relativePath := "tmp/example.txt"

	newContent := []byte("This is a new content")

	if err := gouse.WritePath(relativePath, newContent); err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	println("File updated successfully.")
}
```

## 4. Io create dir

Description: Create folder.<br>Input params: (folderName string)<br>

```go
func IoCreateDir() {
	err2 := gouse.CreateDir("tmp")
	if err2 != nil {
		println(err2.Error())
	}
	println("dir created")
}
```

## 5. Io current dir

Description: Get current folder.<br>

```go
func IoCurrentDir() {
	data, err := gouse.CurrentDir()
	if err != nil {
		println(err.Error())
		return
	}

	println(data)
}
```

## 6. Io hierarchy dir

Description: Get the folder hierarchy.<br>Input params: (position string)<br>

```go
func IoHierarchyDir() {
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

## 7. Io check dir

Description: Check if the folder exists.<br>Input params: (dirName string)<br>

```go
func IoCheckDir() {
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

## 8. Io ls dir

Description: List all files in the folder.<br>Input params: (position string)<br>

```go
func IoLsDir() {
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

## 9. Io remove dir

Description: Remove folder.<br>Input params: (dirName string)<br>

```go
func IoRemoveDir() {
	err3 := gouse.RemoveDir("tmp")
	if err3 != nil {
		println(err3.Error())
	}
	println("dir removed")
}
```

## 10. Io append to file

Description: Append data to the file.<br>Input params: (fileName string, data []string)<br>

```go
func IoAppendToFile() {
	err := gouse.AppendFile("data.json", []string{"this is data 3", "this is data 4"})
	if err != nil {
		println(err.Error())
	}
	println("file appended")
}
```

## 11. Io clean file

Description: Clean the file.<br>Input params: (fileName string)<br>

```go
func IoCleanFile() {
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

## 12. Io copy file

Description: Copy content from one file to another.<br>Input params: (sourceName string, destinationName string)<br>

```go
func IoCopyFile() {
	err := gouse.CopyFile("data.json", "data2.json")
	if err != nil {
		println(err.Error())
	}
	println("file copied")
}
```

## 13. Io create file

Description: Create a file.<br>Input params: (fileName string)<br>

```go
func IoCreateFile() {
	err := gouse.CreateFile("data.json")
	if err != nil {
		println(err.Error())
	}
	println("file created")
}
```

## 14. Io file info

Description: Get file info.<br>Input params: (fileName string)<br>

```go
func IoFileInfo() {
	data, err := gouse.FileInfo("main.go")
	if err != nil {
		println(err.Error())
	}
	fmt.Printf("File info: %+v\n", data.All)
	fmt.Println("File info (with name):", data.Name)
	fmt.Printf("File info (with size): %d bytes\n", data.Size)
	fmt.Println("File info (with permissions):", data.Mode)
	fmt.Println("File info (with last modified):", data.ModTime)
	fmt.Println("File info (with directory check): ", data.IsDir)
	fmt.Printf("File info (with system process): %+v\n", data.Sys)
}
```

## 15. Io check file

Description: Check if the file exists.<br>Input params: (fileName string)<br>

```go
func IoCheckFile() {
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

## 16. Io file obj

Description: Write and update object to the file.<br>Input params: (fileName string, data interface{})<br>

```go
func IoFileObj() {
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
	// fmt.Printf("data: %+v\n", data)

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
	fmt.Printf("data: %+v\n", data)

	// unmarshal data into a slice of User
	var users []User
	if err := json.Unmarshal(data, &users); err != nil {
		println(err.Error())
		return
	}

	// create a new user
	newUser := User{
		Name: fmt.Sprintf("name %d", len(users)+1),
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

## 17. Io read file by line

Description: Read file by line.<br>Input params: (fileName string)<br>

```go
func IoReadFileByLine() {
	data, err := gouse.ReadFileByLine("main.go")
	if err != nil {
		println(err.Error())
	}
	for _, v := range data {
		println(v)
	}
}
```

## 18. Io remove file

Description: Remove file.<br>Input params: (fileName string)<br>

```go
func IoRemoveFile() {
	err := gouse.RemoveFile("data.json")
	if err != nil {
		println(err.Error())
	}
	println("file removed")
}
```

## 19. Io rename file

Description: Rename file.<br>Input params: (oldName string, newName string)<br>

```go
func IoRenameFile() {
	err := gouse.RenameFile("data.json", "data2.json")
	if err != nil {
		println(err.Error())
	}
	println("file renamed")
}
```

## 20. Io truncate file

Description: Truncate data of file.<br>Input params: (fileName string, size int64)<br>

```go
func IoTruncateFile() {
	err := gouse.TruncateFile("data.json", 10)
	if err != nil {
		println(err.Error())
	}
	println("file truncated to 10 bytes")
}
```

## 21. Io write to file

Description: Write data to file.<br>Input params: (fileName string, data []string)<br>

```go
func IoWriteToFile() {
	err := gouse.WriteFile("data.json", []string{"this is data 1", "this is data 2"})
	if err != nil {
		println(err.Error())
	}
	println("file written")
}
```

## 22. Io zip

Description: Zip files.<br>Input params: (zipFileName string, filesToZip []string)<br>

```go
func IoZip() {
	filesToZip := []string{"file1.txt", "file2.txt"}
	zipFileName := "archive.zip"
	err := gouse.Zip(zipFileName, filesToZip)
	if err != nil {
		fmt.Println("Error zipping files:", err)
	}

	println("Files zipped successfully:", zipFileName)
}
```

## 23. Io extract

Description: Unzip files.<br>Input params: (zipFileName string, destFolder string)<br>

```go
func IoExtract() {
	destFolder := "unzipped"
	zipFileName := "archive.zip"
	err := gouse.Extract(zipFileName, destFolder)
	if err != nil {
		fmt.Println("Error unzipping files:", err)
	}

	println("Files unzipped successfully to:", destFolder)
}
```
