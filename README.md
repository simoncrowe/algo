To run unit tests:
```shell
 go test algo/internal/sorting
```

Algorithms cam be run on like so:
```shell
cat ~/Downloads/largeUF.txt | time go run ./cmd/union_find WeightedQuickUnion
```

Inputs like `largeUF.txt` can be found on [this booksite](https://algs4.cs.princeton.edu/home/).
