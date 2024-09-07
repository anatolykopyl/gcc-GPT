package main

import (
	"log"
	"os"
)

func main() {
	sourceFile := os.Args[1]
	content, err := os.ReadFile(sourceFile)
	if err != nil {
		log.Fatalf("Error reading file: %v\n", err)
	}

	assemblyCode, err := GenerateAssembly(string(content))
	if err != nil {
		log.Fatalf("Error generating assembly: %v\n", err)
	}

	os.WriteFile("a.out", []byte(assemblyCode), 0644)
}
