package middlewares

import (
	"be-mklinik/f"
	"net/http"
)

func RoleCheckMiddleware(allowedRoles ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Misalnya, role pengguna disimpan di header Authorization (hanya contoh)
			userRole := r.Header.Get("Role")

			// Memeriksa apakah role pengguna ada di daftar allowedRoles
			for _, role := range allowedRoles {
				if userRole == role {
					next.ServeHTTP(w, r) // Lanjutkan ke handler berikutnya
					return
				}
			}

			// Jika role tidak cocok, return 403 Forbidden
			f.WriteToJsonError(w, r, "Role Not Found")
		})
	}
}
