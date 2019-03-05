
package configs

import (
	"github.com/diegoandrepoli/gompany/configs"
	"testing"
)

/**
 * Test configuration return values
 */
func TestConfigurationValues(t *testing.T) {

	//get configurations
	var configs = configs.GetConfiguration()

	//mock values
	configs.FileClientName = "xyz"
	configs.FileCatalogName = "wxz"

	//test get file client name
	if configs.FileClientName != "xyz" {
		t.Fatalf("Expecting FileClientName to be xyz, has value: %v", configs.FileClientName)
	}

	//test get file catalog name
	if configs.FileCatalogName != "wxz" {
		t.Fatalf("Expecting FileCatalogName to be wxz, company: %v", configs.FileCatalogName)
	}

}