package gouse

import (
	"image/color"
	"math"
	"testing"
)

func TestHex2Rgba(t *testing.T) {
	tests := []struct {
		hex  string
		want color.RGBA
	}{
		{"#ffffff", color.RGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff}},
		{"#000000", color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0xff}},
		{"#ff0000", color.RGBA{R: 0xff, G: 0x00, B: 0x00, A: 0xff}},
		{"#00ff00", color.RGBA{R: 0x00, G: 0xff, B: 0x00, A: 0xff}},
		{"#0000ff", color.RGBA{R: 0x00, G: 0x00, B: 0xff, A: 0xff}},
		{"#f00", color.RGBA{R: 0xff, G: 0x00, B: 0x00, A: 0xff}},
	}

	for _, tt := range tests {
		t.Run(tt.hex, func(t *testing.T) {
			got, _ := Hex2Rgba(tt.hex)
			if got != tt.want {
				t.Errorf("Hex2Rgba() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRgba2Hlsa(t *testing.T) {
	tests := []struct {
		r, g, b, a uint8
		expectedH  float64
		expectedS  float64
		expectedL  float64
		expectedA  float64
	}{
		{255, 0, 0, 255, 0, 1, 0.5, 1},
		{0, 255, 0, 255, 120, 1, 0.5, 1},
		{0, 0, 255, 255, 240, 1, 0.5, 1},
		{255, 255, 255, 255, 0, 0, 1, 1},
		{0, 0, 0, 255, 0, 0, 0, 1},
		{128, 128, 128, 255, 0, 0, 0.5, 1},
		{255, 0, 0, 128, 0, 1, 0.5, 0.5},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			h, s, l, a := Rgba2Hlsa(test.r, test.g, test.b, test.a)

			// Allow some small error in float comparisons due to precision issues
			if math.Abs(h-test.expectedH) > 0.01 {
				t.Errorf("expected hue %v, got %v", test.expectedH, h)
			}
			if math.Abs(s-test.expectedS) > 0.01 {
				t.Errorf("expected saturation %v, got %v", test.expectedS, s)
			}
			if math.Abs(l-test.expectedL) > 0.01 {
				t.Errorf("expected lightness %v, got %v", test.expectedL, l)
			}
			if math.Abs(a-test.expectedA) > 0.01 {
				t.Errorf("expected alpha %v, got %v", test.expectedA, a)
			}
		})
	}
}

func TestHue2Rgb(t *testing.T) {
	tests := []struct {
		name     string
		p, q, t  float64
		expected float64
	}{
		{"t is greater than 360", 0.5, 1.0, 390, 0.75},
		{"t is less than 60", 0.5, 1.0, 30, 0.75},
		{"t is less than 180", 0.5, 1.0, 120, 1.0},
		{"t is greater than or equal to 240", 0.5, 1.0, 300, 0.5},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Hue2Rgb(test.p, test.q, test.t)
			if result != test.expected {
				t.Errorf("Test %q failed: expected %f, got %f", test.name, test.expected, result)
			}
		})
	}
}

func Test_Hlsa2Rgba(t *testing.T) {
	tests := []struct {
		h, l, s, a float64
		expected   color.RGBA
	}{
		// Example 1: Testing with H=0 (Red), L=0.5 (midlight), S=1 (full saturation), A=1 (full opacity)
		{
			h: 0, l: 0.5, s: 1, a: 1,
			expected: color.RGBA{R: 255, G: 0, B: 0, A: 255},
		},
		// Example 2: Testing with H=120 (Green), L=0.5 (midlight), S=1 (full saturation), A=1 (full opacity)
		{
			h: 120, l: 0.5, s: 1, a: 1,
			expected: color.RGBA{R: 0, G: 255, B: 0, A: 255},
		},
		// Example 3: Testing with H=240 (Blue), L=0.5 (midlight), S=1 (full saturation), A=1 (full opacity)
		{
			h: 240, l: 0.5, s: 1, a: 1,
			expected: color.RGBA{R: 0, G: 0, B: 255, A: 255},
		},
		// Example 4: Testing with H=0 (Red), L=0.5 (midlight), S=0 (no saturation), A=1 (full opacity) - gray color
		{
			h: 0, l: 0.5, s: 0, a: 1,
			expected: color.RGBA{R: 127, G: 127, B: 127, A: 255}, // Should return gray
		},
		// Example 5: Testing with H=0 (Red), L=0.3 (darker), S=1 (full saturation), A=0.5 (semi-transparent)
		{
			h: 0, l: 0.3, s: 1, a: 0.5,
			expected: color.RGBA{R: 153, G: 0, B: 0, A: 127}, // Corrected expected value
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			result := Hlsa2Rgba(tt.h, tt.l, tt.s, tt.a)
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}
