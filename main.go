package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"os"
)

func main() {
	var ask map[float64]float64
	f, err := os.Open("../data/asks.gob")
	if err != nil {
		log.Fatal(err)
	}
	if err := gob.NewDecoder(f).Decode(&ask); err != nil {
		log.Fatal(err)
	}

	for price, size := range ask {
		fmt.Println("%.2f - %.2f\n")
	}
}
