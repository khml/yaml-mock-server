package routing

import (
	"log"
	"net/http"
)

func httpRouting() {
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
