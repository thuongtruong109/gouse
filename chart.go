package gouse

import (
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"

	"github.com/PuerkitoBio/goquery"
)

func customTitle(dist, title string) {
	file, err1 := os.Open(dist)
	if err1 != nil {
		println(err1)
		return
	}
	defer file.Close()

	doc, err2 := goquery.NewDocumentFromReader(file)
	if err2 != nil {
		println(err2)
		return
	}

	doc.Find("title").SetText("Gouse - " + title)

	outFile, err3 := os.Create(dist)
	if err3 != nil {
		println(err3)
		return
	}
	defer outFile.Close()

	html, err4 := doc.Html()
	if err4 != nil {
		println(err4)
		return
	}

	_, err5 := outFile.WriteString(html)
	if err5 != nil {
		println(err5)
	}
}

type IBarChartItem struct {
	Name   string
	Values []float64
}

type IBarChart struct {
	Output   string
	Title    string
	Subtitle string
	XAxis    []string
	Items    []IBarChartItem
}

func generateBarItems(values []float64) []opts.BarData {
	items := make([]opts.BarData, 0)
	for i := range values {
		items = append(items, opts.BarData{Value: values[i]})
	}
	return items
}

func BarChart(options *IBarChart) {
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

	customTitle(options.Output, options.Title)
}

type ILineChartItem struct {
	Name   string
	Values []float64
}

type ILineChart struct {
	Output   string
	Title    string
	Subtitle string
	XAxis    []string
	Items    []ILineChartItem
}

func generateLineItems(values []float64) []opts.LineData {
	items := make([]opts.LineData, 0)
	for i := range values {
		items = append(items, opts.LineData{Value: values[i]})
	}
	return items
}

func LineChart(options *ILineChart) {
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
			SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: opts.Bool(true)}))
	}

	f, _ := os.Create(options.Output)
	_ = line.Render(f)

	customTitle(options.Output, options.Title)
}

type IPieChartItem struct {
	Name   string
	Values float64
}

type IPieChart struct {
	Output     string
	Title      string
	Subtitle   string
	Annotation string
	Radius     float64
	Format     string
	ShowLabel  bool
	Items      []IPieChartItem
}

func generatePieItems(ele []IPieChartItem) []opts.PieData {
	items := make([]opts.PieData, 0)
	for i := range ele {
		items = append(items, opts.PieData{
			Name:  ele[i].Name,
			Value: ele[i].Values,
		})
	}
	return items
}

func PieChart(options *IPieChart) {
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
	pie.AddSeries(options.Annotation,
		generatePieItems(options.Items)).
		SetSeriesOptions(
			charts.WithPieChartOpts(
				opts.PieChart{
					Radius: options.Radius,
				},
			),
			charts.WithLabelOpts(
				opts.Label{
					Show:      opts.Bool(options.ShowLabel),
					Formatter: options.Format,
				},
			),
		)
	f, _ := os.Create(options.Output)
	_ = pie.Render(f)

	customTitle(options.Output, options.Title)
}

type IScatterChart struct {
	Output     string
	Title      string
	Subtitle   string
	Annotation string
	XAxis      []string
	Items      []float64
}

func ScatterChart(options *IScatterChart) {
	scatter := charts.NewScatter()

	scatter.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),
		charts.WithTitleOpts(opts.Title{Title: options.Title, Subtitle: options.Subtitle}),
	)

	temps := make([]opts.ScatterData, 0)
	for i := range options.Items {
		temps = append(temps, opts.ScatterData{Value: options.Items[i]})
	}

	scatter.SetXAxis(options.XAxis).AddSeries(options.Annotation, temps)

	f, _ := os.Create(options.Output)
	scatter.Render(f)

	customTitle(options.Output, options.Title)
}
