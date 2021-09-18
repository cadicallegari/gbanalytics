package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path"

	"github.com/cadicallegari/gbanalytics/csv"
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

	loader := csv.NewLoader(csv.Config{
		ActorsFile:  path.Join(cfg.dataPath, "actors.csv"),
		CommitsFile: path.Join(cfg.dataPath, "commits.csv"),
		EventsFile:  path.Join(cfg.dataPath, "events.csv"),
		ReposFile:   path.Join(cfg.dataPath, "repos.csv"),
	})

	ctx := context.Background()

	dt, err := loader.Load(ctx)
	if err != nil {
		fmt.Fprintln(os.Stderr, fmt.Errorf("unable to load data: %s", err))
		os.Exit(1)
	}

	switch cfg.query {
	case "top-active-users":
		users, err := dt.MostActiveUsers(10)
		if err != nil {
			fmt.Fprintln(os.Stderr, fmt.Errorf("unable get most active users: %s", err))
			os.Exit(1)
		}

		for i, u := range users {
			fmt.Println(i+1, u.Username)
		}

	case "top-active-repos":
		repos, err := dt.MostActiveRepos(10)
		if err != nil {
			fmt.Fprintln(os.Stderr, fmt.Errorf("unable get most active repos: %s", err))
			os.Exit(1)
		}

		for i, r := range repos {
			fmt.Println(i+1, r.Name)
		}

	case "top-watch-repos":
		repos, err := dt.MostWachedRepos(10)
		if err != nil {
			fmt.Fprintln(os.Stderr, fmt.Errorf("unable get most watched repos: %s", err))
			os.Exit(1)
		}

		for i, r := range repos {
			fmt.Println(i+1, r.Name)
		}

	default:
		// move it to the validation step
		fmt.Fprintln(os.Stderr, "unkown query")
		os.Exit(1)
	}
}
