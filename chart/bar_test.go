package chart

import (
	"os"
	"testing"
)

const bar_output = "../mockdata/bar_chart.html"

func TestCreateBarChart(t *testing.T) {
	options := &BarChartOpts{
		Output:   bar_output,
		Title:    "Test Title",
		Subtitle: "Test Subtitle",
		XAxis:    []string{"A", "B", "C"},
		Items: []BarChartItem{
			{
				Name:   "Series1",
				Values: []float64{1.0, 2.0, 3.0},
			},
			{
				Name:   "Series2",
				Values: []float64{4.0, 5.0, 6.0},
			},
		},
	}

	CreateBarChart(options)

	if _, err := os.Stat(bar_output); os.IsNotExist(err) {
		t.Errorf("Expected output file 'bar_chart.html' not found")
	}

	_ = os.Remove(bar_output)
}

func TestGenerateBarItems(t *testing.T) {
	values := []float64{1.0, 2.0, 3.0}
	result := generateBarItems(values)

	if len(result) != len(values) {
		t.Errorf("Expected %d items, but got %d", len(values), len(result))
	}

	for i, v := range values {
		if result[i].Value != v {
			t.Errorf("Expected value %.2f at index %d, but got %.2f", v, i, result[i].Value)
		}
	}
}
