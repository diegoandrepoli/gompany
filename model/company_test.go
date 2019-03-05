package model

import "testing"

/**
 * Test company model
 */
func TestCompanyModel(t *testing.T) {

	//populate model
	var company = Company{ID: 1, Name: "My company", Zip: "42342", Website: "www.gcc.com"}

	//check model id
	if(company.ID != 1){
		t.Fatalf("company ID expected value 1, but got size %v", company.ID)
	}

	//check name
	if(company.Name != "My company"){
		t.Fatalf("company Name expected value 1, but got size %v", company.Name)
	}

	//check zip
	if(company.Zip != "42342"){
		t.Fatalf("company Zip expected value 1, but got size %v", company.Zip)
	}

	//check website
	if(company.Website != "www.gcc.com"){
		t.Fatalf("company Website expected value 1, but got size %v", company.Website)
	}
}

/**
 * Test get lines by array index
 */
func TestGetValueArrayByIndex(t *testing.T) {

	//generate list values
	values := []string{"a", "b", "c", "d"}

	//get value as list
	var value = getValueArrayByIndex(values, 1)

	//test value
	if value != "b" {
		t.Fatalf("Expecting comma to be a, has value: %v", value)
	}
}

/**
 * Test to create company by array index
 */
func TestCreateCompanyByArray(t *testing.T) {

	//generate list values
	values := []string{"adb", "csy", "wxd"}

	//get value as list
	var company = CreateCompanyByArray(values)

	//test company name
	if company.Name != "adb" {
		t.Fatalf("Expecting comma to adb a, has value: %v", company.Name)
	}

	//test company zip
	if company.Zip != "csy" {
		t.Fatalf("Expecting comma to csy a, has value: %v", company.Zip)
	}

	//test company website
	if company.Website != "wxd" {
		t.Fatalf("Expecting comma to wxd a, has value: %v", company.Website)
	}
}