/*
Copyright © 2020 Ramon Moraes <ramonmoraes8080@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package config

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"

	"gitlab.com/velvetkeyboard/go-quickbackup/utils"
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
	Defaults ConfigurationDefaults `yaml:",flow"`
	// Backends  map[interface{}]interface{}      // `yaml:",flow"`
	Backends  map[string]interface{}           `yaml:",flow"`
	Locations map[string]ConfigurationLocation `yaml:",flow"`
	Schemas   map[string][]string
}

func (c *Configuration) Init(filePath string) {
	filePath = utils.ExpandUser(filePath)
	fileBytes, _ := ioutil.ReadFile(filePath)
	// c.Backends = make(map[interface{}]interface{})
	// c.Backends = make(map[string]interface{})
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

// We will run some verification routines on the given location name:
// - It it mapped on the config file?
// - Does the Engine associated is currentry supported?
func (c *Configuration) CheckLocationStatus(locationName string) (bool, error) {
	if c.Locations[locationName].Backend == "" {
		return false, errors.New(fmt.Sprintf(
			"backend is not defined for location \"%s\"", locationName))
	}
	if c.Locations[locationName].Path == "" {
		return false, errors.New(fmt.Sprintf(
			"path is not defined for location \"%s\"", locationName))
	}
	return true, nil
}
