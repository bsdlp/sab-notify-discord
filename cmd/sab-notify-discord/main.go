package main

import (
	"flag"
	"log"
)

func main() {
	flag.Parse()
	parameters := parseParameters()
	err := notify(parameters)
	if err != nil {
		log.Fatal(err)
	}
}
