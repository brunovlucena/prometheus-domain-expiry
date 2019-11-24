// Prometheur Exporter. Domain Expiry.
// Author: Bruno Lucena <bvlg900f@gmail.com>

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/brunovlucena/prometheus-domain-expiry/src/collector"
	"github.com/brunovlucena/prometheus-domain-expiry/src/utils"
)

var (
	scrapedDomains = flag.String(
		"scraped-domains", os.Getenv("DOMAINS"),
		"List of domains to be scraped",
	)
	fetchInterval = flag.Duration(
		"fetch-interval", 100*time.Second,
		"Interval to fetch stats",
	)
	listenAddress = flag.String(
		"web.listen-address", ":"+os.Getenv("PORT"),
		"Address to listen on for web interface and telemetry.",
	)
	metricPath = flag.String(
		"web.telemetry-path", "/metrics",
		"Path under which to expose metrics.",
	)
	domainsExpiration = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "domain_expiration_days",
			Help: "Domain Expirations days.",
		},
		[]string{"hostname"},
	)
)

func main() {
	flag.Parse()

	log.Println("Start scraping...")

	prometheus.MustRegister(domainsExpiration)

	// start scraping
	domains := strings.Split(*scrapedDomains, " ")
	// TODO: Verify if the concorrency (updates) is occouring
	go utils.ForInterval(func() {
		for _, d := range domains {
			domainsExpiration.WithLabelValues(d).Set(float64(collector.VerifyExpire(d)))
		}
	}, *fetchInterval)

	http.Handle(*metricPath, promhttp.Handler())
	log.Printf("Starting Server: %s", *listenAddress)

	// Adding healthcheck
	http.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "live\n")
	})
	log.Fatal(http.ListenAndServe(*listenAddress, nil))

}
