package server

import (
	"net/http"
)

// FileServer create a file server.
func FileServer(addr, dir string) error {
	return http.ListenAndServe(addr, http.FileServer(http.Dir(dir)))
}
