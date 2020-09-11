package graphs

import (
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/go-echarts/go-echarts/charts"
	"github.com/gorilla/mux"
)

type server struct {
	Router *mux.Router
}

func (s *server) Routes() {
	s.Router.HandleFunc("/{zipcode}", s.handleGraph()).Methods("HEAD", "GET")
}

func (s *server) handleGraph() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		params := mux.Vars(req)
		zip := params["zipcode"]
		log.Debugf("creating graph for %s...", zip)

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
		return
	}
}

func NewServer(r *mux.Router) *server {
	s := &server{
		Router: r,
	}
	return s
}
