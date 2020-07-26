package main

import (
	"log"
	"net/http"
	"math/rand"
	"time"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

func main () {
	histogram := promauto.NewHistogram(prometheus.HistogramOpts{
		Name:    "histogram_metric",
		Buckets: []float64{1.0, 2.0, 3.0, 4.0, 5.0},
	})
	go func() {
		for {
			histogram.Observe(rand.Float32() * 5.0)
			time.Sleep(1 * time.Second)	
		}
	}()
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
