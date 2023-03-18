package main

import (
	"hoge/cmd/metrics/metrics"
	"os"

	"github.com/jedib0t/go-pretty/table"
)

func main() {
	metrics.Code()
	metrics.Coupling()

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Modules", "Code (%)"})
	t.AppendRows([]table.Row{
		{1, "Arya"},
		{20, "Jon"},
	})
	t.Render()
}
