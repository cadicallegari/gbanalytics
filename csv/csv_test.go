package csv_test

import (
	"context"
	"path"
	"testing"

	"github.com/cadicallegari/gbanalytics/csv"
)

func TestLoadData(t *testing.T) {
	dataPath := "./testdata"

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

	if len(dt.Actors) != 17 {
		t.Fatalf("loaded %d actors, want %d", len(dt.Actors), 17)
	}

	for id, a := range dt.Actors {
		if a.ID == "" || a.Username == "" {
			t.Fatalf("actor with empty values [%q] %+v", id, a)
		}
	}

	if len(dt.Repos) != 74 {
		t.Fatalf("loaded %d repos, want %d", len(dt.Repos), 74)
	}

	for id, r := range dt.Repos {
		if r.ID == "" || r.Name == "" {
			t.Fatalf("repo with empty values [%q] %+v", id, r)
		}
	}

	if len(dt.Events) != 96 {
		t.Fatalf("loaded %d events, want %d", len(dt.Events), 96)
	}

	for i, e := range dt.Events {
		if e.ID == "" || e.Type == "" || e.ActorID == "" || e.RepoID == "" {
			t.Fatalf("repo with empty values [%d] %+v", i, e)
		}

		if e.Type == "PushEvent" && len(e.Commits) == 0 {
			t.Fatalf("missing commits for push event [%d] %+v", i, e)
		}
	}
}

func TestLoadData_WrongDir(t *testing.T) {
	dataPath := "."

	loader := csv.NewLoader(csv.Config{
		ActorsFile:  path.Join(dataPath, "actors.csv"),
		CommitsFile: path.Join(dataPath, "commits.csv"),
		EventsFile:  path.Join(dataPath, "events.csv"),
		ReposFile:   path.Join(dataPath, "repos.csv"),
	})

	_, err := loader.Load(context.TODO())
	if err == nil {
		t.Fatalf("expecting error when could not load the files, got nil")
	}

}
