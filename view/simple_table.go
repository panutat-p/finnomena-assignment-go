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
			{Align: simpletable.AlignCenter, Text: "mstar_id"},
			{Align: simpletable.AlignCenter, Text: "thailand_fund_code"},
			{Align: simpletable.AlignCenter, Text: "nav_return"},
			{Align: simpletable.AlignCenter, Text: "nav"},
			{Align: simpletable.AlignCenter, Text: "nav_date (UTC)"},
			{Align: simpletable.AlignCenter, Text: "avg_return"},
		},
	}

	for i, row := range *data {
		i += 1
		r := []*simpletable.Cell{
			{Align: simpletable.AlignRight, Text: fmt.Sprintf("%d", i)},
			{Text: row.MstarID},
			{Text: row.ThailandFundCode},
			{Align: simpletable.AlignRight, Text: fmt.Sprintf("$ %.2f", row.NavReturn)},
			{Align: simpletable.AlignRight, Text: fmt.Sprintf("$ %.2f", row.Nav)},
			{Text: row.NavDate.Format("2006-02-01")},
			{Align: simpletable.AlignRight, Text: fmt.Sprintf("$ %.2f", row.AvgReturn)},
		}

		table.Body.Cells = append(table.Body.Cells, r)
	}

	table.SetStyle(simpletable.StyleCompactLite)
	fmt.Println(table.String())
}
