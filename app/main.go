package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: MetricNamespace,
		Subsystem: Subsystem,
		Name:      "processed_ops_total",
		Help:      "The total number of processed events",
	})
	myCounter = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: MetricNamespace,
		Subsystem: Subsystem,
		Name:      "random_counter_total",
		Help:      "Random counter",
	})
)

const MetricNamespace = "my_app"
const Subsystem = "experiment"

func recordMetrics() {
	go func() {
		for {
			opsProcessed.Inc()
			time.Sleep(2 * time.Second)
		}
	}()
}

func main() {
	rand.Seed(time.Now().Unix())

	recordMetrics()

	http.Handle("/metrics", promhttp.Handler())
	ip := os.Getenv("BIND_IP")
	if ip == "" {
		ip = "0.0.0.0"
	}
	port := os.Getenv("BIND_PORT")
	if port == "" {
		port = "2112"
	}
	http.ListenAndServe(fmt.Sprintf("%s:%s", ip, port), nil)
}
