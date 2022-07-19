package main

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"io"
	"net/http"
	"strings"
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
			r, _ := http.Get("https://storage.googleapis.com/finno-ex-re-v2-static-staging/recruitment-test/fund-ranking-1Y.json")
			fmt.Println(r.Status)
			fmt.Println(r.Header["Content-Type"])

			buf := new(strings.Builder)
			_, err := io.Copy(buf, r.Body)
			if err != nil {
				panic(err)
			}
			fmt.Println(buf.String())
		}
	}
}
