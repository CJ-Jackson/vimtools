package main

import (
	"encoding/json"
	"log"
	"os"
)

func main() {
	var m any

	if err := json.NewDecoder(os.Stdin).Decode(&m); err != nil {
		log.Fatalln(err)
	}

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	if err := enc.Encode(m); err != nil {
		log.Fatalln(err)
	}
}
