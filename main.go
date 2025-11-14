package main

import (
	"math/rand"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/snapshot-chromedp/render"
)

var (
	itemCnt = 7
	weeks   = []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
)

func generateBarItems() []opts.BarData {
	items := make([]opts.BarData, 0)
	for i := 0; i < itemCnt; i++ {
		items = append(items, opts.BarData{Value: rand.Intn(300)})
	}
	return items
}

func barTitle() *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{
			BackgroundColor: "#FFFFFF",
		}),
		// Don't forget disable the Animation
		charts.WithAnimation(false),
		charts.WithTitleOpts(opts.Title{
			Title:    "title and legend options",
			Subtitle: "go-echarts is an awesome chart library written in Golang",
			Link:     "https://github.com/go-echarts/go-echarts",
			Right:    "40%",
		}),
		charts.WithLegendOpts(opts.Legend{Right: "80%"}),
	)
	bar.SetXAxis(weeks).
		AddSeries("Category A", generateBarItems()).
		AddSeries("Category B", generateBarItems())
	return bar
}

func barTooltip() *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithAnimation(false),
		charts.WithTitleOpts(opts.Title{Title: "tooltip options"}),
		charts.WithLegendOpts(opts.Legend{Right: "80px"}),
	)
	bar.SetXAxis(weeks).
		AddSeries("Category A", generateBarItems()).
		AddSeries("Category B", generateBarItems())
	return bar
}

func main() {
	barTitleChart := barTitle()
	barTooltipChart := barTooltip()

	render.MakeChartSnapshot(barTitleChart.RenderContent(), "my-bar-title.png")
	render.MakeChartSnapshot(barTooltipChart.RenderContent(), "my-bar-tooltip.jpg")
}
