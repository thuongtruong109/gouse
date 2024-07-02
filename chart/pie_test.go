package chart

import (
	"os"
	"testing"
	"github.com/stretchr/testify/assert"
)

const pie_output = "../testdata/pie_chart.html"

func TestCreatePieChart(t *testing.T) {
	tests := []struct {
		name           string
		opts           *PieChartOpts
		expectedOutput string
	}{
		{
			name: "Test basic pie chart",
			opts: &PieChartOpts{
				Output:    pie_output,
				Title:     "Test Pie Chart",
				Subtitle:  "Test subtitle",
				Radius:    50,
				Format:    "{b}: {c}",
				ShowLabel: true,
				Items: []PieChartItem{
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
