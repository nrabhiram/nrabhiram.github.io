package main

import (
	"log"
	"sapphire/internal"
)

func main() {
	// Create a new controller instance
	controller := internal.MakeController()

	command, repoPaths, outputPath := controller.ProcessFlags()

	// Execute the command
	switch *command {
	case "generate-schema":
		controller.GenerateSchema(repoPaths, outputPath)
	case "build":
		controller.Build(repoPaths, outputPath)
	default:
		log.Fatalf("Unknown command: %s", *command)
	}
}
