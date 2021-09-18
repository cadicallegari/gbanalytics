package csv_test

import (
	"context"
	"path"
	"testing"

	"github.com/cadicallegari/gbanalytics/csv"
)

func TestLoadData(t *testing.T) {
	dataPath := "../.data"

	loader := csv.NewLoader(csv.Config{
		ActorsFile:  path.Join(dataPath, "actors.csv"),
		CommitsFile: path.Join(dataPath, "commits.csv"),
		EventsFile:  path.Join(dataPath, "events.csv"),
		ReposFile:   path.Join(dataPath, "repos.csv"),
	})

	dt, err := loader.Load(context.TODO())
	if err != nil {
		t.Fatalf("not expected error loading data %s", err)
	}

	if len(dt.Actors) < 10 {
		t.Fatalf("actors not loaded as expected")
	}

	if len(dt.Commits) < 10 {
		t.Fatalf("commits not loaded as expected")
	}

	if len(dt.Events) < 10 {
		t.Fatalf("events not loaded as expected")
	}

	if len(dt.Repos) < 10 {
		t.Fatalf("repos not loaded as expected")
	}
}
