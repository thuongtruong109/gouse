
# <Badge style='font-size: 1.8rem; text-shadow: 1px 1px 2px rgba(0, 0, 0, 0.3); padding: 0.35rem 0.75rem 0.35rem 0;' type='info' text='🔖 Color' />


```go
import (
	"fmt"
	"github.com/thuongtruong109/gouse"
)
```

## 1. Color hue to rgb

Description: Convert color from Hue to RGB format<br>Parameters: (p, q, t float64) float64 -> float64<br>

```go
func ColorHueToRgb() {
	p := 0.2
	q := 0.8
	t := 120.0

	rgb := gouse.Hue2Rgb(p, q, t)

	fmt.Printf("RGB value for Hue2Rgb(%f, %f, %f) = %f\n", p, q, t, rgb)
}
```

## 2. Color hlsa to rgba

Description: Convert color from HLSA to RGBA format<br>Parameters: (h, l, s, a float64) color.RGBA -> color.RGBA<br>

```go
func ColorHlsaToRgba() {
	h := 120.0 // Hue: 120° (green)
	l := 0.5   // Lightness: 50%
	s := 0.75  // Saturation: 75%
	a := 1.0   // Alpha: fully opaque (100%)

	rgba := gouse.Hlsa2Rgba(h, l, s, a)

	fmt.Printf("HLSA(%f, %f, %f, %f) = RGBA(%d, %d, %d, %d)\n", h, l, s, a, rgba.R, rgba.G, rgba.B, rgba.A)
}
```

## 3. Color rgba to hlsa

Description: Convert color from RGBA to HLSA format<br>Parameters: (r, g, b, a uint8) -> (float64, float64, float64, float64)<br>

```go
func ColorRgbaToHlsa() {
	r, g, b, a := uint8(255), uint8(0), uint8(0), uint8(128) // Red with 50% opacity

	h, s, l, af := gouse.Rgba2Hlsa(r, g, b, a)

	fmt.Printf("RGBA(%d, %d, %d, %d) => HLSA: %.2f, %.2f, %.2f, %.2f\n", r, g, b, a, h, s, l, af)
}
```

## 4. Color hex to hlsa

Description: Convert color from Hex to HLSA format<br>Parameters: (hex string) -> (float64, float64, float64, float64, error)<br>

```go
func ColorHexToHlsa() {
	hex := "#FF5733CC" // RGBA: Red=255, Green=87, Blue=51, Alpha=204 (opaque)

	h, s, l, a, err := gouse.Hex2Hlsa(hex)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Hex: %s => HLSA: %.2f, %.2f, %.2f, %.2f\n", hex, h, s, l, a)
}
```

## 5. Color hex to rgba

Description: Convert color from Hex to RGBA format<br>Parameters: (hex string) -> (color.RGBA, error)<br>

```go
func ColorHexToRgba() {
	hex := "#FF5733" // RGB: Red=255, Green=87, Blue=51

	rgba, err := gouse.Hex2Rgba(hex)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Hex: %s => RGBA: %d, %d, %d, %d\n", hex, rgba.R, rgba.G, rgba.B, rgba.A)
}
```
