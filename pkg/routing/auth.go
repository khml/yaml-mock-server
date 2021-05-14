package routing

import "net/http"

func authProxy(
	w http.ResponseWriter, r *http.Request,
	authFunc func(*http.Request) bool,
	handler func(w http.ResponseWriter, r *http.Request)) {

	if authFunc(r) {
		handler(w, r)
		return
	}

	w.Header().Add("WWW-Authenticate", `Basic realm="SECRET AREA"`)
	w.WriteHeader(http.StatusUnauthorized)
	http.Error(w, "Unauthorized", 401)
}

func makeBasicAuth(config *Config) func(*http.Request) bool {
	return func(r *http.Request) bool {
		id, secret, ok := r.BasicAuth()
		if ok == false {
			return false
		}
		return id == config.Auth.Id && secret == config.Auth.Password
	}
}

func needBasicAuth(cnf *Config) bool {
	return len(cnf.Auth.Id) > 0 && len(cnf.Auth.Password) > 0
}
