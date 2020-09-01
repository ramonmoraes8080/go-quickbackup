package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"

	"gitlab.com/velvetkeyboard/go-backup/utils"
)

type ConfigurationBackend struct {
	Path string
}

type ConfigurationDefaults struct {
	Location string `yaml:",flow"`
	Schema   string `yaml:",flow"`
}

type ConfigurationLocation struct {
	Backend string `yaml:",flow"`
	Path    string `yaml:",flow"`
}

type Configuration struct {
	Defaults  ConfigurationDefaults            `yaml:",flow"`
	Backends  map[interface{}]interface{}      // `yaml:",flow"`
	Locations map[string]ConfigurationLocation `yaml:",flow"`
	Schemas   map[string][]string
}

func (c *Configuration) Init(filePath string) {
	filePath = utils.ExpandUser(filePath)
	fileBytes, _ := ioutil.ReadFile(filePath)
	c.Backends = make(map[interface{}]interface{})
	err := yaml.Unmarshal(fileBytes, &c)
	if err != nil {
		panic(err)
	}
}

func (c *Configuration) GetDefaultLocationName() string {
	// TODO should we return error when it's empty?
	return c.Defaults.Location
}

func (c *Configuration) GetDefaultSchemaName() string {
	return c.Defaults.Schema
}
