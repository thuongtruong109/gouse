
# <Badge style='font-size: 1.8rem; text-shadow: 1px 1px 2px rgba(0, 0, 0, 0.3); padding: 0.25rem 0.75rem 0.25rem 0;' type='info' text='🔖 Chart' />


```go
import (
	"github.com/thuongtruong109/gouse"
)
```

### <Badge style='font-size: 1.1rem;' type='tip' text='1. sample chart bar' />



```go
func SampleChartBar() {
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

	gouse.CreateBarChart(newChart)
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='2. sample chart line' />



```go
func SampleChartLine() {
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

	gouse.CreateLineChart(newChart)
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='3. sample chart pie' />



```go
func SampleChartPie() {
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

	gouse.CreatePieChart(newChart)
}
```
