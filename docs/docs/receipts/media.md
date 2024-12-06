
# <Badge style='font-size: 1.8rem; text-shadow: 1px 1px 2px rgba(0, 0, 0, 0.3); padding: 0.35rem 0.75rem 0.35rem 0;' type='info' text='🔖 Media' />


```go
import (
	"fmt"
	"image/png"
	"log"
	"os"
	"time"
	"github.com/thuongtruong109/gouse"
)
```

## 1. Media canvas



```go
func MediaCanvas() {
	avatar, err := gouse.CreateCanvas(200, "#FFFFFF")
	if err != nil {
		log.Fatal(err)
	}
	filename := fmt.Sprintf("mockdata/%d.png", time.Now().Unix())
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	png.Encode(file, avatar)
}
```

## 2. Media png to jpg



```go
func MediaPngToJpg() {
	err := gouse.PNGToJPG("mockdata/1720031107.png", "mockdata/output.jpg")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Conversion successful")
}
```
