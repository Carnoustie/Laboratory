package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type bondprice struct {
	Date  string  `json:"date"`
	Price float64 `json:"value"`
}

func main() {
	response, err := http.Get("https://api.riksbank.se/swea/v1/Observations/JPGVB10Y/1987-01-05")

	if err != nil {
		panic(err)
	}

	//fmt.Println("\n\n\n\nstatus: ", response.Status)

	var buffer []byte

	buffer, err = io.ReadAll(response.Body)

	var bondprices []bondprice

	err = json.Unmarshal(buffer, &bondprices)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	for _, bondprice := range bondprices {
		fmt.Println(bondprice.Price)
		fmt.Println(time.Parse("2006-01-02", bondprice.Date))
	}

	//out := string(buffer)
	//fmt.Println(out)

	//scanner := bufio.NewScanner(response.Body)
	//fmt.Println(buffer)

	//var indented bytes.Buffer
	//err = json.Indent(&indented, buffer, "", "\t")

	//fmt.Println(indented)
}
