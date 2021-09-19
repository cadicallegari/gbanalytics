package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path"

	"github.com/cadicallegari/gbanalytics"
	"github.com/cadicallegari/gbanalytics/csv"
)

var version = "dev" // it will be set on build time

type config struct {
	showVersion bool

	// maybe it is better to have a single var for each file
	dataPath string

	// what is the question to be aswered
	query string

	// size of the rank
	limit int
}

var validQueries = []string{
	"top-active-users",
	"top-active-repos",
	"top-watch-repos",
}

func (cfg config) validate() error {
	if cfg.showVersion {
		return nil
	}

	if cfg.dataPath == "" {
		return fmt.Errorf("missing data path parameter")
	}

	if cfg.limit < 0 {
		return fmt.Errorf("limit must be greater than 0")
	}

	var queryFound bool
	for _, vq := range validQueries {
		if cfg.query == vq {
			queryFound = true
			continue
		}
	}

	if !queryFound {
		return fmt.Errorf("invalid query param")
	}

	return nil
}

func parseArgs() config {
	var cfg config

	flag.BoolVar(&cfg.showVersion, "version", false, "show version")

	flag.StringVar(&cfg.dataPath, "data", "", "where to find the files with the data")
	flag.StringVar(&cfg.query, "query", "", "question to be aswered")
	flag.IntVar(&cfg.limit, "limit", 10, "size of the top rank")

	flag.Parse()

	return cfg
}

func usage() {
	fmt.Fprintf(
		flag.CommandLine.Output(),
		"\nparse data files and answer some questions.\nvalid queries: %s\n",
		validQueries,
	)

	flag.PrintDefaults()
}

func repoNameOrID(repos map[string]*gbanalytics.Repo, id string) string {
	r, ok := repos[id]
	if !ok {
		return id
	}

	return r.Name
}

func usernameOrID(actors map[string]*gbanalytics.Actor, id string) string {
	r, ok := actors[id]
	if !ok {
		return id
	}

	return r.Username
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

	data, err := loader.Load(ctx)
	if err != nil {
		fmt.Fprintln(os.Stderr, fmt.Errorf("unable to load data: %s", err))
		os.Exit(1)
	}

	switch cfg.query {
	case "top-active-users":
		results, err := gbanalytics.MostActiveUsers(data.Events, data.CommitsByEvent, cfg.limit)
		if err != nil {
			fmt.Fprintln(os.Stderr, fmt.Errorf("unable get most active users: %s", err))
			os.Exit(1)
		}

		for i, r := range results {
			fmt.Printf("%3d | %4d - %s\n", i+1, r.Count, usernameOrID(data.Actors, r.ID))
		}

	case "top-active-repos":
		results, err := gbanalytics.MostActiveRepos(data.Events, data.CommitsByEvent, cfg.limit)
		if err != nil {
			fmt.Fprintln(os.Stderr, fmt.Errorf("unable get most active repos: %s", err))
			os.Exit(1)
		}

		for i, r := range results {
			fmt.Printf("%3d | %4d - %s\n", i+1, r.Count, repoNameOrID(data.Repos, r.ID))
		}

	case "top-watch-repos":
		results, err := gbanalytics.MostWachedRepos(data.Events, cfg.limit)
		if err != nil {
			fmt.Fprintln(os.Stderr, fmt.Errorf("unable get most watched repos: %s", err))
			os.Exit(1)
		}

		for i, r := range results {
			fmt.Printf("%3d | %4d - %s\n", i+1, r.Count, repoNameOrID(data.Repos, r.ID))
		}

	default:
		fmt.Fprintln(os.Stderr, "unknown query")
		os.Exit(1)
	}
}
