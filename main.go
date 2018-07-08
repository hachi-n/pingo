package main

import (
	"log"
	"os"

	"github.com/hachi-n/pingo/lib/ping"
)

func main() {
	log.SetOutput(os.Stdout)

	if len(os.Args) != 2 {
		log.Fatal("argument err...")
	}

	hostname := os.Args[1]

	ping.Do(hostname)
}
