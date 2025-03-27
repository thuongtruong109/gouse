package samples

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/thuongtruong109/gouse"
)

/*
Description: Create a path where file to be created.
Input params: (path string)
*/
func IoCreatePath() {
	relativePath := "tmp/example.txt"

	if err := gouse.CreatePath(relativePath); err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	println("File created successfully.")
}

/*
Description: Read the path where file to be read.
Input params: (path string)
*/
func IoReadPath() {
	relativePath := "tmp/example.txt"

	content, err := gouse.ReadPath(relativePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println("File content:", string(content))
}

/*
Description: Write content to the path where file to be updated.
Input params: (path string, content []byte)
*/
func IoWritePath() {
	relativePath := "tmp/example.txt"

	newContent := []byte("This is a new content")

	if err := gouse.WritePath(relativePath, newContent); err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	println("File updated successfully.")
}

/*
Description: Create folder.
Input params: (folderName string)
*/
func IoCreateDir() {
	err2 := gouse.CreateDir("tmp")
	if err2 != nil {
		println(err2.Error())
	}
	println("dir created")
}

/*
Description: Get current folder.
*/
func IoCurrentDir() {
	data, err := gouse.CurrentDir()
	if err != nil {
		println(err.Error())
		return
	}

	println(data)
}

/*
Description: Get the folder hierarchy.
Input params: (position string)
*/
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

/*
Description: Check if the folder exists.
Input params: (dirName string)
*/
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

/*
Description: List all files in the folder.
Input params: (position string)
*/
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

/*
Description: Remove folder.
Input params: (dirName string)
*/
func IoRemoveDir() {
	err3 := gouse.RemoveDir("tmp")
	if err3 != nil {
		println(err3.Error())
	}
	println("dir removed")
}

/*
Description: Append data to the file.
Input params: (fileName string, data []string)
*/
func IoAppendToFile() {
	err := gouse.AppendFile("data.json", []string{"this is data 3", "this is data 4"})
	if err != nil {
		println(err.Error())
	}
	println("file appended")
}

/*
Description: Clean the file.
Input params: (fileName string)
*/
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

/*
Description: Copy content from one file to another.
Input params: (sourceName string, destinationName string)
*/
func IoCopyFile() {
	err := gouse.CopyFile("data.json", "data2.json")
	if err != nil {
		println(err.Error())
	}
	println("file copied")
}

/*
Description: Create a file.
Input params: (fileName string)
*/
func IoCreateFile() {
	err := gouse.CreateFile("data.json")
	if err != nil {
		println(err.Error())
	}
	println("file created")
}

/*
Description: Get file info.
Input params: (fileName string)
*/
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

/*
Description: Check if the file exists.
Input params: (fileName string)
*/
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

/*
Description: Write and update object to the file.
Input params: (fileName string, data interface{})
*/
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

/*
Description: Read file by line.
Input params: (fileName string)
*/
func IoReadFileByLine() {
	data, err := gouse.ReadFileByLine("main.go")
	if err != nil {
		println(err.Error())
	}
	for _, v := range data {
		println(v)
	}
}

/*
Description: Remove file.
Input params: (fileName string)
*/
func IoRemoveFile() {
	err := gouse.RemoveFile("data.json")
	if err != nil {
		println(err.Error())
	}
	println("file removed")
}

/*
Description: Rename file.
Input params: (oldName string, newName string)
*/
func IoRenameFile() {
	err := gouse.RenameFile("data.json", "data2.json")
	if err != nil {
		println(err.Error())
	}
	println("file renamed")
}

/*
Description: Truncate data of file.
Input params: (fileName string, size int64)
*/
func IoTruncateFile() {
	err := gouse.TruncateFile("data.json", 10)
	if err != nil {
		println(err.Error())
	}
	println("file truncated to 10 bytes")
}

/*
Description: Write data to file.
Input params: (fileName string, data []string)
*/
func IoWriteToFile() {
	err := gouse.WriteFile("data.json", []string{"this is data 1", "this is data 2"})
	if err != nil {
		println(err.Error())
	}
	println("file written")
}

/*
Description: Zip files.
Input params: (zipFileName string, filesToZip []string)
*/
func IoZip() {
	filesToZip := []string{"file1.txt", "file2.txt"}
	zipFileName := "archive.zip"
	err := gouse.Zip(zipFileName, filesToZip)
	if err != nil {
		fmt.Println("Error zipping files:", err)
	}

	println("Files zipped successfully:", zipFileName)
}

/*
Description: Unzip files.
Input params: (zipFileName string, destFolder string)
*/
func IoUnzip() {
	destFolder := "unzipped"
	zipFileName := "archive.zip"
	err := gouse.Extract(zipFileName, destFolder)
	if err != nil {
		fmt.Println("Error unzipping files:", err)
	}

	println("Files unzipped successfully to:", destFolder)
}
