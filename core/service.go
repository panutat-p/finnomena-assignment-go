package core

import (
	"encoding/json"
	"fmt"
	"github.com/panutat-p/finn-thai-funds-go/config"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Response struct {
	Status bool   `json:"status"`
	Error  string `json:"error"`
	Data   []struct {
		MstarID          string    `json:"mstar_id"`
		ThailandFundCode string    `json:"thailand_fund_code"`
		NavReturn        float64   `json:"nav_return"`
		Nav              float64   `json:"nav"`
		NavDate          time.Time `json:"nav_date"`
		AvgReturn        float64   `json:"avg_return"`
	} `json:"data"`
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
}
