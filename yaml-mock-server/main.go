package main

import (
	"flag"
	"fmt"
	"yaml-mock-server/pkg/routing"
)

func main() {
	const defaultYmlConfig = "yms.yml"

	var (
		s = flag.String("c", defaultYmlConfig, "path to Routing YAML Config")
	)
	flag.Parse()

	routes, err := routing.ReadConfig(*s)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(routes)
}
