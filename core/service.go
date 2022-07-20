package core

import (
	"encoding/json"
	"fmt"
	"github.com/panutat-p/finn-thai-funds-go/config"
	"github.com/panutat-p/finn-thai-funds-go/fund"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
)

type Response struct {
	Status bool   `json:"status"`
	Error  string `json:"error"`
	Data   []fund.Fund
}

func GetFundByRanking(tf string) []fund.Fund {
	r, _ := http.Get(fmt.Sprintf("%s/%s", config.BASE_URL, tf))
	fmt.Println(r.Status)
	fmt.Println(r.Header["Content-Type"])

	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		panic(err)
	}

	var response Response
	err = json.Unmarshal(buf, &response)
	if err != nil {
		log.Println(err)
		panic(err)
	}

	return response.Data
}

func SortFundsByNavDate(funds []fund.Fund) {
	sort.Slice(funds, func(i, j int) bool {
		return funds[i].NavDate.Before(funds[j].NavDate)
	})
}
