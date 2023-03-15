package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type user struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var users = []user{{
	ID:   0,
	Name: "Alice",
}, {
	ID:   2,
	Name: "John",
}}

var (
	httpRequestsTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Count of all HTTP requests",
	}, []string{"code", "method"})
)

func init() {
	prometheus.MustRegister(httpRequestsTotal)
}

func main() {
	slowDuration := time.Millisecond * 10

	fmt.Printf("starting_server: :9999\n")

	http.Handle("/slow", Handler{
		AdditionalLatency: slowDuration,
	})

	http.Handle("/fast", Handler{
		AdditionalLatency: time.Millisecond * 0,
	})

	foundHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(r.RequestURI))
	})

	http.HandleFunc("/test", promhttp.InstrumentHandlerCounter(httpRequestsTotal, foundHandler))

	http.HandleFunc("/header", promhttp.InstrumentHandlerCounter(httpRequestsTotal, foundHandler))

	http.HandleFunc("/twoheader", promhttp.InstrumentHandlerCounter(httpRequestsTotal, foundHandler))

	http.Handle("/metrics", promhttp.Handler())

	log.Fatal(http.ListenAndServe(":9999", nil))
}

type Handler struct {
	AdditionalLatency time.Duration
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	time.Sleep(h.AdditionalLatency)

	diff := time.Since(start)
	fmt.Fprintf(w, "Took: %s\n", diff)
}
