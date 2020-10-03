package routing

import (
	"io/ioutil"
	"log"
	"net/http"
)

func RunServer(setting Setting) {
	println("Running Server now")

	for _, route := range setting.Routes {
		makeHandler(route)
	}

	_ = http.ListenAndServe(":"+setting.Config.Port, nil)
}

func loggingRequest(r *http.Request) {
	log.Printf("[%s] %s FROM %s\n", r.Method, r.URL.Path, r.RemoteAddr)
}

func makeHandler(route Route) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		loggingRequest(r)

		buf, err := ioutil.ReadFile(route.File)

		if err != nil {
			println(err)
			return
		}

		_, err = w.Write(buf)
		if err != nil {
			println(err)
		}
	}
	http.HandleFunc(route.Path, fn)
}
