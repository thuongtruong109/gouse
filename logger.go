package gouse

import (
	"fmt"
	"log"
	"os"
	"time"
)

func WriteLogDefault(logID, logMessage, output string) {
	file, err := os.OpenFile(output, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.Printf("Failed to open log file: %v", err)
	}

	if err := (os.Chmod(output, 0600)); err != nil {
		log.Printf("Error setting chmod log file handle permissions: %v", err)
	}

	defer func() {
		if err := file.Close(); err != nil {
			log.Printf("Failed to close log file: %v", err)
		}
	}()

	logger := log.New(file, logID, log.Ldate|log.Ltime|log.Lshortfile)

	logger.Println(logMessage)
}

type ILog struct {
	Prefix  string
	Message string
	Output  string
}

func WriteLog(prefix, msg, filePath string, err ...any) {
	logParam := &ILog{
		Prefix:  prefix,
		Message: "",
		Output:  filePath,
	}
	logId := RandStr(8)

	if len(err) > 0 {
		errStr := DetectErr(err)
		logParam.Message = fmt.Sprintf("Message: %s - ID: %s - Error: %s - Date: \n", msg, logId, errStr)
	} else {
		logParam.Message = fmt.Sprintf("Message: %s - ID: %s - Date: \n", msg, logId)
	}

	WriteLogDefault(logParam.Output, fmt.Sprintf("%s %s", logParam.Prefix, logParam.Message), logParam.Output)
}

func WriteErrorLog(msg string, err any) {
	WriteLog(ERROR_LOG_PREFIX, msg, ERROR_LOG_PATH, err)
}

func WriteInfoLog(msg string) {
	WriteLog(INFO_LOG_PREFIX, msg, INFO_LOG_PATH)
}

func WriteAccessLog(msg string) {
	WriteLog(ACCESS_LOG_PREFIX, msg, ACCESS_LOG_PATH)
}

func truncateLog(file *os.File) error {
	err := file.Truncate(0)
	if err != nil {
		return err
	}

	_, err = file.Seek(0, 0)
	if err != nil {
		return err
	}

	return nil
}

func RotateLog(output string, maxFileSize ...int64) {
	maxSize := int64(1 * 1024 * 1024) // 1MB

	if len(maxFileSize) > 0 {
		maxSize = int64(maxFileSize[0])
	}

	logFile, err := os.OpenFile(output, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Printf("Error opening log file: %v", err)
	}

	if err := (os.Chmod(output, 0600)); err != nil {
		log.Printf("Error setting chmod rotate log file permissions: %v", err)
	}

	defer func() {
		if err := logFile.Close(); err != nil {
			log.Print("Failed to close log file: ", err)
		}
	}()

	log.SetOutput(logFile)

	for {
		info, err := logFile.Stat()
		if err != nil {
			log.Printf("Error getting log file stats: %v", err)
		}

		if info.Size() >= maxSize {
			err = truncateLog(logFile)
			if err != nil {
				log.Printf("Error truncating log file: %v", err)
			}
		}

		time.Sleep(time.Minute)
	}
}
