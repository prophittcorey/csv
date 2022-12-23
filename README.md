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
  "bytes"
  "io"
  "log"
  "strings"

  "github.com/prophittcorey/csv"
)

// Plain text data. A regular ol' CSV file.
var plaindata = strings.NewReader(`
id,first_name,age
1,john,21
2,sam,23
`)

// Raw gzip bytes from the compressed CSV file.
var gzdata = bytes.NewReader([]byte{
  31, 139, 8, 8, 236, 213, 165, 99, 0, 3, 102, 111, 111, 46, 99, 115, 118, 0,
  203, 76, 209, 73, 203, 44, 42, 46, 137, 207, 75, 204, 77, 213, 73, 76, 79,
  229, 50, 212, 201, 202, 207, 200, 211, 49, 50, 228, 50, 210, 41, 78, 204, 213,
  49, 50, 230, 2, 0, 95, 174, 139, 212, 37, 0, 0, 0,
})

func main() {
  for _, r := range []io.Reader{plaindata, gzdata} {
      if reader, err := csv.NewReader(r); err == nil {
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
}
```

## License

The source code for this repository is licensed under the MIT license, which you can
find in the [LICENSE](LICENSE.md) file.
