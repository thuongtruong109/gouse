package gouse

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateBarChart(t *testing.T) {
	const bar_output = "./mockdata/bar_chart.html"

	options := &IBarChartOpts{
		Output:   bar_output,
		Title:    "Test Title",
		Subtitle: "Test Subtitle",
		XAxis:    []string{"A", "B", "C"},
		Items: []IBarChartItem{
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

func TestCreateLineChart(t *testing.T) {
	const line_output = "./mockdata/line_chart.html"

	options := &ILineChartOpts{
		Output:   line_output,
		Title:    "Test Title",
		Subtitle: "Test Subtitle",
		XAxis:    []string{"A", "B", "C"},
		Items: []ILineChartItem{
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

func TestCreatePieChart(t *testing.T) {
	const pie_output = "./mockdata/pie_chart.html"

	tests := []struct {
		name           string
		opts           *IPieChartOpts
		expectedOutput string
	}{
		{
			name: "Test basic pie chart",
			opts: &IPieChartOpts{
				Output:    pie_output,
				Title:     "Test Pie Chart",
				Subtitle:  "Test subtitle",
				Radius:    50,
				Format:    "{b}: {c}",
				ShowLabel: true,
				Items: []IPieChartItem{
					{Name: "A", Values: 100},
					{Name: "B", Values: 200},
					{Name: "C", Values: 300},
				},
			},
			expectedOutput: pie_output,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CreatePieChart(tt.opts)

			_, err := os.Stat(tt.expectedOutput)
			assert.NoError(t, err, "output file should be created")
		})
	}
}
