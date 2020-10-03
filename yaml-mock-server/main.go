package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"gopkg.in/yaml.v2"
)

func main() {
	readRouting()
}

type Route struct {
	Path string
	File string
}

func readRouting() {
	buf, err := ioutil.ReadFile("./sample.yml")
	if err != nil {
		fmt.Println(err)
		return
	}

	routes, err := readRouteingFromYaml(buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(routes)
}

func readRouteingFromYaml(fileBuffer []byte) ([]Route, error) {
	defaultSize := 10
	routes := make([]Route, defaultSize)
	err := yaml.Unmarshal(fileBuffer, &routes)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return routes, nil
}

func routing() {
	println("Running Server now")

	http.HandleFunc("/", helloWorld)

	_ = http.ListenAndServe(":3000", nil)
}

func loggingRequest(r *http.Request) {
	log.Printf("[%s] %s FROM %s\n", r.Method, r.URL.Path, r.RemoteAddr)
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	loggingRequest(r)

	_, err := w.Write([]byte("Hello World"))
	if err != nil {
		println(err)
	}
}
