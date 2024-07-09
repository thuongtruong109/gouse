package media

import (
	"image/jpeg"
	"image/png"
	"os"
	"fmt"
	"errors"
)

func PngToJpg(input, output string) error {
	pngFile, err := os.Open(input)
	if err != nil {
		return formatError("Error opening PNG file", err )
	}
	defer pngFile.Close()

	pngImage, err := png.Decode(pngFile)
	if err != nil {
		return formatError("Error decoding PNG file", err)
	}

	jpegFile, err := os.Create(output)
	if err != nil {
		return formatError("Error creating JPEG file", err)
	}
	defer jpegFile.Close()

	err = jpeg.Encode(jpegFile, pngImage, nil)
	if err != nil {
		return formatError("Error encoding image to JPEG", err)
	}

	return nil
}

func formatError(message string, err error) error {
	return errors.New(fmt.Sprintf("%s: %v", message, err))
}
