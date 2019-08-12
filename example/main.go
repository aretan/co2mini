package main

import (
    "github.com/DataDog/datadog-go/statsd"
    "github.com/gashirar/co2mini"
    "log"
)

func main() {
    var co2mini co2mini.Co2mini
    var co2 int
    var temp float64

    if err := co2mini.Connect(); err != nil {
        log.Fatal(err)
    }

    dogstatsd, err := statsd.New("127.0.0.1:8125",
        statsd.WithNamespace("co2mini."),
    )

    if err != nil {
        log.Fatal(err)
    }
	
    go func() {
        if err := co2mini.Start(); err != nil {
            log.Fatal(err)
        }
    }()

    for {
        select {
        case co2 = <-co2mini.Co2Ch:
		dogstatsd.Gauge("co2", float64(co2), []string{}, 1)
        case temp = <-co2mini.TempCh:
		dogstatsd.Gauge("temp", float64(temp), []string{}, 1)
        }
    }
}
