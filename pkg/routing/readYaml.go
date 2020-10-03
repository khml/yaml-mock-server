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
	Cfg struct {
		Port string
	}
	Routes []Route
}

func ReadConfig(filename string) (Config, error) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return Config{}, err
	}

	return ReadConfigFromYaml(buf)
}

func ReadConfigFromYaml(fileBuffer []byte) (Config, error) {
	setting := Config{}
	err := yaml.Unmarshal(fileBuffer, &setting)
	if err != nil {
		fmt.Println(err)
		return Config{}, err
	}
	return setting, nil
}

func LoggingConfig(c Config) {
	fmt.Printf("Port = %s\n", c.Cfg.Port)
	for _, route := range c.Routes {
		fmt.Printf("path = %s, file = %s\n", route.Path, route.File)
	}
}
