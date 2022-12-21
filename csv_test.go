package csv

import (
	"testing"
)

func TestProcessingPlainGzippedCSV(t *testing.T) {
	config := Config{Header: true, Gzipped: true}

	processed, err := ForEachFile("testfiles/people.csv.gz", config, func(*Row) error {
		return nil /* NOP */
	})

	if err != nil {
		t.Fatalf("test failed process file; %s", err)
	}

	if processed != 4 {
		t.Fatalf("test failed process correct row count; got %d", processed)
	}
}

func TestProcessingPlainCSV(t *testing.T) {
	config := Config{Header: true}

	processed, err := ForEachFile("testfiles/people.csv", config, func(*Row) error {
		return nil /* NOP */
	})

	if err != nil {
		t.Fatalf("test failed process file; %s", err)
	}

	if processed != 4 {
		t.Fatalf("test failed process correct row count; got %d", processed)
	}
}
