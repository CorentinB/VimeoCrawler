package main

import (
	"fmt"
	"os"

	"github.com/akamensky/argparse"
)

var arguments = struct {
	Concurrency int
	Verbose     bool
}{}

func parseArgs(args []string) {
	// Create new parser object
	parser := argparse.NewParser("VimeoCrawler", "Crawler for vimeo.com")

	concurrency := parser.Int("j", "concurrency", &argparse.Options{
		Required: false,
		Help:     "Concurrency",
		Default:  4})

	verbose := parser.Flag("v", "verbose", &argparse.Options{
		Required: false,
		Help:     "Verbose output",
		Default:  false})

	// Parse input
	err := parser.Parse(args)
	if err != nil {
		// In case of error print error and print usage
		// This can also be done by passing -h or --help flags
		fmt.Print(parser.Usage(err))
		os.Exit(0)
	}

	// Fill arguments structure
	arguments.Concurrency = *concurrency
	arguments.Verbose = *verbose
}
