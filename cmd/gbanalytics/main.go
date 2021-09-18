package main

import (
	"flag"
	"fmt"
	"os"
)

var version = "dev" // it will be set on build time

type config struct {
	showVersion bool

	// maybe it is better to have a single var for each file
	dataPath string

	// what is the question to be aswered
	query string
}

func (cfg config) validate() error {
	if cfg.showVersion {
		return nil
	}

	if cfg.dataPath == "" {
		return fmt.Errorf("missing data path parameter")
	}

	if cfg.query == "" {
		return fmt.Errorf("missing query parameter")
	}

	return nil
}

func parseArgs() config {
	var cfg config

	flag.BoolVar(&cfg.showVersion, "version", false, "Show version")

	flag.StringVar(&cfg.dataPath, "data", "", "where to find the files with the data")
	flag.StringVar(&cfg.query, "query", "", "question to be aswered")

	flag.Parse()

	return cfg
}

func usage() {
	fmt.Fprintf(
		flag.CommandLine.Output(),
		"parse data files and answer some questions. Version: %s\n",
		version,
	)

	flag.PrintDefaults()
}

func main() {
	flag.Usage = usage

	cfg := parseArgs()
	if err := cfg.validate(); err != nil {
		fmt.Println(err)
		flag.Usage()
		os.Exit(1)
	}

	if cfg.showVersion {
		fmt.Println("version:", version)
		return
	}

	fmt.Println("end")
}
