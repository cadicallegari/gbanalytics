package gbanalytics

import "sort"

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
}

// sha	message	event_id
type Commit struct {
	EventID string
	SHA     string
	Message string
}

type Data struct {
	Actors  map[string]*Actor
	Repos   map[string]*Repo
	Events  []*Event
	Commits []*Commit
}

type Result struct {
	ID    string
	Count int
}

// Top 10 active users sorted by amount of PRs created and commits pushed
func MostActiveUsers(events []*Event, n int) ([]*Result, error) {
	return nil, nil
}

// Top 10 repositories sorted by amount of commits pushed
func MostActiveRepos(events []*Event, n int) ([]*Result, error) {
	return nil, nil
}

// Top 10 repositories sorted by amount of watch events
func MostWachedRepos(events []*Event, n int) ([]*Result, error) {
	rank := make(map[string]int)

	for _, e := range events {
		if e.Type == "WatchEvent" {
			rank[e.RepoID]++
		}
	}

	results := make([]*Result, 0, len(rank))

	for k, v := range rank {
		results = append(results, &Result{ID: k, Count: v})
	}

	sort.SliceStable(results, func(i, j int) bool {
		// order desc
		return results[i].Count > results[j].Count
	})

	return results[:n], nil
}
