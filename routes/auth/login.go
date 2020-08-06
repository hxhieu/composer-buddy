package auth

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/jwtauth"
	"github.com/go-chi/render"
	"github.com/hxhieu/composer-buddy/models"
)

// LoginRequest struct
type LoginRequest struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

// Login is validating the payload and issue the corresponding JWT
func Login(w http.ResponseWriter, r *http.Request) {
	signingKey := os.Getenv("COMPOSER_BUDDY_SIGNING_KEY")
	acceptUser := os.Getenv("COMPOSER_BUDDY_USER")
	acceptPassword := os.Getenv("COMPOSER_BUDDY_PASSWORD")
	var defaultExp int64
	defaultExp = 86400
	if exp := os.Getenv("COMPOSER_BUDDY_TOKEN_EXP"); len(exp) > 0 {
		if n, err := strconv.ParseInt(exp, 10, 64); err != nil {
			defaultExp = n
		}
	}

	var loginRequest LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
		render.Status(r, 401)
		render.JSON(w, r, models.HTTPResponse{Error: err.Error()})
		return
	}

	if loginRequest.User != acceptUser || loginRequest.Password != acceptPassword {
		render.Status(r, 401)
		render.JSON(w, r, models.HTTPResponse{Error: "Invalid user or password"})
		return
	}

	tokenAuth := jwtauth.New("HS256", []byte(signingKey), nil)
	_, jwt, _ := tokenAuth.Encode(jwt.MapClaims{
		"iss": "composer-buddy",
		"exp": time.Now().Unix() + defaultExp,
	})

	render.JSON(w, r, models.HTTPResponse{Data: jwt})
}
