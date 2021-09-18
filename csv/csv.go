package csv

import "github.com/cadicallegari/gbanalytics"

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

func (*Loader) Load() (*gbanalytics.Data, error) {
	return &gbanalytics.Data{}, nil
}
