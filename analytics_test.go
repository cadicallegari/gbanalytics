package gbanalytics_test

import (
	"testing"

	"github.com/cadicallegari/gbanalytics"
)

func baseEvents() []*gbanalytics.Event {
	return []*gbanalytics.Event{
		{
			ID:     "1",
			RepoID: "001",
			Type:   "CreateEvent",
		},
		{
			ID:     "2",
			RepoID: "001",
			Type:   "WatchEvent",
		},
		{
			ID:     "3",
			RepoID: "001",
			Type:   "WatchEvent",
		},
		{
			ID:     "4",
			RepoID: "001",
			Type:   "PushEvent",
		},
		{
			ID:     "5",
			RepoID: "001",
			Type:   "PullRequestEvent",
		},

		{
			ID:     "6",
			RepoID: "002",
			Type:   "WatchEvent",
		},
		{
			ID:     "7",
			RepoID: "002",
			Type:   "WatchEvent",
		},
		{
			ID:     "8",
			RepoID: "002",
			Type:   "PushEvent",
		},
		{
			ID:     "9",
			RepoID: "002",
			Type:   "PushEvent",
		},
		{
			ID:     "10",
			RepoID: "002",
			Type:   "WatchEvent",
		},

		{
			ID:     "11",
			RepoID: "003",
			Type:   "WatchEvent",
		},
		{
			ID:     "12",
			RepoID: "003",
			Type:   "PushEvent",
		},
	}
}

func baseCommitsByEvent() map[string][]*gbanalytics.Commit {
	return map[string][]*gbanalytics.Commit{
		// repo id 001
		"4": {
			{Message: "m", SHA: "s"},
			{Message: "m", SHA: "s"},
		},

		// repo id 003
		"12": {
			{Message: "a", SHA: "a"},
			{Message: "b", SHA: "b"},
			{Message: "c", SHA: "c"},
			{Message: "d", SHA: "d"},
			{Message: "e", SHA: "e"},
		},
	}
}

func Test_MostActiveRepos(t *testing.T) {
	results, err := gbanalytics.MostActiveRepos(
		baseEvents(), baseCommitsByEvent(), 2,
	)
	if err != nil {
		t.Fatalf("unexpeted error %s", err)
	}

	expected := []*gbanalytics.Result{
		{
			ID:    "003",
			Count: 5,
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

func Test_MostWachedRepos(t *testing.T) {
	results, err := gbanalytics.MostWachedRepos(baseEvents(), 2)
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
