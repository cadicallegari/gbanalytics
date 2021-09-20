package gbanalytics

import (
	"sort"
)

type Actor struct {
	ID       string
	Username string
}

type Repo struct {
	ID   string
	Name string
}

type Event struct {
	ID      string
	Type    string
	ActorID string
	RepoID  string
	Commits []*Commit
}

type Commit struct {
	EventID string
	SHA     string
	Message string
}

type Result struct {
	ID    string
	Count int
}

func rankToResults(rank map[string]int, n int) []*Result {
	results := make([]*Result, 0, len(rank))

	for k, v := range rank {
		results = append(results, &Result{ID: k, Count: v})
	}

	sort.SliceStable(results, func(i, j int) bool {
		// order desc
		return results[i].Count > results[j].Count
	})

	if n > 0 && len(results) > n {
		return results[:n]
	}

	return results
}

// MostActiveUsers active users sorted by amount of PRs created and commits pushed
func MostActiveUsers(events []*Event, n int) ([]*Result, error) {
	rank := make(map[string]int)

	for _, e := range events {
		if e.Type == "PullRequestEvent" {
			rank[e.ActorID]++
		}

		if e.Type == "PushEvent" {
			rank[e.ActorID] += len(e.Commits)
		}
	}

	return rankToResults(rank, n), nil
}

// MostActiveRepos repositories sorted by amount of commits pushed
func MostActiveRepos(events []*Event, n int) ([]*Result, error) {
	rank := make(map[string]int)

	for _, e := range events {
		if e.Type == "PushEvent" {
			rank[e.RepoID] += len(e.Commits)
		}
	}

	return rankToResults(rank, n), nil
}

// MostWachedRepos repositories sorted by amount of watch events
func MostWachedRepos(events []*Event, n int) ([]*Result, error) {
	rank := make(map[string]int)

	for _, e := range events {
		if e.Type == "WatchEvent" {
			rank[e.RepoID]++
		}
	}

	return rankToResults(rank, n), nil
}
