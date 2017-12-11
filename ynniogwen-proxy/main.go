package main

// Provides a JSON & Prometeus HTTP proxy for the Ynni Ogwen turbine.
// JSON data is cached for 1 minute, Prometheus for 5 seconds.

import (
	"context"
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/lukepfarrar/ynniogwen/scraper"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const timeout = 15

func registerGauge(name, help string) prometheus.Gauge {
	g := prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "ynni_ogwen",
		Subsystem: "turbine",
		Name:      name,
		Help:      help,
	})
	if err := prometheus.Register(g); err != nil {
		log.Fatalf("%q guage could not be registered: %v", name, err)
	}
	return g
}

func errStatus(ctx context.Context) int {
	if ctx.Err() != nil {
		return http.StatusRequestTimeout
	}
	return http.StatusInternalServerError
}

func jsonHandler(url string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
		defer cancel()
		values, err := scraper.PollTurbine(ctx, url, time.Minute) // data must be no more than a minute old
		if err != nil {
			log.Println(err)
			w.WriteHeader(errStatus(ctx))
			return
		}

		enc := json.NewEncoder(w)
		enc.SetIndent("", "    ")
		if err := enc.Encode(values); err != nil {
			log.Printf("Error: %v", err)
			w.WriteHeader(500)
		}
	}
}

type guageMap map[string]prometheus.Gauge

func turbineHandler(url string, guages guageMap) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
		defer cancel()
		values, err := scraper.PollTurbine(ctx, url, 5*time.Second) // data must no more than 5 seconds old
		if err != nil {
			log.Println(err)
			w.WriteHeader(errStatus(ctx))
			return
		}
		for label, guage := range guages {
			guage.Set(values[label])
		}
		promhttp.Handler().ServeHTTP(w, r)
	}
}

func main() {
	url := flag.String("url", "", "The turbine data URL. Mandatory.")
	flag.Parse()

	if *url == "" {
		flag.Usage()
		os.Exit(1)
	}

	guages := guageMap{
		"Power Output":        registerGauge("power_output", "Power output of the turbine (KWh)."),
		"River Level":         registerGauge("river_level", "River level (mm?)"),
		"Sump Level":          registerGauge("sump_level", "Sump level (unit unknown)"),
		"Speed":               registerGauge("speed", "Speed (RPM?)"),
		"GV1 position":        registerGauge("gv1_position", "GV1 position."),
		"GV2 position":        registerGauge("gv2_position", "GV2 position."),
		"G59 mains OK":        registerGauge("g59_mains_ok", "Mains OK (boolean)."),
		"Softstarter Closed":  registerGauge("softstarted_closed", "Soft starter closed (boolean)."),
		"Gearbox Temperature": registerGauge("gearbox_temperature", "Gearbox temperature (°C)"),
		"P1 Mains Voltage":    registerGauge("p1_mains_voltage", "Mains Voltage."),
		"Theoretical Power":   registerGauge("theoretical_power", "Theoretical power output (KWh)."),
		"mWh":                 registerGauge("mWh", "Megawatt hours."),
		"kWh":                 registerGauge("kWh", "Kilowatt hours."),
	}

	http.HandleFunc("/", jsonHandler(*url))
	http.HandleFunc("/metrics", turbineHandler(*url, guages))

	log.Println("Listening...")
	log.Fatal(http.ListenAndServe(":9000", nil))
}
