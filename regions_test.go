package regions

import (
	"testing"
)

func TestGetByNumeric(t *testing.T) {

  expectedRegion := "NORTH_AMERICA"
	region, _ := Region("US")

	if region != expectedRegion {
		t.Fatalf("Region failed: %s != %s", region, expectedRegion)
	}

}
