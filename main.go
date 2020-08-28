package main

import (
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/go-echarts/go-echarts/charts"
)

func handler(w http.ResponseWriter, _ *http.Request) {
	log.Debug("creating graph...")
	nameItems := []string{"one", "tow", "three", "four", "five", "six"}
	bar := charts.NewBar()
	bar.SetGlobalOptions(charts.TitleOpts{Title: "Bar Graph Test"})
	bar.AddXAxis(nameItems).
		AddYAxis("A", []int{20, 30, 40, 10, 24, 36}).
		AddYAxis("B", []int{35, 14, 25, 60, 44, 23})
	f, err := os.Create("bar.html")
	if err != nil {
		log.Println(err)
	}
	log.Debug("rendering...")
	bar.Render(w, f)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8081", nil)
}
