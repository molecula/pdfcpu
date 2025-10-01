package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s input.pdf output.pdf\n", os.Args[0])
		os.Exit(1)
	}

	inFile := os.Args[1]
	outFile := os.Args[2]

	conf := model.NewDefaultConfiguration()
	conf.OptimizeResourceDicts = false
	conf.ValidationMode = model.ValidationRelaxed
	conf.Offline = true
	conf.Optimize = false
	conf.DecodeAllStreams = false
	conf.WriteXRefStream = false

	// Trim directly without optimization
	fmt.Println("Trimming PDF to pages 1-5...")
	err := api.TrimFile(inFile, outFile, []string{"1-5"}, conf)
	if err != nil {
		log.Fatalf("Error extracting pages: %v", err)
	}

	fmt.Printf("Successfully extracted pages 1-5 from %s to %s\n", inFile, outFile)
}
