# Github Analytics (gbanalytics)

Get some insights about repositories and users based on the Github events


## Assumptions

1. Give that one hour of data uses ~6MB (~150MB per day), will fit in memory, I've decided to load all the data in memory before processing it.
For longer periods some optimization could be necessary, e.g. load only essencial data to be processed (excluding commit message and commit hash).

1. Pull requests and commits have the same weight while sorting.

1. It's expected that all files are in the same directory with name
`actors.csv`, `commits.csv`, `events.csv` and `repos.csv` for sake of simplicity. A next step could be pass a parameter where to locale each file, or a single zip containing all files.

## Running

You can run the main go file, build the binary before run or install before run

```
make install

gbanalytics -path data/path -query top-active-users
```

the current availiable queries are:

```
top-active-users
top-active-repos
top-watch-repos
```

To see the complete list of avaliable options:

```
gbanalytics -help
```


## Developing

### Testing

Run:

```
make test
```

### Help

For more options you can type:

```
make help
```


