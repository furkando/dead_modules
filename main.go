package main

import (
	"dead_modules/ui"
	"log"
)

func main() {
	if err := ui.StartApp(); err != nil {
		log.Fatalf("Error running application: %v", err)
	}

}
