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

/**
 * Test get lines by index
 */
func TestGetLinesByIndex(t *testing.T) {

	//generate list values
	values := []string{"a", "b", "c", "d"}

	//get value as list
	var value = getLineReaderByIndex(values, 1)

	//test value
	if value != "b" {
		t.Fatalf("Expecting comma to be a, has value: %v", value)
	}
}
