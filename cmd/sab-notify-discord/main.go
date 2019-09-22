package main

import (
	"log"
)

func main() {
	parameters, err := parseParameters()
	if err != nil {
		log.Fatal(err)
	}

	err = notify(parameters)
	if err != nil {
		log.Fatal(err)
	}
}
