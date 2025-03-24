package samples

import (
	"github.com/thuongtruong109/gouse"
)

/*
Description: Create a bar chart and export it to a html file
Input params: (*gouse.IBarChart)
*/
func ChartBar() {
	newChart := &gouse.IBarChart{
		Output:   "test_sample_data/bar.html",
		Title:    "Bar chart in Go",
		Subtitle: "This is fun to use!",
		XAxis:    []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun"},
		Items: []gouse.IBarChartItem{
			{Name: "Category A", Values: []float64{100, 200, 300, 400, 500, 600}},
			{Name: "Category B", Values: []float64{200, 300, 400, 500, 600, 700}},
			{Name: "Category C", Values: []float64{300, 400, 500, 600, 700, 800}},
		},
	}

	gouse.BarChart(newChart)
}

/*
Description: Create a line chart and export it to a html file
Input params: (*gouse.ILineChart)
*/
func ChartLine() {
	newChart := &gouse.ILineChart{
		Output:   "test_sample_data/line.html",
		Title:    "Line chart in Go",
		Subtitle: "This is fun to use!",
		XAxis:    []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"},
		Items: []gouse.ILineChartItem{
			{Name: "Category A", Values: []float64{70, 200, 10, 300, 310, 900}},
			{Name: "Category B", Values: []float64{680, 290, 356, 434, 900, 100}},
		},
	}

	gouse.LineChart(newChart)
}

/*
Description: Create a pie chart and export it to a html file
Input params: (*gouse.IPieChart)
*/
func ChartPie() {
	newChart := &gouse.IPieChart{
		Output:     "test_sample_data/pie.html",
		Title:      "Pie chart in Go",
		Subtitle:   "This is fun to use!",
		Annotation: "Monthly revenue",
		Radius:     200,
		Format:     "{b}: {c} ({d}%)",
		ShowLabel:  true,
		Items: []gouse.IPieChartItem{
			{Name: "Category A", Values: 335},
			{Name: "Category B", Values: 310},
			{Name: "Category C", Values: 234},
			{Name: "Category D", Values: 135},
			{Name: "Category E", Values: 1548},
			{Name: "Category F", Values: 1548},
		},
	}

	gouse.PieChart(newChart)
}

/*
Description: Create a scatter chart and export it to a html file
Input params: (*gouse.IScatterChart)
*/
func ChartScatter() {
	newChart := &gouse.IScatterChart{
		Output:     "test_sample_data/scatter.html",
		Title:      "Scatter chart in Go",
		Subtitle:   "This is fun to use!",
		Annotation: "Temperature",
		XAxis: []string{"Jan 1", "Jan 10", "Jan 12", "Jan 20", "Jan 30", "Feb 1",
			"Feb 2", "Feb 5", "Feb 8", "Feb 12"},
		Items: []float64{-7.3, -3.4, -5.0, -0.9, -2.2, 4.8, 5.1, -1.9, 0, 2.6},
	}

	gouse.ScatterChart(newChart)
}
