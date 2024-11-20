package gouse

import (
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
)

/* Bar chart */

type IBarChartItem struct {
	Name   string
	Values []float64
}

type IBarChartOpts struct {
	Output   string
	Title    string
	Subtitle string
	XAxis    []string
	Items    []IBarChartItem
}

func generateBarItems(values []float64) []opts.BarData {
	items := make([]opts.BarData, 0)
	for i := 0; i < len(values); i++ {
		items = append(items, opts.BarData{Value: values[i]})
	}
	return items
}

func CreateBarChart(options *IBarChartOpts) {
	bar := charts.NewBar()

	bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    options.Title,
		Subtitle: options.Subtitle,
	}))

	for _, item := range options.Items {
		bar.SetXAxis(options.XAxis).
			AddSeries(item.Name, generateBarItems(item.Values))
	}

	f, _ := os.Create(options.Output)
	_ = bar.Render(f)
}

/* Line chart */

type ILineChartItem struct {
	Name   string
	Values []float64
}

type ILineChartOpts struct {
	Output   string
	Title    string
	Subtitle string
	XAxis    []string
	Items    []ILineChartItem
}

func generateLineItems(values []float64) []opts.LineData {
	items := make([]opts.LineData, 0)
	for i := 0; i < len(values); i++ {
		items = append(items, opts.LineData{Value: values[i]})
	}
	return items
}

func CreateLineChart(options *ILineChartOpts) {
	line := charts.NewLine()

	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{
			Theme: types.ThemeInfographic,
		}),
		charts.WithTitleOpts(opts.Title{
			Title:    options.Title,
			Subtitle: options.Subtitle,
		}),
	)

	for _, item := range options.Items {
		line.SetXAxis(options.XAxis).
			AddSeries(item.Name, generateLineItems(item.Values)).
			SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: true}))
	}

	f, _ := os.Create(options.Output)
	_ = line.Render(f)
}

/* Pie chart */

type IPieChartItem struct {
	Name   string
	Values float64
}

type IPieChartOpts struct {
	Output    string
	Title     string
	Subtitle  string
	Radius    float64
	Format    string
	ShowLabel bool
	Items     []IPieChartItem
}

func generatePieItems(ele []IPieChartItem) []opts.PieData {
	items := make([]opts.PieData, 0)
	for i := 0; i < len(ele); i++ {
		items = append(items, opts.PieData{
			Name:  ele[i].Name,
			Value: ele[i].Values,
		})
	}
	return items
}

func CreatePieChart(options *IPieChartOpts) {
	pie := charts.NewPie()
	pie.SetGlobalOptions(
		charts.WithTitleOpts(
			opts.Title{
				Title:    options.Title,
				Subtitle: options.Subtitle,
			},
		),
	)
	pie.SetSeriesOptions()
	pie.AddSeries("Monthly revenue",
		generatePieItems(options.Items)).
		SetSeriesOptions(
			charts.WithPieChartOpts(
				opts.PieChart{
					Radius: options.Radius,
				},
			),
			charts.WithLabelOpts(
				opts.Label{
					Show:      options.ShowLabel,
					Formatter: options.Format,
				},
			),
		)
	f, _ := os.Create(options.Output)
	_ = pie.Render(f)
}
