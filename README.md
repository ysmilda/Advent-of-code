# Advent-of-code

These are my solutions for the [Advent of code](https://adventofcode.com/) puzzles.

To generate the next day run:

```
go run ./cmd/generate/main.go -year xxxx -day xx
```

Puzzle inputs can be fetched with:

```
go run ./cmd/fetch/main.go -session-token xxx
```

This fetches the input for all years and days defined in the `solutions` directory. If the input already exists it will not be fetched again. These inputs won't be commited to the repository as per the [Advent of code faq](https://adventofcode.com/about).

To include the new solution in the main file run:

```
go run ./cmd/update/main.go
```

This updates the days.go file to include all solutions defined in the `solutions` directory.

To run the solutions:

```
go run . -year xxxx -day xx
```

Setting the day to 0 runs all solutions for that year.