package samples

import "github.com/thuongtruong109/gouse"

/*
Description: Write log to file (as default)
Input params: (logId, logMessage, output string)
*/
func WriteLogDefault() {
	gouse.WriteLogDefault("id_1", "log message", "output.log")
}

/*
Description: Write log to file (as Gouse style format)
Input params: (prefix, msg, filePath string, err ...interface{})
*/
func WriteLogAsGouse() {
	gouse.WriteLog("[POST]", "message", "output.log", "error")
}

/*
Description: Write log by types to file (as Gouse style format)
Input params: (msg string, err interface{})
*/
func WriteLogWithType() {
	gouse.WriteErrorLog("log message", "error")
	gouse.WriteInfoLog("log message")
	gouse.WriteAccessLog("log message")
}

/*
Description: Truncate log file
Input params: (output string, maxFileSize ...int64)
*/
func AutoRotateTruncateLog() {
	gouse.RotateLog("output.log", 1024)
}