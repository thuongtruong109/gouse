package gouse

import (
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"log"
	"os"
)

func CreateCanvas(size int, hexColor string) (*image.RGBA, error) {
	width, height := size, size
	bgColor, err := Hex2Rgba(hexColor)
	if err != nil {
		log.Fatal(err)
	}
	background := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(background, background.Bounds(), &image.Uniform{C: bgColor},
		image.Point{}, draw.Src)
	return background, err
}

func Png2Jpg(input, output string) error {
	pngFile, err := os.Open(input)
	if err != nil {
		return ErrMsg("Error opening PNG file", err)
	}
	defer pngFile.Close()

	pngImage, err := png.Decode(pngFile)
	if err != nil {
		return ErrMsg("Error decoding PNG file", err)
	}

	jpegFile, err := os.Create(output)
	if err != nil {
		return ErrMsg("Error creating JPEG file", err)
	}
	defer jpegFile.Close()

	err = jpeg.Encode(jpegFile, pngImage, nil)
	if err != nil {
		return ErrMsg("Error encoding image to JPEG", err)
	}

	return nil
}
