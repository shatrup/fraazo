package main

import (
	"frazzo/api"
	"log"
	"time"
)

func main() {
	start := time.Now()
	cache := map[string]string{}
	a := api.API{Cache: cache}
	err := a.LoadCSVFile()

	if err != nil {
		log.Fatal("Error while reading CSV...", err.Error())
	}

	log.Println("Total time in updating cache: ", time.Since(start))
	
	a.Start()
}

