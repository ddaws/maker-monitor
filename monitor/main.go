package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func recordMetrics() {
	for {
		opsProcessed.Inc()
		time.Sleep(2 * time.Second)
	}
}

var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "maker_processed_ops_total",
		Help: "The total number of processed events",
		ConstLabels: map[string]string{
			"Type": "Fake",
		},
	})
)

func main() {
	fmt.Println("Starting the Maker Monitor metrics server...")

	go recordMetrics()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe("0.0.0.0:8080", nil)
}
