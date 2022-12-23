package csv

import "testing"

func TestProcessingFiles(t *testing.T) {
	files := []string{
		"testfiles/people.csv.gz",
		"testfiles/people.csv",
	}

	for _, f := range files {
		reader, err := NewReaderFromFile(f)

		if err != nil {
			t.Fatalf("test failed open file; %s", err)
		}

		reader.Header = true

		processed, err := reader.ForEach(nil)

		if err != nil {
			t.Fatalf("test failed to process reader; %s", err)
		}

		if processed != 4 {
			t.Fatalf("test failed process correct row count; got %d", processed)
		}
	}
}

func TestForEach(t *testing.T) {
	files := []string{
		"testfiles/people.csv.gz",
		"testfiles/people.csv",
	}

	valid := map[string]bool{
		"jones": true,
		"smith": true,
	}

	for _, f := range files {
		reader, err := NewReaderFromFile(f)

		if err != nil {
			t.Fatalf("test failed open file; %s", err)
		}

		reader.Header = true

		processed, err := reader.ForEach(func(row *Row) error {
			if lname, ok := row.Get("last_name"); ok {
				if _, ok := valid[lname]; !ok {
					t.Fatalf("found invalid last name; %s", lname)
				}
			}

			return nil
		})

		if err != nil {
			t.Fatalf("test failed to process reader; %s", err)
		}

		if processed != 4 {
			t.Fatalf("test failed process correct row count; got %d", processed)
		}
	}
}
