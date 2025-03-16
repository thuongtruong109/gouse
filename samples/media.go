package samples

import (
	"image/png"
	"log"
	"os"
	"time"

	"github.com/thuongtruong109/gouse"
)

/*
Description: Create a canvas with a white background and save it as a PNG file
Input params: (size int, background string)
*/
func Canvas() {
	avatar, err := gouse.CreateCanvas(200, "#FFFFFF")
	if err != nil {
		log.Fatal(err)
	}
	filename := gouse.Sprintf("mockdata/%d.png", time.Now().Unix())
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	png.Encode(file, avatar)
}

/*
Description: Convert a PNG image to a JPG image
Input params: (jpgPath string, pngPath string)
*/
func PngToJpg() {
	err := gouse.Png2Jpg("mockdata/1720031107.png", "mockdata/output.jpg")
	if err != nil {
		gouse.Println(err)
		return
	}

	gouse.Println("Conversion successful")
}
