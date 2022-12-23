package csv

import (
	"bufio"
	"compress/gzip"
	"encoding/csv"
	"io"
	"os"
)

type Row struct {
	// Headers stores a slice of the header columns (if Reader's Header is true).
	Headers []string

	// Data is the actual row data.
	Data []string

	headermapping map[string]int
}

func (r Row) Get(header string) (string, bool) {
	if index, ok := r.headermapping[header]; ok && index < len(r.Data) {
		return r.Data[index], true
	}

	return "", false
}

type Reader struct {
	*csv.Reader

	Header bool /* if true, will extract the header row first */
}

func (r *Reader) ForEach(cb func(*Row) error) (int, error) {
	var processed int
	var headers []string

	hmap := map[string]int{}

	/* shifts the first row off and stores as the headers */
	if r.Header {
		hs, err := r.Read()

		if err != nil {
			return 0, err
		}

		headers = append(headers, hs...)

		for index := range headers {
			hmap[headers[index]] = index
		}
	}

	for {
		row, err := r.Read()

		if err != nil {
			if err == io.EOF {
				return processed, nil
			}

			return processed, err
		}

		if cb != nil {
			if err := cb(&Row{Headers: headers, Data: row, headermapping: hmap}); err != nil {
				return processed, err
			}
		}

		processed++
	}

	return processed, nil
}

func NewReader(r io.Reader) (*Reader, error) {
	reader := bufio.NewReader(r)

	bs, err := reader.Peek(2)

	if err != nil {
		return nil, err
	}

	/* sniff the first two bytes: 0x1f8b == gzip */
	if bs[0] == 0x1f && bs[1] == 0x8b {
		gr, err := gzip.NewReader(reader)

		if err != nil {
			return nil, err
		}

		r = io.Reader(gr)
	}

	return &Reader{
		Reader: csv.NewReader(r),
	}, nil
}

func NewReaderFromFile(file string) (*Reader, error) {
	f, err := os.Open(file)

	if err != nil {
		return nil, err
	}

	return NewReader(f)
}
