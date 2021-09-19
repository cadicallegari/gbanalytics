package csv

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"os"

	"github.com/cadicallegari/gbanalytics"
	"golang.org/x/sync/errgroup"
)

type Config struct {
	ActorsFile  string
	CommitsFile string
	EventsFile  string
	ReposFile   string
}

type Loader struct {
	config Config
}

func NewLoader(cfg Config) *Loader {
	return &Loader{config: cfg}
}

func (l *Loader) Load(ctx context.Context) (*gbanalytics.Data, error) {
	dt := gbanalytics.Data{
		Actors: make(map[string]*gbanalytics.Actor),
		Repos:  make(map[string]*gbanalytics.Repo),
	}

	g, _ := errgroup.WithContext(ctx)

	g.Go(func() error {
		actors, err := loadActors(l.config.ActorsFile)
		if err != nil {
			return fmt.Errorf("unable to load actors: %w", err)
		}

		for _, a := range actors {
			dt.Actors[a.ID] = a
		}

		return nil
	})

	g.Go(func() error {
		commits, err := loadCommitsFile(l.config.CommitsFile)
		if err != nil {
			return fmt.Errorf("unable to load actors: %w", err)
		}

		dt.Commits = commits

		return nil
	})

	g.Go(func() error {
		events, err := loadEvents(l.config.EventsFile)
		if err != nil {
			return fmt.Errorf("unable to load actors: %w", err)
		}

		dt.Events = events

		return nil
	})

	g.Go(func() error {
		repos, err := loadRepos(l.config.ReposFile)
		if err != nil {
			return fmt.Errorf("unable to load actors: %w", err)
		}

		for _, r := range repos {
			dt.Repos[r.ID] = r
		}

		return nil
	})

	err := g.Wait()
	if err != nil {
		return nil, err
	}

	return &dt, nil
}

func loadActors(fn string) ([]*gbanalytics.Actor, error) {
	lines, err := readLines(fn)
	if err != nil {
		return nil, err
	}

	actors := make([]*gbanalytics.Actor, 0)

	for _, col := range lines {
		actors = append(actors, &gbanalytics.Actor{
			ID:       col["id"],
			Username: col["username"],
		})
	}

	return actors, nil
}

func loadCommitsFile(fn string) ([]*gbanalytics.Commit, error) {
	lines, err := readLines(fn)
	if err != nil {
		return nil, err
	}

	commits := make([]*gbanalytics.Commit, 0)

	for _, col := range lines {
		commits = append(commits, &gbanalytics.Commit{
			EventID: col["event_id"],
			SHA:     col["sha"],
			Message: col["message"],
		})
	}

	return commits, nil
}

func loadEvents(fn string) ([]*gbanalytics.Event, error) {
	lines, err := readLines(fn)
	if err != nil {
		return nil, err
	}

	events := make([]*gbanalytics.Event, 0)

	for _, col := range lines {
		events = append(events, &gbanalytics.Event{
			ID:      col["id"],
			Type:    col["type"],
			ActorID: col["actor_id"],
			RepoID:  col["repo_id"],
		})
	}

	return events, nil
}

func loadRepos(fn string) ([]*gbanalytics.Repo, error) {
	lines, err := readLines(fn)
	if err != nil {
		return nil, err
	}

	repos := make([]*gbanalytics.Repo, 0)

	for _, col := range lines {
		repos = append(repos, &gbanalytics.Repo{
			ID:   col["id"],
			Name: col["name"],
		})
	}

	return repos, nil
}

func readLines(fn string) ([]map[string]string, error) {
	r, err := os.Open(fn)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	cr := csv.NewReader(r)
	// cr.Comma = ','
	cr.TrimLeadingSpace = true
	cr.LazyQuotes = true

	lines := make([]map[string]string, 0)
	header := []string{}

	var count int
	for {
		count++

		l, err := cr.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}

		if count == 1 {
			header = l
			continue
		}

		lm := map[string]string{}
		for i, v := range l {
			lm[header[i]] = v
		}

		lines = append(lines, lm)
	}

	return lines, nil
}
