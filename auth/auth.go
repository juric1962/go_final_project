package auth

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt"
	_ "github.com/mattn/go-sqlite3"
)

var Pass []byte

func verifyUser(token string) bool {
	jwtToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return Pass, nil
	})
	if err != nil {
		fmt.Printf("Failed to parse token: %s\n", err)
		return false
	}
	if !jwtToken.Valid {
		return false
	}
	// приводим поле Claims к типу jwt.MapClaims
	res, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Printf("failed to typecast to jwt.MapCalims")
		return false
	}
	passhashRaw := res["passhash"]
	passhash, ok := passhashRaw.(string)
	if !ok {
		fmt.Printf("failed to typecast to string")
		return false
	}
	crc := sha256.Sum256([]byte(Pass))
	hashString := hex.EncodeToString(crc[:])
	return hashString == passhash
}

// HandleApiAuthPostTesting
// сервисный обработчик для проверки ключа через Cookie
func HandleApiAuthPostTestingCookie(w http.ResponseWriter, r *http.Request) {
	var jwt string // JWT-токен из куки
	// получаем куку
	cookie, err := r.Cookie("token")
	if err == nil {
		jwt = cookie.Value
	} else {
		// возвращаем ошибку авторизации 401
		http.Error(w, "Authentification required", http.StatusUnauthorized)
		return
	}
	if verifyUser(jwt) {
		fmt.Fprintf(w, "This is your secret: Hello world\n")
		return
	}
	http.Error(w, "Unauthorized", http.StatusUnauthorized)
}

// AuthCookie
// middleware функция
//
//	( jwt = cookie.Value)
//	Со сгенерированным токеном.
func AuthCookie(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// смотрим наличие пароля
		if len(Pass) > 0 {
			var jwt string // JWT-токен из куки
			// получаем куку
			cookie, err := r.Cookie("token")
			if err == nil {
				jwt = cookie.Value
			} else {
				// возвращаем ошибку авторизации 401
				http.Error(w, "Authentification required", http.StatusUnauthorized)
				return
			}
			if !verifyUser(jwt) {
				http.Error(w, "Authentification required", http.StatusUnauthorized)
				return
			}
		}
		// возвращаем original
		next(w, r)
	})
}
