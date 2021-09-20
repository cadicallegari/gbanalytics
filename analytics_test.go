package gbanalytics_test

import (
	"testing"

	"github.com/cadicallegari/gbanalytics"
)

func baseEvents() []*gbanalytics.Event {
	return []*gbanalytics.Event{
		{
			ID:      "1",
			RepoID:  "001",
			ActorID: "888",
			Type:    "CreateEvent",
		},
		{
			ID:      "2",
			RepoID:  "001",
			ActorID: "888",
			Type:    "WatchEvent",
		},
		{
			ID:      "3",
			RepoID:  "001",
			ActorID: "999",
			Type:    "WatchEvent",
		},
		{
			ID:      "4",
			RepoID:  "001",
			ActorID: "999",
			Type:    "PushEvent",
			Commits: []*gbanalytics.Commit{
				{Message: "m", SHA: "s"},
				{Message: "m", SHA: "s"},
			},
		},
		{
			ID:      "5",
			RepoID:  "001",
			ActorID: "999",
			Type:    "PullRequestEvent",
		},

		{
			ID:      "6",
			RepoID:  "002",
			ActorID: "777",
			Type:    "WatchEvent",
		},
		{
			ID:      "7",
			RepoID:  "002",
			ActorID: "999",
			Type:    "WatchEvent",
		},
		{
			ID:      "8",
			RepoID:  "002",
			ActorID: "777",
			Type:    "PushEvent",
			Commits: []*gbanalytics.Commit{
				{Message: "g", SHA: "g"},
			},
		},
		{
			ID:      "9",
			RepoID:  "002",
			ActorID: "888",
			Type:    "PushEvent",
			Commits: []*gbanalytics.Commit{
				{Message: "u", SHA: "u"},
			},
		},
		{
			ID:      "10",
			RepoID:  "002",
			ActorID: "777",
			Type:    "WatchEvent",
		},

		{
			ID:      "11",
			RepoID:  "003",
			ActorID: "888",
			Type:    "WatchEvent",
		},
		{
			ID:      "12",
			RepoID:  "003",
			ActorID: "777",
			Type:    "PushEvent",
			Commits: []*gbanalytics.Commit{
				{Message: "a", SHA: "a"},
				{Message: "b", SHA: "b"},
				{Message: "c", SHA: "c"},
				{Message: "d", SHA: "d"},
				{Message: "e", SHA: "e"},
			},
		},

		{
			ID:      "13",
			RepoID:  "001",
			ActorID: "888",
			Type:    "PullRequestEvent",
		},
	}
}

func assertResult(t *testing.T, got, want []*gbanalytics.Result) {
	t.Helper()

	if len(got) != len(want) {
		t.Fatalf("got list with %d elemements, expecting %d ", len(got), len(want))
	}

	for i, e := range want {
		if got[i].ID != e.ID {
			t.Fatalf("wrong order on element [%d] got %q, expecting %q", i, got[i].ID, e.ID)
		}

		if got[i].Count != e.Count {
			t.Fatalf("wrong count value element [%d] got %d, expecting %d", i, got[i].Count, e.Count)
		}

	}
}

func Test_MostActiveUsers(t *testing.T) {
	results, err := gbanalytics.MostActiveUsers(baseEvents(), 50)
	if err != nil {
		t.Fatalf("unexpeted error %s", err)
	}

	want := []*gbanalytics.Result{
		{
			ID:    "777",
			Count: 6,
		},
		{
			ID:    "999",
			Count: 3,
		},
		{
			ID:    "888",
			Count: 2,
		},
	}

	assertResult(t, results, want)
}

func Test_MostActiveRepos(t *testing.T) {
	results, err := gbanalytics.MostActiveRepos(baseEvents(), 2)
	if err != nil {
		t.Fatalf("unexpeted error %s", err)
	}

	want := []*gbanalytics.Result{
		{
			ID:    "003",
			Count: 5,
		},
		{
			ID:    "001",
			Count: 2,
		},
	}

	assertResult(t, results, want)
}

func Test_MostWachedRepos(t *testing.T) {
	results, err := gbanalytics.MostWachedRepos(baseEvents(), 1)
	if err != nil {
		t.Fatalf("unexpeted error %s", err)
	}

	want := []*gbanalytics.Result{
		{
			ID:    "002",
			Count: 3,
		},
	}

	assertResult(t, results, want)
}
