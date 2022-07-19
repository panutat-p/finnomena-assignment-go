package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
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
