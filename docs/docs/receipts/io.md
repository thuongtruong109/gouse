
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

 Samples for io path functions<br>

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



```go
func IoFileObj() {
	type User struct {
		Name string
		Age  int
	}

	exampleFile := "data.json"

	// read file
	// data, err := io.ReadFileObj[User](exampleFile)
	// if err != nil {
	// 	println(err.Error())
	// }
	// fmt.Printf("data: %+v\n", data)

	// // write file
	// updateData := append(data, User{
	// 	Name: fmt.Sprintf("name %d", len(data)+1),
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
	// if err := ioutil.WriteFile(exampleFile, updatedData, 0644); err != nil {
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

## 23. Io unzip



```go
func IoUnzip() {
	destFolder := "unzipped"
	zipFileName := "archive.zip"
	err := gouse.Unzip(zipFileName, destFolder)
	if err != nil {
		fmt.Println("Error unzipping files:", err)
	}

	println("Files unzipped successfully to:", destFolder)
}
```
