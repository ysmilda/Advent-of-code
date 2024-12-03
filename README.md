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

To run the solutions:

```
go run . -year xxxx -day xx
```

Setting the day to 0 runs all solutions for that year.