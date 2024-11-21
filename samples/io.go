package samples

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/thuongtruong109/gouse"
)

/* Samples for io path functions */
func SampleIoCreatePath() {
	relativePath := "tmp/example.txt"

	if err := gouse.CreatePath(relativePath); err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	println("File created successfully.")
}

func SampleIoReadPath() {
	relativePath := "tmp/example.txt"

	content, err := gouse.ReadPath(relativePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println("File content:", string(content))
}

func SampleIoWritePath() {
	relativePath := "tmp/example.txt"

	newContent := []byte("This is a new content")

	if err := gouse.WritePath(relativePath, newContent); err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	println("File updated successfully.")
}

/* Directory functions */

func SampleIoCreateDir() {
	err2 := gouse.CreateDir("tmp")
	if err2 != nil {
		println(err2.Error())
	}
	println("dir created")
}

func SampleIoCurrentDir() {
	data, err := gouse.CurrentDir()
	if err != nil {
		println(err.Error())
		return
	}

	println(data)
}

func SampleIoHierarchyDir() {
	data, err := gouse.HierarchyDir(".")
	if err != nil {
		println(err.Error())
		return
	}

	for _, v := range data {
		println(v)
	}
}

func SampleIoCheckDir() {
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

func SampleIoLsDir() {
	data, err := gouse.LsDir(".")
	if err != nil {
		println(err.Error())
		return
	}

	for _, v := range data {
		println(v)
	}
}

func SampleIoRemoveDir() {
	err3 := gouse.RemoveDir("tmp")
	if err3 != nil {
		println(err3.Error())
	}
	println("dir removed")
}

/* File functions */

func SampleIoAppendToFile() {
	err := gouse.AppendFile("data.json", []string{"this is data 3", "this is data 4"})
	if err != nil {
		println(err.Error())
	}
	println("file appended")
}

func SampleIoCleanFile() {
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

func SampleIoCopyFile() {
	err := gouse.CopyFile("data.json", "data2.json")
	if err != nil {
		println(err.Error())
	}
	println("file copied")
}

func SampleIoCreateFile() {
	err := gouse.CreateFile("data.json")
	if err != nil {
		println(err.Error())
	}
	println("file created")
}

func SampleIoFileInfo() {
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

func SampleIoCheckFile() {
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

func SampleIoFileObj() {
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

func SampleIoReadFileByLine() {
	data, err := gouse.ReadFileByLine("main.go")
	if err != nil {
		println(err.Error())
	}
	for _, v := range data {
		println(v)
	}
}

func SampleIoRemoveFile() {
	err := gouse.RemoveFile("data.json")
	if err != nil {
		println(err.Error())
	}
	println("file removed")
}

func SampleIoRenameFile() {
	err := gouse.RenameFile("data.json", "data2.json")
	if err != nil {
		println(err.Error())
	}
	println("file renamed")
}

func SampleIoTruncateFile() {
	err := gouse.TruncateFile("data.json", 10)
	if err != nil {
		println(err.Error())
	}
	println("file truncated to 10 bytes")
}

func SampleIoWriteToFile() {
	err := gouse.WriteFile("data.json", []string{"this is data 1", "this is data 2"})
	if err != nil {
		println(err.Error())
	}
	println("file written")
}

/* Utility functions */

func SampleIoZip() {
	filesToZip := []string{"file1.txt", "file2.txt"}
	zipFileName := "archive.zip"
	err := gouse.Zip(zipFileName, filesToZip)
	if err != nil {
		fmt.Println("Error zipping files:", err)
	}

	println("Files zipped successfully:", zipFileName)
}

func SampleIoUnzip() {
	destFolder := "unzipped"
	zipFileName := "archive.zip"
	err := gouse.Unzip(zipFileName, destFolder)
	if err != nil {
		fmt.Println("Error unzipping files:", err)
	}

	println("Files unzipped successfully to:", destFolder)
}
