package middlewares

import (
	"be-mklinik/f"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
)

/*
1. To display a page , Front End needed a page token
2. page token for all page , exclude request page
3. login token for all page , exclude login page and request page
*/

var excludedPathsPage []string = []string{"request-page-token", "ws-sample", "ws-chat", "ws-call", "ws-listen", "panggil-antrian", "ws-antrian", "loket"}

func CheckPageToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Check for current URL already in excluded paths
		path_split := strings.Split(r.URL.Path, "/")
		if path_split[1] != "favicon.ico" {
			for _, path := range excludedPathsPage {

				if len(path_split) > 2 && path_split[2] == path {
					next.ServeHTTP(w, r) // next handler
					return
				}
			}
		}

		w.Header().Set("Content-Type", "application/json")
		pageToken := r.Header.Get("Page-Token")
		if pageToken == "" {
			f.WriteToJsonError(w, r, "Missing Page Token")
			return
		}
		err := VerifyTokenPage(pageToken)
		if err != nil {
			f.WriteToJsonError(w, r, err.Error())
			return
		}
		next.ServeHTTP(w, r)

	})
}

func VerifyTokenPage(tokenString string) error {
	var SecretKey = []byte(os.Getenv("SECRET_KEY_PAGE"))
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}

var excludedPaths []string = []string{"login", "request-page-token", "check-version", "monitor-antrian", "monitor-antrian-x", "ws-sample", "ws-chat", "ws-call", "ws-listen", "panggil-antrian", "ws-antrian", "loket"}

func CheckLoginToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Check for current URL already in excluded paths
		for _, path := range excludedPaths {
			path_split := strings.Split(r.URL.Path, "/")
			if path_split[2] == path {
				next.ServeHTTP(w, r) // next handler
				return
			}
		}
		//selain cek token cek token page juga
		w.Header().Set("Content-Type", "application/json")
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			f.WriteToJsonError(w, r, "Missing authorization header")
			return
		}
		tokenString = tokenString[len("Bearer "):]
		err := VerifyLoginToken(tokenString)
		if err != nil {
			f.WriteToJsonError(w, r, err.Error())
			return
		}
		next.ServeHTTP(w, r)

	})
}

func VerifyLoginToken(tokenString string) error {
	var SecretKey = []byte(os.Getenv("SECRET_KEY"))
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}
