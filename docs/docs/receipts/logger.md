
# <Badge style='font-size: 1.8rem; text-shadow: 1px 1px 2px rgba(0, 0, 0, 0.3); padding: 0.35rem 0.75rem 0.35rem 0;' type='info' text='🔖 Logger' />


```go
import (
	"github.com/thuongtruong109/gouse"
)
```

## 1. Logger write default

Description: Write log to file (as default)<br>Input params: (logId, logMessage, output string)<br>

```go
func LoggerWriteDefault() {
	gouse.WriteLogDefault("id_1", "log message", "output.log")
}
```

## 2. Logger write as gouse

Description: Write log to file (as Gouse style format)<br>Input params: (prefix, msg, filePath string, err ...interface{})<br>

```go
func LoggerWriteAsGouse() {
	gouse.WriteLog("[POST]", "message", "output.log", "error")
}
```

## 3. Logger write with type

Description: Write log by types to file (as Gouse style format)<br>Input params: (msg string, err interface{})<br>

```go
func LoggerWriteWithType() {
	gouse.WriteErrorLog("log message", "error")
	gouse.WriteInfoLog("log message")
	gouse.WriteAccessLog("log message")
}
```

## 4. Logger auto rotate truncate

Description: Truncate log file<br>Input params: (output string, maxFileSize ...int64)<br>

```go
func LoggerAutoRotateTruncate() {
	gouse.RotateLog("output.log", 1024)
}
```
