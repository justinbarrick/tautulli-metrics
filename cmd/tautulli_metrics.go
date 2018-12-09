package main

import (
	"github.com/justinbarrick/tautulli-metrics/pkg/metrics"
	"github.com/justinbarrick/tautulli-metrics/pkg/tautulli"
	"log"
	"os"
	"time"
)

func main() {
	metricsUrl := os.Getenv("INFLUX_URL")
	metricsDb := os.Getenv("INFLUX_DB")

	if metricsUrl == "" || metricsDb == "" {
		log.Fatal("Both INFLUX_URL and INFLUX_DB must be set.")
	}

	metrics, err := metrics.NewMetrics(metricsUrl, metricsDb, os.Getenv("INFLUX_USER"), os.Getenv("INFLUX_PASS"))
	if err != nil {
		log.Fatal(err)
	}

	tautulliUrl := os.Getenv("TAUTULLI_URL")
	tautulliApiKey := os.Getenv("TAUTULLI_API_KEY")

	if tautulliUrl == "" || tautulliApiKey == "" {
		log.Fatal("Both TAUTULLI_URL and TAUTULLI_API_KEY must be set.")
	}

	tautulliApi := tautulli.NewTautulliAPI(tautulliUrl, tautulliApiKey)

	lastTimestamp, err := metrics.MostRecentHistoryTimestamp()
	if err != nil {
		log.Fatal(err)
	}

	for {
		log.Println("Getting watch history")
		streams, err := tautulliApi.GetHistory(lastTimestamp)
		if err != nil {
			log.Fatal(err)
		}

		err = metrics.InsertHistory(streams)
		if err != nil {
			log.Fatal(err)
		}

		if len(streams) > 0 {
			log.Printf("Inserted %d item(s)\n", len(streams))
			lastTimestamp = streams[0].Started
		}

		time.Sleep(10 * time.Second)
	}
}
