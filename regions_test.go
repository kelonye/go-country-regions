package regions

import (
	"testing"
)

func TestGetByNumeric(t *testing.T) {

  expectedRegion := "North America"
	region, _ := Region("US")

	if region != expectedRegion {
		t.Fatalf("Region failed: %s != %s", region, expectedRegion)
	}

}
