package main

import (
	"fmt"
	"yaml-mock-server/pkg/routing"
)

func main() {
	routes, err := routing.ReadRouting("./sample.yml")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(routes)
}
