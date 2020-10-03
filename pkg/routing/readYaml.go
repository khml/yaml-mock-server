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

func ReadRouting(filename string) ([]Route, error) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return ReadRouteingFromYaml(buf)
}

func ReadRouteingFromYaml(fileBuffer []byte) ([]Route, error) {
	defaultSize := 10
	routes := make([]Route, defaultSize)
	err := yaml.Unmarshal(fileBuffer, &routes)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return routes, nil
}
