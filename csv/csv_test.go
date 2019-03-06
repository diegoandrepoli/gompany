package csv

import (
	"testing"
)

/**
 * Test configuration return values
 */
func TestCommaValue(t *testing.T) {

	//test get file client name
	if comma != ';' {
		t.Fatalf("Expecting comma to be ;, has value: %v", comma)
	}
}
