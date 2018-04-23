package pagination_test

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/ren90/paginator/pagination"
)

func TestComputePagination(t *testing.T) {
	fixtures := []struct {
		Subject  *pagination.Footer
		Expected []string
	}{}

	if err := readFromJSON("testdata", "compute_pagination.json", &fixtures); err != nil {
		t.Fatalf("Error unmarshalling fixtures: %v", err)
	}

	for i, f := range fixtures {
		t.Logf("Testing ComputePagination() case %d", i)

		actual := f.Subject.ComputePagination()
		expected := f.Expected
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("Unmatched expectations. Expected: %s | Actual: %s", expected, actual)
		}
	}
}

func readFromJSON(path, jsonFilename string, target interface{}) error {
	tableBytes, err := ioutil.ReadFile(filepath.Join(path, jsonFilename))
	if err != nil {
		return err
	}
	if err := json.Unmarshal(tableBytes, target); err != nil {
		return err
	}
	return nil
}
