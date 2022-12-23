# CSV

A golang package that extends the standard library's `csv` package with the
ability to iterate through a CSV (or a gzipped CSV) file without loading the
entire file in memory.

## Package Usage

```go
import "github.com/prophittcorey/csv"

csv.ForEachFile("testfiles/people.csv.gz", csv.Config{Header: true}, func(row *csv.Row) error {
  if fname, ok := row.Get("first_name"); ok {
    fmt.Printf("Hello, %s!\n", fname)
  }

  return nil
})
```

## License

The source code for this repository is licensed under the MIT license, which you can
find in the [LICENSE](LICENSE.md) file.
