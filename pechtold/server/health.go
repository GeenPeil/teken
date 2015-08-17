package server

import "net/http"

func (s *Server) newHealthCheckHandlerFunc() http.HandlerFunc {
	okBytes := []byte(`ok`)

	return func(w http.ResponseWriter, r *http.Request) {
		r.Body.Close()
		w.Write(okBytes)
	}
}
