package csv

import (
	"testing"
)

func TestProcessingPlainGzippedCSV(t *testing.T) {
	config := Config{Header: true, Gzipped: true}

	processed, err := ForEachFile("testfiles/people.csv.gz", config, nil)

	if err != nil {
		t.Fatalf("test failed process file; %s", err)
	}

	if processed != 4 {
		t.Fatalf("test failed process correct row count; got %d", processed)
	}
}

func TestValidGetsIndexNonZero(t *testing.T) {
	config := Config{Header: true, Gzipped: true}

	valid := map[string]bool{
		"jones": true,
		"smith": true,
	}

	processed, err := ForEachFile("testfiles/people.csv.gz", config, func(row *Row) error {
		if lname, ok := row.Get("last_name"); ok {
			if _, ok := valid[lname]; !ok {
				t.Fatalf("found invalid name; %s", lname)
			}
		}

		return nil
	})

	if err != nil {
		t.Fatalf("test failed process file; %s", err)
	}

	if processed != 4 {
		t.Fatalf("test failed process correct row count; got %d", processed)
	}
}

func TestValidGetsIndexZero(t *testing.T) {
	config := Config{Header: true, Gzipped: true}

	valid := map[string]bool{
		"alex":    true,
		"alice":   true,
		"bob":     true,
		"barbara": true,
	}

	processed, err := ForEachFile("testfiles/people.csv.gz", config, func(row *Row) error {
		if fname, ok := row.Get("first_name"); ok {
			if _, ok := valid[fname]; !ok {
				t.Fatalf("found invalid name; %s", fname)
			}
		}

		return nil
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

	processed, err := ForEachFile("testfiles/people.csv", config, nil)

	if err != nil {
		t.Fatalf("test failed process file; %s", err)
	}

	if processed != 4 {
		t.Fatalf("test failed process correct row count; got %d", processed)
	}
}
