package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
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

	dataset := make(plotter.XYs, len(bondprices))
	for index, bondprice := range bondprices {
		fmt.Println(bondprice.Price)
		fmt.Println(time.Parse("2006-01-02", bondprice.Date))
		trueDate, _ := time.Parse("2006-01-02", bondprice.Date)
		dateFloat := float64(trueDate.Unix())
		dataset[index].X = dateFloat
		dataset[index].Y = bondprice.Price
	}

	p := plot.New()
	p.Title.Text = "Japanese Bond yields"
	p.X.Tick.Marker = plot.TimeTicks{Format: "2006-01-02"}
	line, err := plotter.NewLine(dataset)
	p.Add(line)
	p.X.Padding = 10
	p.Y.Padding = 10
	p.Y.Tick.Marker = plot.DefaultTicks{}
	err = p.Save(6*vg.Inch, 3*vg.Inch, "bondyields.svg")

	if err != nil {
		panic(err)
	}

	//out := string(buffer)
	//fmt.Println(out)

	//scanner := bufio.NewScanner(response.Body)
	//fmt.Println(buffer)

	//var indented bytes.Buffer
	//err = json.Indent(&indented, buffer, "", "\t")

	//fmt.Println(indented)
}
