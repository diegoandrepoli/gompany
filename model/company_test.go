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