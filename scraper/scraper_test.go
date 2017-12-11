package scraper

import (
	"fmt"
	"os"
	"testing"
	"time"
)

const URL_ENV = "YNNI_OGWEN_URL"

func getURL() (string, error) {
	if url := os.Getenv(URL_ENV); url != "" {
		return url, nil
	}
	return "", fmt.Errorf("Environmental variable %q missing", URL_ENV)
}

func TestPollTurbine(t *testing.T) {
	url, err := getURL()
	if err != nil {
		t.Error(err)
		return
	}

	data, err := PollTurbine(url, time.Hour)
	if err != nil {
		t.Error(err)
	}

	if len(data) != 15 {
		t.Errorf("Expected 15 keys, but got %d", len(data))
	}

	for _, expectedKey := range []string{
		"Power Output",
		"River Level",
		"Sump Level",
		"Speed",
		"GV1 position",
		"GV2 position",
		"G59 mains OK",
		"Softstarter Closed",
		"mWh",
		"kWh",
		"Gearbox Temperature",
		"P1 Mains Voltage",
		"Time (hours)",
		"Time (mins)",
		"Theoretical Power",
	} {
		if _, ok := data[expectedKey]; !ok {
			t.Errorf("Expected key %q was not found", expectedKey)
		}
	}
}
