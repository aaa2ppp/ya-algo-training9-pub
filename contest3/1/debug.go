package main

import (
	"log"
	"os"
)

func init() {
	_, debug = os.LookupEnv("DEBUG")
	log.SetFlags(log.Llongfile)
}
