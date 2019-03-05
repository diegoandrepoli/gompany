
package configs

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

/**
 * File configuration name
 */
var fileconfig = "config.yaml"

/**
 * Structure of file configuration
 */
type conf struct {
	FileCatalogName string `yaml:"file.catalog.name"`
	FileClientName  string `yaml:"file.client.name"`
}

/**
 * Get config file
 * @param configuration variable
 */
func (c *conf) getConf() *conf {

	//read file configuration
	yamlFile, err := ioutil.ReadFile(fileconfig)

	if err != nil {
		log.Printf("Config file error   #%v ", err)
	}

	//content file configuraion
	err = yaml.Unmarshal(yamlFile, c)

	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}

/**
 * Get configurations
 * @return all configuration in config file
 */
func GetConfiguration() conf {
	var config conf

	//get configurations
	config.getConf()
	return config
}