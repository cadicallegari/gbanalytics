package gbanalytics

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
	Actors  []Actor
	Repos   []Repo
	Events  []Event
	Commits []Commit
}

// Top 10 active users sorted by amount of PRs created and commits pushed
func (*Data) MostActiveUsers(n int) ([]Actor, error) {
	return nil, nil
}

// Top 10 repositories sorted by amount of commits pushed
func (*Data) MostActiveRepos(n int) ([]Repo, error) {
	return nil, nil
}

// Top 10 repositories sorted by amount of watch events
func (*Data) MostWachedRepos(n int) ([]Repo, error) {
	return nil, nil
}
