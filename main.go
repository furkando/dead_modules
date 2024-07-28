package main

import (
	"dead_modules/ui"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	// Define the --version flag
	versionFlag := flag.Bool("version", false, "Print the version number and exit")
	flag.Parse()

	// Check if the --version flag was provided
	if *versionFlag {
		fmt.Printf("Dead Modules %s\n", ui.Version)
		os.Exit(0)
	}

	if err := ui.StartApp(); err != nil {
		log.Fatalf("Error running application: %v", err)
	}

}
