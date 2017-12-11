package scraper

import (
	"context"
	"encoding/xml"
	"github.com/pkg/errors"
	"net/http"
	"strconv"
	"sync"
	"time"

	"golang.org/x/net/context/ctxhttp"
)

/*
"Power Output": 99.000000
"River Level": 240.000000
"Sump Level": 7.000000
"Speed": 400.000000
"GV1 position": 68.000000
"GV2 position": 98.000000
"G59 mains OK": 1.000000
"Softstarter Closed": 1.000000
"mWh": 241.000000
"kWh": 258.000000
"Gearbox Temperature": 56.000000
"P1 Mains Voltage": 252.000000
"Time (hours)": 11.000000
"Time (mins)": 53.000000
"Theoretical Power": 98.000000
*/

type Value struct {
	Label string `xml:"label"`
	Value string `xml:"span"`
}

type HTML struct {
	XMLName xml.Name `xml:"html"`
	Values  []Value  `xml:"body>div"`
}

var (
	retrievedAt time.Time
	cached      map[string]float64
	cacheLock   = &sync.RWMutex{}
)

// PollTurbine data - Retrieves & scrapes turbine data
//                    May return cached data newer than time.Now()-recency
func PollTurbine(ctx context.Context, url string, recency time.Duration) (map[string]float64, error) {

	// Check cache
	cacheLock.RLock()
	if retrievedAt.Add(recency).After(time.Now()) {
		defer cacheLock.RUnlock()
		return cached, nil
	}
	cacheLock.RUnlock()

	cacheLock.Lock()
	defer cacheLock.Unlock()

	// Retrieve as XML/HTML
	resp, err := ctxhttp.Get(ctx, http.DefaultClient, url)
	if err != nil {
		return nil, errors.Wrapf(err, "Error fetching turbine values from %s", url)
	}
	defer resp.Body.Close()

	dec := xml.NewDecoder(resp.Body)
	var p HTML
	if err := dec.Decode(&p); err != nil {
		return nil, errors.Wrap(err, "Problem parsing turbine data")
	}

	// Copy to map
	cached = make(map[string]float64)
	for _, v := range p.Values {
		fv, err := strconv.ParseFloat(v.Value, 64)
		if err != nil {
			return nil, errors.Wrapf(err, "Error parsing float for %q", v.Label)
		}
		cached[v.Label] = fv
	}
	retrievedAt = time.Now()
	return cached, nil
}
