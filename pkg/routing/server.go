package routing

import (
	"log"
	"net/http"
)

func RunServer(setting Setting) {
	println("Running Server now")

	if setting.Config.Public {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			loggingRequest(r, "/", setting.Config.Debug)
			http.ServeFile(w, r, r.URL.Path[1:])
		})
	} else {
		for _, route := range setting.Routes {
			makeHandler(route, setting)
		}
	}

	if setting.Config.Browser.Open {
		_ = openBrowser(setting)
	}

	err := http.ListenAndServe(":"+setting.Config.Port, nil)
	if err != nil {
		println(err)
	}
}

func loggingRequest(r *http.Request, route string, debug bool) {
	log.Printf("Route=%s [%s] %s FROM %s\n", route, r.Method, r.URL.Path, r.RemoteAddr)

	if debug {
		log.Printf("%v", r)
	}
}

func makeHandler(route Route, setting Setting) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		loggingRequest(r, route.Path, setting.Config.Debug)
		http.ServeFile(w, r, route.File)
	}
	http.HandleFunc(route.Path, fn)
}
