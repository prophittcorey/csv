package csv

import (
	"os"
	"testing"
)

func TestProcessingPlainGzippedCSV(t *testing.T) {
	f, err := os.Open("testfiles/people.csv.gz")

	if err != nil {
		t.Fatalf("test failed to open file; %s", err)
	}

	defer (func() {
		if err := f.Close(); err != nil {
			t.Fatalf("test failed to close file; %s", err)
		}
	})()

	config := Config{Header: true, Gzipped: true}

	processed, err := ForEach(f, config, func(*Row) error {
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
	f, err := os.Open("testfiles/people.csv")

	if err != nil {
		t.Fatalf("test failed to open file; %s", err)
	}

	defer (func() {
		if err := f.Close(); err != nil {
			t.Fatalf("test failed to close file; %s", err)
		}
	})()

	config := Config{Header: true}

	processed, err := ForEach(f, config, func(*Row) error {
		return nil /* NOP */
	})

	if err != nil {
		t.Fatalf("test failed process file; %s", err)
	}

	if processed != 4 {
		t.Fatalf("test failed process correct row count; got %d", processed)
	}
}
