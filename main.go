package main

import (
	"log"

	"dead_modules/ui"
)

func main() {
	if err := ui.StartApp(); err != nil {
		log.Fatalf("Error running application: %v", err)
	}
}
