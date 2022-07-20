package view

import (
	"fmt"
	"github.com/alexeyco/simpletable"
	"github.com/panutat-p/finn-thai-funds-go/fund"
)

func PrintCompactTable(data *[]fund.Fund) {
	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "Fund ID"},
			{Align: simpletable.AlignCenter, Text: "Code"},
			{Align: simpletable.AlignCenter, Text: "NAV Rank"},
			{Align: simpletable.AlignCenter, Text: "Price"},
			{Align: simpletable.AlignCenter, Text: "Date (UTC)"},
			{Align: simpletable.AlignCenter, Text: "Average Return"},
		},
	}

	for i, row := range *data {
		i += 1
		r := []*simpletable.Cell{
			{Align: simpletable.AlignRight, Text: fmt.Sprintf("%d", i)},
			{Text: row.MstarID},
			{Text: blue(row.ThailandFundCode)},
			{Align: simpletable.AlignRight, Text: fmt.Sprintf("%.2f", row.NavReturn)},
			{Align: simpletable.AlignRight, Text: green(fmt.Sprintf("%.2f ฿", row.Nav))},
			{Text: row.NavDate.Format("2006-02-01")},
			{Align: simpletable.AlignRight, Text: fmt.Sprintf("%.2f", row.AvgReturn)},
		}

		table.Body.Cells = append(table.Body.Cells, r)
	}

	table.SetStyle(simpletable.StyleCompactLite)
	fmt.Println(table.String())
}
