package routing

import (
	"log"
	"net/http"
	"os/exec"
	"runtime"
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

	if setting.Config.Browser.Open {
		_ = openBrowser(setting)
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

func openBrowser(setting Setting) error {
	var url string = "http://localhost:" + setting.Config.Port + setting.Config.Browser.OpenPath

	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	case "linux":
		cmd = "xdg-open"
	default:
		return nil
	}
	args = append(args, url)
	return exec.Command(cmd, args...).Start()
}
