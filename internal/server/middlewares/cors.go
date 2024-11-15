package middlewares

import (
	"net/http"
	"strings"

	"github.com/duc-cnzj/mars/v5/internal/mlog"
)

func preflightHandler(w http.ResponseWriter, r *http.Request) {
	headers := []string{"Content-Type", "Accept", "X-Requested-With", "Authorization", "Accept-Language"}
	w.Header().Set("Access-Control-Allow-Headers", strings.Join(headers, ","))
	methods := []string{"GET", "HEAD", "POST", "PUT", "PATCH", "DELETE"}
	w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))
}

func AllowCORS(logger mlog.Logger, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			if r.Method == "OPTIONS" && r.Header.Get("Access-Control-Request-Method") != "" {
				preflightHandler(w, r)
				return
			}
		}
		h.ServeHTTP(w, r)
	})
}
