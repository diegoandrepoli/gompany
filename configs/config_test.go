
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
	configs.DatabaseConfigHost = "edt"
	configs.DatabaseConfigPort = "iyt"
	configs.DatabaseConfigUser = "vrw"
	configs.DatabaseConfigPassword = "spq"
	configs.DatabaseConfigDbname = "vle"
	configs.DatabaseConfigDriver = "cpp"

	//test get file client name
	if configs.FileClientName != "xyz" {
		t.Fatalf("Expecting FileClientName to be xyz, has value: %v", configs.FileClientName)
	}

	//test get file catalog name
	if configs.FileCatalogName != "wxz" {
		t.Fatalf("Expecting FileCatalogName to be wxz, company: %v", configs.FileCatalogName)
	}

	//test get database config host
	if configs.DatabaseConfigHost != "edt" {
		t.Fatalf("Expecting DatabaseConfigHost to be edt, company: %v", configs.DatabaseConfigHost)
	}

	//test get database config port
	if configs.DatabaseConfigPort != "iyt" {
		t.Fatalf("Expecting DatabaseConfigPort to be iyt, company: %v", configs.DatabaseConfigPort)
	}

	//test get database config user
	if configs.DatabaseConfigUser != "vrw" {
		t.Fatalf("Expecting DatabaseConfigUser to be vrw, company: %v", configs.DatabaseConfigUser)
	}

	//test get database config password
	if configs.DatabaseConfigPassword != "spq" {
		t.Fatalf("Expecting DatabaseConfigPassword to be spq, company: %v", configs.DatabaseConfigPassword)
	}

	//test get database config db name
	if configs.DatabaseConfigDbname != "vle" {
		t.Fatalf("Expecting FileCatalogName to be vle, company: %v", configs.DatabaseConfigDbname)
	}

	//test get database config driver
	if configs.DatabaseConfigDriver != "cpp" {
		t.Fatalf("Expecting FileCatalogName to be cpp, company: %v", configs.DatabaseConfigDriver)
	}

}