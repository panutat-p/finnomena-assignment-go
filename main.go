package main

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/panutat-p/finn-thai-funds-go/core"
	"github.com/panutat-p/finn-thai-funds-go/view"
)

func main() {
	for {
		prompt := promptui.Select{
			Label: "Select Time Range",
			Items: []string{"1Y", "1M", "1W", "1D", "🟥"},
		}

		_, result, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			break
		}

		switch result {
		case "1Y":
			funds := core.GetFundByRanking("fund-ranking-1Y.json")
			view.PrintCompactTable(&funds)
		case "1M":
			funds := core.GetFundByRanking("fund-ranking-1M.json")
			view.PrintCompactTable(&funds)
		case "1W":
			funds := core.GetFundByRanking("fund-ranking-1W.json")
			view.PrintCompactTable(&funds)
		case "1D":
			funds := core.GetFundByRanking("fund-ranking-1D.json")
			view.PrintCompactTable(&funds)
		case "🟥":
			return
		}
	}
}
