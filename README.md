# CSV

[![Go Reference](https://pkg.go.dev/badge/github.com/prophittcorey/csv.svg)](https://pkg.go.dev/github.com/prophittcorey/csv)

A golang package that extends the standard library's `csv.Reader` with
`ForEach` to make stream processing large files a breeze without loading the
entire file in memory. Furthermore, this package seamlessly adds gzip support
for even easier processing at scale.

## Package Usage

```go
package main

import (
  "log"
  "strings"

  "github.com/prophittcorey/csv"
)

var mockdata = strings.NewReader(`
id,first_name,age
1,john,21
2,sam,23
`)

func main() {
  if reader, err := csv.NewReader(mockdata); err == nil {
    /* give the reader a head's up... no pun intended, kinda */
    reader.Header = true

    rows, err := reader.ForEach(func(row *csv.Row) error {
      if firstName, ok := row.Get("first_name"); ok {
              log.Printf("Hello, %s!\n", firstName)
      }

      /* stream processing halts if non-nil is returned */
      return nil
    })

    log.Printf("We processed %d rows; errors? %v\n", rows, err != nil)
  }
}
```

## License

The source code for this repository is licensed under the MIT license, which you can
find in the [LICENSE](LICENSE.md) file.
