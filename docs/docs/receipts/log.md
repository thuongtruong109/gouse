
# <Badge style='font-size: 1.8rem; text-shadow: 1px 1px 2px rgba(0, 0, 0, 0.3); padding: 0.35rem 0.75rem 0.35rem 0;' type='info' text='🔖 Log' />


```go
import (
	"github.com/thuongtruong109/gouse"
)
```

## 1. Write log default

Description: Write log to file (as default)<br>Input params: (logId, logMessage, output string)<br>

```go
func WriteLogDefault() {
	gouse.WriteLogDefault("id_1", "log message", "output.log")
}
```

## 2. Write log as gouse

Description: Write log to file (as Gouse style format)<br>Input params: (prefix, msg, filePath string, err ...interface{})<br>

```go
func WriteLogAsGouse() {
	gouse.WriteLog("[POST]", "message", "output.log", "error")
}
```

## 3. Write log with type

Description: Write log by types to file (as Gouse style format)<br>Input params: (msg string, err interface{})<br>

```go
func WriteLogWithType() {
	gouse.WriteErrorLog("log message", "error")
	gouse.WriteInfoLog("log message")
	gouse.WriteAccessLog("log message")
}
```

## 4. Auto rotate truncate log

Description: Truncate log file<br>Input params: (output string, maxFileSize ...int64)<br>

```go
func AutoRotateTruncateLog() {
	gouse.RotateLog("output.log", 1024)
}
```
