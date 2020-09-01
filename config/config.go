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
