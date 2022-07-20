package core

import (
	"encoding/json"
	"fmt"
	"github.com/panutat-p/finn-thai-funds-go/config"
	"github.com/panutat-p/finn-thai-funds-go/fund"
	"github.com/panutat-p/finn-thai-funds-go/view"
	"io/ioutil"
	"log"
	"net/http"
)

type Response struct {
	Status bool   `json:"status"`
	Error  string `json:"error"`
	Data   []fund.Fund
}

func GetFundRanking1Y() {
	r, _ := http.Get(config.BASE_URL)
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
	fmt.Println("status:", response.Status)
	fmt.Println("error:", response.Error)
	fmt.Println("len", len(response.Data))

	view.PrintCompactTable(&response.Data)
}
