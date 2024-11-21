package gouse

import (
	"archive/zip"
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

/* File */

func IsExistFile(path string) (bool, error) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		_, err := os.Create(path)
		if err != nil {
			return false, err
		}
	}
	return true, nil
}

func CreateFile(path string) error {
	_, err := os.Create(path)
	if err != nil {
		return err
	}
	return nil
}

func RemoveFile(path string) error {
	_, err := IsExistFile(path)
	if err != nil {
		return err
	}

	err = os.Remove(path)
	if err != nil {
		return err
	}

	return nil
}

func WriteFile(path string, data []string) error {
	os.WriteFile(path, []byte(Join(data, " ")), 0600)
	return nil
}

func AppendFile(path string, data []string) error {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return nil
	}

	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	for _, v := range data {
		if _, err := file.WriteString(v + "\n"); err != nil {
			return err
		}
	}

	return nil
}

func ReadFileByLine(path string) ([]string, error) {
	_, err := IsExistFile(path)
	if err != nil {
		return nil, err
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}

type FileInfoStruct struct {
	Name    string
	Size    int64
	Mode    os.FileMode
	ModTime string
	IsDir   bool
	Sys     interface{}
	All     os.FileInfo
}

func FileInfo(path string) (*FileInfoStruct, error) {
	_, err := IsExistFile(path)
	if err != nil {
		return nil, err
	}

	fileInfo, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	return &FileInfoStruct{
		Name:    fileInfo.Name(),
		Size:    fileInfo.Size(),
		Mode:    fileInfo.Mode(),
		ModTime: ISODate(fileInfo.ModTime()),
		IsDir:   fileInfo.IsDir(),
		Sys:     fileInfo.Sys(),
		All:     fileInfo,
	}, nil
}

func RenameFile(oldPath, newPath string) error {
	_, err := IsExistFile(oldPath)
	if err != nil {
		return err
	}

	err = os.Rename(oldPath, newPath)
	if err != nil {
		return err
	}

	return nil
}

func CopyFile(oldPath, newPath string) error {
	_, err := IsExistFile(oldPath)
	if err != nil {
		return err
	}

	source, err := os.Open(oldPath)
	if err != nil {
		return err
	}

	defer func() {
		if err := source.Close(); err != nil {
			panic(err)
		}
	}()

	target, targetError := os.OpenFile(newPath, os.O_RDWR|os.O_CREATE, 0666)
	if targetError != nil {
		return targetError
	}

	defer func() {
		if err := target.Close(); err != nil {
			panic(err)
		}
	}()

	_, copyError := io.Copy(target, source)
	if copyError != nil {
		return copyError
	}

	return nil
}

func TruncateFile(path string, size int64) error {
	_, err := IsExistFile(path)
	if err != nil {
		return err
	}

	err = os.Truncate(path, size)
	if err != nil {
		return err
	}

	return nil
}

func CleanFile(path string) error {
	TruncateFile(path, 0)
	return nil
}

func ReadFileObj[T any](path string) ([]T, error) {
	_, err := IsExistFile(path)
	if err != nil {
		return nil, err
	}

	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	if len(content) == 0 {
		return nil, nil
	}

	var data []T

	err = json.Unmarshal(content, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func WriteFileObj[T any](path string, data T) error {
	_, err := IsExistFile(path)
	if err != nil {
		err2 := CreateDir(path)
		if err2 != nil {
			return err2
		}
	}

	content, err := json.Marshal(data)
	if err != nil {
		return nil
	}

	err = os.WriteFile(path, content, 0644)
	if err != nil {
		return nil
	}
	return nil
}

/* Directory */

func IsExistDir(dir string) (bool, error) {
	_, err := os.Stat(dir)
	if err == nil {
		return false, err
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return true, nil
}

func CreateDir(dir string) error {
	isExist, err := IsExistDir(dir)
	if err != nil {
		return err
	}
	if !isExist {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}
	}

	return nil
}

func RemoveDir(dir string) error {
	_, err := IsExistDir(dir)
	if err != nil {
		return err
	}

	err = os.RemoveAll(dir)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func LsDir(dir string) ([]string, error) {
	_, err := IsExistDir(dir)
	if err != nil {
		return nil, err
	}

	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var result []string
	for _, file := range files {
		result = append(result, file.Name())
	}

	return result, nil
}

func HierarchyDir(dir string) ([]string, error) {
	_, err := IsExistDir(dir)
	if err != nil {
		return nil, err
	}

	var result []string
	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		result = append(result, path)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return result, nil
}

func CurrentDir() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return dir, nil
}

/* Path */

func CreatePath(relativePath string) error {
	absolutePath, err := filepath.Abs(relativePath)
	if err != nil {
		return err
	}

	dir := filepath.Dir(absolutePath)

	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}

	file, err := os.Create(absolutePath)
	if err != nil {
		return err
	}
	defer file.Close()

	return nil
}

func ReadPath(relativePath string) ([]byte, error) {
	absolutePath, err := filepath.Abs(relativePath)
	if err != nil {
		return nil, err
	}

	content, err := os.ReadFile(absolutePath)
	if err != nil {
		return nil, err
	}

	return content, nil
}

func WritePath(relativePath string, content []byte) error {
	absolutePath, err := filepath.Abs(relativePath)
	if err != nil {
		return err
	}

	err = os.WriteFile(absolutePath, content, 0644)
	if err != nil {
		return err
	}

	return nil
}

/* Compress */

func Zip(zipFileName string, files []string) error {
	zipFile, err := os.Create(zipFileName)
	if err != nil {
		return err
	}
	defer func() {
		if closeErr := zipFile.Close(); closeErr != nil {
			fmt.Println("Error closing zip file:", closeErr)
		}
	}()

	zipWriter := zip.NewWriter(zipFile)
	defer func() {
		if closeErr := zipWriter.Close(); closeErr != nil {
			fmt.Println("Error closing zip writer:", closeErr)
		}
	}()

	for _, file := range files {
		fileToZip, err := os.Open(file)
		if err != nil {
			return err
		}

		zipEntryWriter, err := zipWriter.Create(file)
		if err != nil {
			fileToZip.Close()
			return err
		}

		_, err = io.Copy(zipEntryWriter, fileToZip)
		fileToZip.Close()
		if err != nil {
			return err
		}
	}

	return nil
}

func Unzip(zipFile, destFolder string) error {
	const maxFileSize = 100 * 1024 * 1024  // 100 MB
	const maxTotalSize = 500 * 1024 * 1024 // 500 MB

	zipReader, err := zip.OpenReader(zipFile)
	if err != nil {
		return err
	}
	defer func() {
		if closeErr := zipReader.Close(); closeErr != nil {
			fmt.Println("Error closing zip reader:", closeErr)
		}
	}()

	var totalSize int64

	for _, file := range zipReader.File {
		filePath := filepath.Join(destFolder, file.Name)

		if file.FileInfo().IsDir() {
			os.MkdirAll(filePath, os.ModePerm)
			continue
		}

		if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
			return err
		}

		fileToExtract, err := file.Open()
		if err != nil {
			return err
		}
		defer fileToExtract.Close()

		targetFile, err := os.Create(filePath)
		if err != nil {
			return err
		}
		defer targetFile.Close()

		// Check the size of the file before copying
		fileSize, err := io.CopyN(targetFile, fileToExtract, maxFileSize)
		if err != nil && err != io.EOF {
			return err
		}

		totalSize += fileSize

		if totalSize > maxTotalSize {
			return errors.New("exceeded maximum total size limit for extraction")
		}
	}

	return nil
}
