package routing

import (
	"log"
	"net/http"
	"os"
)

func RunServer(setting *Setting) {
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
	log.Printf("[%s] %s FROM %s; Route=%s \n", r.Method, r.URL.Path, r.RemoteAddr, route)

	if debug {
		log.Printf("%v\n", r)
	}
}

func isNotFound(route *Route, w *http.ResponseWriter, r *http.Request) bool {
	if r.URL.Path != route.Path {
		http.NotFound(*w, r)
		return true
	}
	return false
}

func isNotFileExists(filepath string) error {
	_, err := os.Stat(filepath)
	if os.IsNotExist(err) {
		return err
	}
	return nil
}

func makeHandler(route Route, setting *Setting) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		loggingRequest(r, route.Path, setting.Config.Debug)

		if isNotFound(&route, &w, r) {
			log.Printf("Not Found: %s\n", r.URL.Path)
			return
		}

		if err := isNotFileExists(route.File); err != nil {
			log.Printf("%s\n", err)
			return
		}

		if setting.Config.NoCache {
			w.Header().Set("Cache-Control", "no-cache")
		}

		http.ServeFile(w, r, route.File)
	}
	http.HandleFunc(route.Path, fn)
}
