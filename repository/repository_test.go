package repository

import (
	"testing"
)


/**
 * Get database test
 */
func TesGetDatabase(t *testing.T) {

	if(getDatabase() != "postgress"){
		t.Fatalf("company databanse expected value postgres, but got size %v", getDatabase())
	}
}

/**
 * TODO: use https://github.com/DATA-DOG/go-sqlmock to mock batabase methods
 */