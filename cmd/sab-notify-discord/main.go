package main

import (
	"flag"
	"log"
)

func main() {
	cfg, err := LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	flag.Parse()
	parameters := parseParameters()
	err = notify(cfg.WebhookURL, parameters)
	if err != nil {
		log.Fatal(err)
	}
}
