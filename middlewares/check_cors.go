package middlewares

import (
	"net/http"
)

var allowedOrigins = []string{
	"http://localhost:5173",
	"http://localhost:3001",
	"http://127.0.0.1:5500",
}

func EnableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Izinkan akses dari origin tertentu
		// w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173 , http://127.0.0.1:5500")
		origin := r.Header.Get("Origin")

		// Cek apakah origin ada dalam daftar origin yang diizinkan
		for _, allowedOrigin := range allowedOrigins {
			if origin == allowedOrigin {
				// Jika origin sesuai, set header CORS
				w.Header().Set("Access-Control-Allow-Origin", origin)
				break
			}
		}

		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS ,HEAD")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Request-Page-Token, Page-Token")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Upgrade", "websocket")
		w.Header().Set("Connection", "Upgrade")

		// Tangani preflight request untuk CORS (OPTIONS request)
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
