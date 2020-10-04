package routing

import (
	"log"
	"net/http"
)

func RunServer(setting Setting) {
	println("Running Server now")

	if setting.Config.Public {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			loggingRequest(r)
			http.ServeFile(w, r, r.URL.Path[1:])
		})
	} else {
		for _, route := range setting.Routes {
			makeHandler(route)
		}
	}

	err := http.ListenAndServe(":"+setting.Config.Port, nil)
	if err != nil {
		println(err)
	}
}

func loggingRequest(r *http.Request) {
	log.Printf("[%s] %s FROM %s\n", r.Method, r.URL.Path, r.RemoteAddr)
}

func makeHandler(route Route) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		loggingRequest(r)
		http.ServeFile(w, r, route.File)
	}
	http.HandleFunc(route.Path, fn)
}
