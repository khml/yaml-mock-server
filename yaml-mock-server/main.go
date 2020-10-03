package main

import (
	"fmt"
	"os"
	"yaml-mock-server/pkg/routing"
)

func main() {
	const defaultYmlConfig = "yms.yml"

	pathToYml := defaultYmlConfig
	if len(os.Args) >= 2 {
		pathToYml = os.Args[1]
	}

	config, err := routing.ReadConfig(pathToYml)
	if err != nil {
		fmt.Println(err)
		return
	}
	routing.LoggingConfig(config)

	routing.RunServer(config.Cfg.Port)
}
