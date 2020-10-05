package routing

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Route struct {
	Path string
	File string
}

type Config struct {
	Debug   bool
	NoCache bool `yaml:"noCache"`
	Port    string
	Public  bool
	Browser struct {
		Open     bool
		OpenPath string `yaml:"openPath"`
	}
}

type Setting struct {
	Config Config `yaml:"cfg"`
	Routes []Route
}

func ReadSetting(filename string) (Setting, error) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return Setting{}, err
	}

	return readSettingFromYaml(buf)
}

func readSettingFromYaml(fileBuffer []byte) (Setting, error) {
	setting := Setting{}
	setting.Config.NoCache = true

	err := yaml.Unmarshal(fileBuffer, &setting)
	if err != nil {
		fmt.Println(err)
		return Setting{}, err
	}
	return setting, nil
}

func LoggingSetting(c Setting) {
	fmt.Println("Config ------")
	fmt.Printf("Port = %s\n", c.Config.Port)
	fmt.Printf("Public = %t \n", c.Config.Public)
	fmt.Printf("noCache = %t \n", c.Config.NoCache)

	if c.Config.Public {
		return
	}

	fmt.Println("\nRouting ------")
	for _, route := range c.Routes {
		fmt.Printf("path = %s, file = %s\n", route.Path, route.File)
	}
	fmt.Println()
}
