
package model

/**
 * Gompany model structure
 */
type Company struct {
	ID       int      `json:"id,omitempty"`
	Name     string   `json:"name,omitempty"`
	Zip      string   `json:"zip,omitempty"`
	Website  string   `json:"website,omitempty"`
}

/**
 * Index empty as array value not found
 */
var indexEmpty = ""

/**
 * Create company by line
 * @param list of items - csv line
 * @return company
 */
func CreateCompanyByArray(items []string) Company {
	var company Company

	//populate company
	company.Name = getValueArrayByIndex(items, 0)
	company.Zip = getValueArrayByIndex(items, 1)
	company.Website = getValueArrayByIndex(items, 2)

	return company
}

/**
 * Get line index
 * @param list of items
 * @param index of item
 * @return value as item or empty is out of range
 */
func getValueArrayByIndex(items []string, index int) string{
	if(len(items) > index){
		return items[index]
	}

	return indexEmpty
}



