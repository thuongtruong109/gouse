package gouse

import (
	"fmt"
	"log"
	"os"
	"time"
)

type ILog struct {
	Prefix  string
	Message string
	Output  string
}

func DetectError(err interface{}) string {
	switch e := err.(type) {
	case error:
		return e.Error()
	case string:
		return e
	default:
		return fmt.Sprintf("%v", err)
	}
}

func WriteLog(level, msg, filePath string, err ...interface{}) {
	logParam := &ILog{
		Prefix:  level,
		Message: "",
		Output:  filePath,
	}
	logId := RandStr(8)

	if len(err) > 0 {
		errStr := DetectError(err)
		logParam.Message = fmt.Sprintf("Message: %s - ID: %s - Error: %s - Date: \n", msg, logId, errStr)
	} else {
		logParam.Message = fmt.Sprintf("Message: %s - ID: %s - Date: \n", msg, logId)
	}

	HandleLog(logParam.Output, fmt.Sprintf("%s %s", logParam.Prefix, logParam.Message), logParam.Output)
}

func WriteErrorLog(msg string, err interface{}) {
	WriteLog(ERROR_LOG_PREFIX, msg, ERROR_LOG_PATH, err)
}

func WriteInfoLog(msg string) {
	WriteLog(INFO_LOG_PREFIX, msg, INFO_LOG_PATH)
}

func WriteAccessLog(msg string) {
	WriteLog(ACCESS_LOG_PREFIX, msg, ACCESS_LOG_PATH)
}

func WriteCustomLog(prefix, msg string, filePath ...string) {
	WriteLog(prefix, msg, filePath[0])
}

func HandleLog(logID, logMessage, output string) {
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

	// Create a new logger that writes to the log file
	logger := log.New(file, logID, log.Ldate|log.Ltime|log.Lshortfile)

	// Write some log messages
	logger.Println(logMessage)
}

func RotateTruncateLog(output string, maxFileSize ...int64) {
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
			err = TruncateLog(logFile)
			if err != nil {
				log.Printf("Error truncating log file: %v", err)
			}
		}

		time.Sleep(time.Minute)
	}
}

func TruncateLog(file *os.File) error {
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
