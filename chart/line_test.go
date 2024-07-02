package chart

import (
	"os"
	"testing"
)

const line_output = "../testdata/line_chart.html"

func TestCreateLineChart(t *testing.T) {
	options := &LineChartOpts{
		Output:  line_output,
		Title:    "Test Title",
		Subtitle: "Test Subtitle",
		XAxis:    []string{"A", "B", "C"},
		Items: []LineChartItem{
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

	CreateLineChart(options)

	if _, err := os.Stat(line_output); os.IsNotExist(err) {
		t.Errorf("Expected output file 'line_test.html' not found")
	}

	_ = os.Remove(line_output)
}

func TestGenerateLineItems(t *testing.T) {
	values := []float64{1.0, 2.0, 3.0}
	result := generateLineItems(values)

	if len(result) != len(values) {
		t.Errorf("Expected %d items, but got %d", len(values), len(result))
	}

	for i, v := range values {
		if result[i].Value != v {
			t.Errorf("Expected value %.2f at index %d, but got %.2f", v, i, result[i].Value)
		}
	}
}
