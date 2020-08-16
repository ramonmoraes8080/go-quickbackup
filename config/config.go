package config

import (
    "fmt"
    "io/ioutil"
    "gopkg.in/yaml.v2"
)

type ConfigurationBackend struct {
    Path string
}

type ConfigurationDefaults struct {
    Location string `yaml:",flow"`
    Schema string `yaml:",flow"`
}

type ConfigurationLocation struct {
    Backend string `yaml:",flow"`
    Path string `yaml:",flow"`
}

type Configuration struct {
    Defaults ConfigurationDefaults `yaml:",flow"`
    Backends map[interface{}]interface{} //`yaml:",flow"`
    Locations map[string]ConfigurationLocation `yaml:",flow"`
    Schemas map[string][]string
}

func (c *Configuration) Init(filePath string) {
    fmt.Println("[Configuration] reading", filePath)
    fileBytes, _ := ioutil.ReadFile(filePath)
    c.Backends = make(map[interface{}]interface{})
    err := yaml.Unmarshal(fileBytes, &c)
    if err != nil {
        panic(err)
    }
}
