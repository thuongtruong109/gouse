package gouse

import (
	"fmt"
	"image/color"
	"math"
)

func Hue2Rgb(p, q, t float64) float64 {
	if t < 0 {
		t += 360
	}
	if t > 360 {
		t -= 360
	}
	if t < 60 {
		return p + (q-p)*t/60
	}
	if t < 180 {
		return q
	}
	if t < 240 {
		// p + (q-p)*(240-t)/60
		return p + (q-p)*(t-180)/60
	}
	return p
}

func Hlsa2Rgba(h, l, s, a float64) color.RGBA {
	var (
		r, g, b float64
	)
	if s == 0 {
		r = l
		g = l
		b = l
	} else {
		var (
			q float64
			p float64
		)
		if l < 0.5 {
			q = l * (1 + s)
		} else {
			q = l + s - l*s
		}
		p = 2*l - q
		r = Hue2Rgb(p, q, h+120)
		g = Hue2Rgb(p, q, h)
		b = Hue2Rgb(p, q, h-120)
	}

	return color.RGBA{R: uint8(r * 255), G: uint8(g * 255), B: uint8(b * 255), A: uint8(a * 255)}
}

func Rgba2Hlsa(r, g, b, a uint8) (float64, float64, float64, float64) {
	// Normalize the RGB values to the range [0, 1]
	rf := float64(r) / 255.0
	gf := float64(g) / 255.0
	bf := float64(b) / 255.0

	min := math.Min(rf, math.Min(gf, bf))
	max := math.Max(rf, math.Max(gf, bf))
	delta := max - min

	l := (max + min) / 2.0

	// Saturation
	var s float64
	if delta == 0 {
		s = 0
	} else {
		if l < 0.5 {
			s = delta / (max + min)
		} else {
			s = delta / (2.0 - max - min)
		}
	}

	// Hue
	var h float64
	if delta == 0 {
		h = 0
	} else {
		if max == rf {
			h = (gf - bf) / delta
		} else if max == gf {
			h = (bf-rf)/delta + 2.0
		} else {
			h = (rf-gf)/delta + 4.0
		}
		h *= 60.0
		if h < 0 {
			h += 360.0
		}
	}

	// Alpha (transparency)
	af := float64(a) / 255.0

	return h, s, l, af
}

func Hex2Hlsa(hex string) (float64, float64, float64, float64, error) {
	rgba, err := Hex2Rgba(hex)
	if err != nil {
		return 0, 0, 0, 0, err
	}

	h, s, l, a := Rgba2Hlsa(rgba.R, rgba.G, rgba.B, rgba.A)

	return h, s, l, a, nil
}

func Hex2Rgba(hex string) (color.RGBA, error) {
	var (
		rgba             color.RGBA
		err              error
		errInvalidFormat = fmt.Errorf("invalid")
	)
	rgba.A = 0xff
	if hex[0] != '#' {
		return rgba, errInvalidFormat
	}
	hexToByte := func(b byte) byte {
		switch {
		case b >= '0' && b <= '9':
			return b - '0'
		case b >= 'a' && b <= 'f':
			return b - 'a' + 10
		case b >= 'A' && b <= 'F':
			return b - 'A' + 10
		}
		err = errInvalidFormat
		return 0
	}
	switch len(hex) {
	case 7:
		rgba.R = hexToByte(hex[1])<<4 + hexToByte(hex[2])
		rgba.G = hexToByte(hex[3])<<4 + hexToByte(hex[4])
		rgba.B = hexToByte(hex[5])<<4 + hexToByte(hex[6])
	case 4:
		rgba.R = hexToByte(hex[1]) * 17
		rgba.G = hexToByte(hex[2]) * 17
		rgba.B = hexToByte(hex[3]) * 17
	default:
		err = errInvalidFormat
	}
	return rgba, err
}
