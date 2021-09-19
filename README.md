# Github Analytics (gbanalytics)

Get some insights about repositories and users based on the Github events


## Assumptions

1. All the data is loaded into memory.
Given that one hour of data has ~6 MB of data,
if we start to process a bigger period of time, 24 hours, for instance, we will have ~150 MB, which is not little, but it is still an acceptable amount of data to process.
If the number of events increases and multiplies the total amount of data,
it's possible to reduce the amount of data loaded saving only the attributes needed to get the metrics,
excluding commit message and SHA for example.

1. Pull requests and commits have the same weight.

1. To keep the cli simple, we are expecting a directory having the files with the
expected names inside.
To improve this, we could enable every file to be specified as an argument,
receive a zip file or a mix of them all

## Running

You can run the main go file, build the binary before run or install before run

```
make install
gbanalytics -data ./.data -query top-active-users
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



## Develop

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


