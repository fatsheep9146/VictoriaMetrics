package promscrape

import (
	"fmt"
	"testing"
	"time"

	"github.com/VictoriaMetrics/VictoriaMetrics/lib/auth"
	"github.com/VictoriaMetrics/VictoriaMetrics/lib/promauth"
	"github.com/VictoriaMetrics/VictoriaMetrics/lib/prompbmarshal"
	"github.com/VictoriaMetrics/VictoriaMetrics/lib/proxy"
)

func TestScrapeFromHttpServer(t *testing.T) {
	sw := &ScrapeWork{
		ScrapeURL:      "http://localhost:9001/metrics",
		ScrapeTimeout:  10 * time.Second,
		ScrapeInterval: 10 * time.Second,
		HonorLabels:    false,
		AuthConfig:     &promauth.Config{},
	}

	pushData := func(at *auth.Token, wr *prompbmarshal.WriteRequest) {
		for i, ts := range wr.Timeseries {
			fmt.Printf("ts[%v]: %v\n", i, ts)
		}
	}

	scraper := newScraper(sw, "test", pushData)
	scraper.sw.run(nil, nil)
}

func TestScrapeFromSocks5Server(t *testing.T) {
	url := proxy.MustNewURL("socks5://user:password@127.0.0.1:8000")

	sw := &ScrapeWork{
		ScrapeURL:      "http://localhost:9001/metrics",
		ScrapeTimeout:  10 * time.Second,
		ScrapeInterval: 10 * time.Second,
		HonorLabels:    false,
		ProxyURL:       url,
		AuthConfig:     &promauth.Config{},
	}

	pushData := func(at *auth.Token, wr *prompbmarshal.WriteRequest) {
		for i, ts := range wr.Timeseries {
			fmt.Printf("ts[%v]: %v\n", i, ts)
		}
	}

	scraper := newScraper(sw, "test", pushData)
	scraper.sw.run(nil, nil)
}
