package csv

import (
	"compress/gzip"
	"encoding/csv"
	"io"
	"log"
	"os"
)

type Config struct {
	Header      bool /* has a reader row ? */
	Gzipped     bool /* if true, will be processed as a gzipped CSV */
	ColSep      rune /* default ',' */
	LazyQuotes  bool
	ReuseRecord bool
}

type Row struct {
	Headers   []string       /* all headers */
	HeaderMap map[string]int /* maps headet to index */
	Data      []string       /* the actual row data */
}

func (r Row) Get(header string) (string, bool) {
	if index, ok := r.HeaderMap[header]; ok && index < len(r.Data) {
		return r.Data[index], true
	}

	return "", false
}

// ForEach iterates through a CSV and calls the callback for each row.
// Returns the number of rows processed and an error if any were detected.
func ForEach(reader io.Reader, config Config, cb func(*Row) error) (int, error) {
	var processed int
	var headers []string

	headerMap := map[string]int{}

	if config.ColSep == 0 {
		config.ColSep = ','
	}

	if config.Gzipped {
		gr, err := gzip.NewReader(reader)

		if err != nil {
			return 0, err
		}

		reader = io.Reader(gr)

		defer (func() {
			if err := gr.Close(); err != nil {
				log.Println(err)
			}
		})()
	}

	cr := csv.NewReader(reader)

	cr.Comma = config.ColSep
	cr.LazyQuotes = config.LazyQuotes
	cr.ReuseRecord = config.ReuseRecord

	/* pop the first row off as it's the header row */
	if config.Header {
		hs, err := cr.Read()

		if err != nil {
			return 0, err
		}

		headers = append(headers, hs...)

		for index := range headers {
			headerMap[headers[index]] = index
		}
	}

	for {
		row, err := cr.Read()

		if err != nil {
			if err == io.EOF {
				return processed, nil
			}

			continue
		}

		if cb != nil {
			err = cb(&Row{
				Headers:   headers,
				HeaderMap: headerMap,
				Data:      row,
			})

			if err != nil {
				return processed, err
			}
		}

		processed++
	}

	return processed, nil
}

// ForEachFile is a wrapper around ForEach to enable processing files with less
// boilerplate code.
func ForEachFile(file string, config Config, cb func(*Row) error) (int, error) {
	f, err := os.Open(file)

	if err != nil {
		return 0, err
	}

	defer (func() {
		if err := f.Close(); err != nil {
			log.Printf("failed to close file (%s); %s", file, err)
		}
	})()

	return ForEach(f, config, cb)
}
