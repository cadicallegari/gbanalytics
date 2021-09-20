package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/cadicallegari/gbanalytics"
	"github.com/cadicallegari/gbanalytics/csv"
)

var version = "dev" // it will be set on build time

type config struct {
	showVersion bool
	dataPath    string
	query       string
	limit       int
}

var validQueries = map[string]struct{}{
	"top-active-users": {},
	"top-active-repos": {},
	"top-watch-repos":  {},
}

func (cfg config) validate() error {
	if cfg.showVersion {
		return nil
	}

	if cfg.dataPath == "" {
		return fmt.Errorf("missing data path parameter")
	}

	if cfg.limit <= 0 {
		return fmt.Errorf("limit must be greater than 0")
	}

	if _, ok := validQueries[cfg.query]; !ok {
		return fmt.Errorf("invalid query param")
	}

	return nil
}

func parseArgs() config {
	var cfg config

	flag.BoolVar(&cfg.showVersion, "version", false, "show version")

	flag.StringVar(&cfg.dataPath, "path", "", "where to find the files with the data")
	flag.StringVar(&cfg.query, "query", "", "question to be aswered")
	flag.IntVar(&cfg.limit, "limit", 10, "size of the top rank")

	flag.Parse()

	return cfg
}

func usage() {
	var queries strings.Builder
	for k := range validQueries {
		queries.WriteString(k)
		queries.WriteString(", ")
	}

	fmt.Fprintf(
		flag.CommandLine.Output(),
		"\nparse files with github data and print the top rank.\nvalid queries: %s\n",
		strings.TrimRight(queries.String(), ", "),
	)

	flag.PrintDefaults()
}

func usernameOrID(actors map[string]*gbanalytics.Actor, id string) string {
	r, ok := actors[id]
	if !ok {
		return id
	}

	return r.Username
}

func printActorsResults(rs []*gbanalytics.Result, actors map[string]*gbanalytics.Actor) {
	for i, r := range rs {
		fmt.Printf("%3d | %4d - %s\n", i+1, r.Count, usernameOrID(actors, r.ID))
	}
}

func repoNameOrID(repos map[string]*gbanalytics.Repo, id string) string {
	r, ok := repos[id]
	if !ok {
		return id
	}

	return r.Name
}

func printReposResults(rs []*gbanalytics.Result, repos map[string]*gbanalytics.Repo) {
	for i, r := range rs {
		fmt.Printf("%3d | %4d - %s\n", i+1, r.Count, repoNameOrID(repos, r.ID))
	}
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

	ctx := context.Background()

	data, err := csv.Load(ctx, csv.Config{
		ActorsFile:  path.Join(cfg.dataPath, "actors.csv"),
		CommitsFile: path.Join(cfg.dataPath, "commits.csv"),
		EventsFile:  path.Join(cfg.dataPath, "events.csv"),
		ReposFile:   path.Join(cfg.dataPath, "repos.csv"),
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, fmt.Errorf("unable to load data: %s", err))
		os.Exit(1)
	}

	switch cfg.query {
	case "top-active-users":
		results, err := gbanalytics.MostActiveUsers(data.Events, cfg.limit)
		if err != nil {
			fmt.Fprintln(os.Stderr, fmt.Errorf("unable get most active users: %s", err))
			os.Exit(1)
		}

		printActorsResults(results, data.Actors)

	case "top-active-repos":
		results, err := gbanalytics.MostActiveRepos(data.Events, cfg.limit)
		if err != nil {
			fmt.Fprintln(os.Stderr, fmt.Errorf("unable get most active repos: %s", err))
			os.Exit(1)
		}

		printReposResults(results, data.Repos)

	case "top-watch-repos":
		results, err := gbanalytics.MostWachedRepos(data.Events, cfg.limit)
		if err != nil {
			fmt.Fprintln(os.Stderr, fmt.Errorf("unable get most watched repos: %s", err))
			os.Exit(1)
		}

		printReposResults(results, data.Repos)

	default:
		fmt.Fprintln(os.Stderr, "unknown query")
		os.Exit(1)
	}
}
