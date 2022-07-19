package main

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/panutat-p/finn-thai-funds-go/core"
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

		if result == "🟥" {
			break
		} else {
			fmt.Printf("You choose %q\n", result)
			core.GetFundRanking1Y()
		}
	}
}
