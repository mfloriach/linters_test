package main

import (
	"fmt"
	"hoge/cmd/metrics/metrics"
	"math"
	"os"
	"text/template"
)

const (
	templateHtml = "assets/report.template.html"
	parsedHtml   = "summarize.html"
)

type Report struct {
	LinesOfCode int
	Modules     []Module
}

type Metrics struct {
	Avg             float32
	PropagationCost float32
}

type Module struct {
	Name              string
	Lines             int
	Files             int
	PercentageOfLines float64
	Metrics           Metrics
	Graph             string
	SubModules        []Module
}

func main() {
	modules := metrics.Code()
	metrics.Coupling()

	report := Report{}
	for _, metric := range modules {
		report.LinesOfCode += metric.Code
	}

	var moduleMetric []Module
	for name, metric := range modules {
		moduleMetric = append(moduleMetric, Module{
			Name:              name,
			Lines:             metric.Code,
			PercentageOfLines: math.Floor((float64(metric.Code) / float64(report.LinesOfCode) * 100 * 100) / 100),
			Files:             metric.Files,
			Graph:             "./assets/images/" + name + ".png",
			Metrics: Metrics{
				Avg:             1.2,
				PropagationCost: 23,
			},
			SubModules: []Module{
				{
					Name:              "interfaces",
					Lines:             12,
					PercentageOfLines: 45,
					Files:             1,
				},
				{
					Name:              "repositories",
					Lines:             89,
					PercentageOfLines: 45,
					Files:             4,
				},
				{
					Name:              "services",
					Lines:             83,
					PercentageOfLines: 4,
					Files:             1,
				},
			},
		})
	}

	report.Modules = moduleMetric
	fmt.Println(report)

	tmpl, err := template.ParseFiles(templateHtml,
		"assets/header.html",
		"assets/summarize.html",
		"assets/module_details.html",
		"assets/modules.html",
		"assets/compliance.html",
		"assets/metrics.html",
		"assets/tests.html",
		"assets/endpoints.html",
	)
	if err != nil {
		panic(err)
	}

	f, err := os.Create(parsedHtml)
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(f, report)
	if err != nil {
		panic(err)
	}
	f.Close()

}
