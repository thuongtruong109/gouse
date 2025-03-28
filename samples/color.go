package samples

import (
	"fmt"

	"github.com/thuongtruong109/gouse"
)

/*
Description: Convert color from Hue to RGB format
Parameters: (p, q, t float64) float64 -> float64
*/
func ColorHueToRgb() {
	p := 0.2
	q := 0.8
	t := 120.0

	rgb := gouse.Hue2Rgb(p, q, t)

	fmt.Printf("RGB value for Hue2Rgb(%f, %f, %f) = %f\n", p, q, t, rgb)
}

/*
Description: Convert color from HLSA to RGBA format
Parameters: (h, l, s, a float64) color.RGBA -> color.RGBA
*/
func ColorHlsaToRgba() {
	h := 120.0 // Hue: 120° (green)
	l := 0.5   // Lightness: 50%
	s := 0.75  // Saturation: 75%
	a := 1.0   // Alpha: fully opaque (100%)

	rgba := gouse.Hlsa2Rgba(h, l, s, a)

	fmt.Printf("HLSA(%f, %f, %f, %f) = RGBA(%d, %d, %d, %d)\n", h, l, s, a, rgba.R, rgba.G, rgba.B, rgba.A)
}

/*
Description: Convert color from RGBA to HLSA format
Parameters: (r, g, b, a uint8) -> (float64, float64, float64, float64)
*/
func ColorRgbaToHlsa() {
	r, g, b, a := uint8(255), uint8(0), uint8(0), uint8(128) // Red with 50% opacity

	h, s, l, af := gouse.Rgba2Hlsa(r, g, b, a)

	fmt.Printf("RGBA(%d, %d, %d, %d) => HLSA: %.2f, %.2f, %.2f, %.2f\n", r, g, b, a, h, s, l, af)
}

/*
Description: Convert color from Hex to HLSA format
Parameters: (hex string) -> (float64, float64, float64, float64, error)
*/
func ColorHexToHlsa() {
	hex := "#FF5733CC" // RGBA: Red=255, Green=87, Blue=51, Alpha=204 (opaque)

	h, s, l, a, err := gouse.Hex2Hlsa(hex)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Hex: %s => HLSA: %.2f, %.2f, %.2f, %.2f\n", hex, h, s, l, a)
}

/*
Description: Convert color from Hex to RGBA format
Parameters: (hex string) -> (color.RGBA, error)
*/
func ColorHexToRgba() {
	hex := "#FF5733" // RGB: Red=255, Green=87, Blue=51

	rgba, err := gouse.Hex2Rgba(hex)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Hex: %s => RGBA: %d, %d, %d, %d\n", hex, rgba.R, rgba.G, rgba.B, rgba.A)
}
