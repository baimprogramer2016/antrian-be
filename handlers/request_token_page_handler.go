package handlers

import (
	"be-mklinik/f"
	"be-mklinik/responses"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

// the struct will be set contract for this handler
type pageTokenHandler struct {
	token responses.TokenResponse
}

// this function is an access handler by route
func NewPageTokenHandler() *pageTokenHandler {
	return &pageTokenHandler{}
}

func (t *pageTokenHandler) RequestPageTokenHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	defer r.Body.Close()

	page_request := r.Header.Get("Request-Page-Token")

	if page_request != os.Getenv("SECRET_USER") {
		f.WriteToJsonError(w, r, "Bad Request , Check Token Page")
		return
	}
	tokenString, err := t.CreateTokenPage()

	if err != nil {
		f.WriteToJsonError(w, r, err.Error())
		return
	}

	t.token = responses.TokenResponse{
		Token: tokenString,
	}

	f.WriteToJson(w, r, t.token)
}

func (t *pageTokenHandler) CreateTokenPage() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": os.Getenv("SECRET_USER"),
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	key := []byte(os.Getenv("SECRET_KEY_PAGE"))
	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
