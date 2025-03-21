
# <Badge style='font-size: 1.8rem; text-shadow: 1px 1px 2px rgba(0, 0, 0, 0.3); padding: 0.35rem 0.75rem 0.35rem 0;' type='info' text='🔖 Chart' />


```go
import (
	"github.com/thuongtruong109/gouse"
)
```

## 1. Chart bar

Description: Create a bar chart and export it to a html file<br>Input params: (*gouse.IBarChartOpts)<br>

```go
func ChartBar() {
	newChart := &gouse.IBarChartOpts{
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
```

## 2. Chart line

Description: Create a line chart and export it to a html file<br>Input params: (*gouse.ILineChartOpts)<br>

```go
func ChartLine() {
	newChart := &gouse.ILineChartOpts{
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
```

## 3. Chart pie

Description: Create a pie chart and export it to a html file<br>Input params: (*gouse.IPieChartOpts)<br>

```go
func ChartPie() {
	newChart := &gouse.IPieChartOpts{
		Output:    "test_sample_data/pie.html",
		Title:     "Pie chart in Go",
		Subtitle:  "This is fun to use!",
		Radius:    200,
		Format:    "{b}: {c} ({d}%)",
		ShowLabel: true,
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
```
