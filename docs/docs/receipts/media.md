# Media

## Imports

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
## Functions


### SampleMediaCanvas

```go
func SampleMediaCanvas() {
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
## Imports

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
## Functions


### SampleMediaPngToJpg

```go
func SampleMediaPngToJpg() {
	err := gouse.PNGToJPG("mockdata/1720031107.png", "mockdata/output.jpg")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Conversion successful")
}
```
