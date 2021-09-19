package gbanalytics_test

import (
	"testing"

	"github.com/cadicallegari/gbanalytics"
)

func Test_MostWachedRepos(t *testing.T) {
	events := []*gbanalytics.Event{
		{
			RepoID: "001",
			Type:   "CreateEvent",
		},
		{
			RepoID: "001",
			Type:   "WatchEvent",
		},
		{
			RepoID: "001",
			Type:   "WatchEvent",
		},
		{
			RepoID: "001",
			Type:   "PushEvent",
		},
		{
			RepoID: "001",
			Type:   "PullRequestEvent",
		},

		{
			RepoID: "002",
			Type:   "WatchEvent",
		},
		{
			RepoID: "002",
			Type:   "WatchEvent",
		},
		{
			RepoID: "002",
			Type:   "PushEvent",
		},
		{
			RepoID: "002",
			Type:   "PushEvent",
		},
		{
			RepoID: "002",
			Type:   "WatchEvent",
		},

		{
			RepoID: "003",
			Type:   "WatchEvent",
		},
	}

	results, err := gbanalytics.MostWachedRepos(events, 2)
	if err != nil {
		t.Fatalf("unexpeted error %s", err)
	}

	expected := []*gbanalytics.Result{
		{
			ID:    "002",
			Count: 3,
		},
		{
			ID:    "001",
			Count: 2,
		},
	}

	if len(results) != len(expected) {
		t.Fatalf("got list with %d elemements, expecting %d ", len(results), len(expected))
	}

	for i, e := range expected {
		if results[i].ID != e.ID {
			t.Fatalf("wrong order on element [%d] got %q, expecting %q", i, results[i].ID, e.ID)
		}

		if results[i].Count != e.Count {
			t.Fatalf("wrong count value element [%d] got %d, expecting %d", i, results[i].Count, e.Count)
		}

	}

}
