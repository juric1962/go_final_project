package auth

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt"
	"github.com/juric1962/go_final_project/tasks"
	_ "github.com/mattn/go-sqlite3"
)

var secret = []byte(os.Getenv("TODO_PASSWORD"))

// HandleApiAuthPost
// возвращат подписаный токен в формате json
func HandleApiAuthPost(w http.ResponseWriter, r *http.Request) {
	// получаем пароль
	var psw tasks.PSW
	var buf bytes.Buffer
	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	if err = json.Unmarshal(buf.Bytes(), &psw); err != nil {
		repErr := tasks.RepErr{Error: "ошибка десериализации JSON"}
		resJson, _ := json.Marshal(repErr)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusUnauthorized)
		out := string(resJson)
		w.Write([]byte(out))
		return
	}
	fmt.Println(psw.Password)
	if len(psw.Password) == 0 {
		resJson, _ := json.Marshal(tasks.RepErr{Error: "Authentification required"})
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusUnauthorized)
		out := string(resJson)
		w.Write([]byte(out))
		return
	}
	secret := []byte(psw.Password)
	crc := sha256.Sum256([]byte(psw.Password))
	hashString := hex.EncodeToString(crc[:])
	fmt.Println(hashString)
	// создаём payload
	claims := jwt.MapClaims{
		"passhash": hashString,
		"roles":    "qwerty",
	}
	// создаём jwt и указываем payload
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// получаем подписанный токен
	signedToken, err := jwtToken.SignedString(secret)
	if err != nil {
		fmt.Printf("failed to sign jwt: %s\n", err)
	}
	fmt.Println("Result token: " + signedToken)

	// для контроля парсим токен
	// begin control
	jwtToken, err = jwt.Parse(signedToken, func(t *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		fmt.Printf("failed to parse token: %s\n", err)
		return
	}
	// приводим поле Claims к типу jwt.MapClaims
	res, ok := jwtToken.Claims.(jwt.MapClaims)
	// обязательно используем второе возвращаемое значение ok и проверяем его, потому что
	// если Сlaims вдруг оказжется другого типа, мы получим панику
	if !ok {
		fmt.Printf("failed to typecast to jwt.MapCalims")
		return
	}
	// Так как jwt.Claims — словарь вида map[string]inteface{}, используем синтакис получения
	// занчения по ключу. Получаем значение ключа "login" и "roles"
	passhashRaw := res["passhash"]
	rolesRaw := res["roles"]
	passhash, ok := passhashRaw.(string)
	if !ok {
		fmt.Printf("failed to typecast to string")
		return
	}
	roles, ok := rolesRaw.(string)
	if !ok {
		fmt.Printf("failed to typecast to []interface")
		return
	}
	// выводим payload
	fmt.Println(passhash)
	fmt.Println(roles)
	// end control
	// возвращаем токен в формате json. {"token" : signedToken}
	resJson, _ := json.Marshal(tasks.RepJSON{Token: signedToken})
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	out := string(resJson)
	w.Write([]byte(out))
}

func verifyUser(token string) bool {
	jwtToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return secret, nil
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
	// обязательно используем второе возвращаемое значение ok и проверяем его, потому что
	// если Сlaims вдруг оказжется другого типа, мы получим панику
	if !ok {
		fmt.Printf("failed to typecast to jwt.MapCalims")
		return false
	}
	passhashRaw := res["passhash"]
	rolesRaw := res["roles"]
	passhash, ok := passhashRaw.(string)
	if !ok {
		fmt.Printf("failed to typecast to string")
		return false
	}
	roles, ok := rolesRaw.(string)
	if !ok {
		fmt.Printf("failed to typecast to []interface")
		return false
	}
	// выводим login и roles
	fmt.Printf("hash")
	fmt.Println(passhash)
	fmt.Printf("roles")
	fmt.Println(roles)
	crc := sha256.Sum256([]byte(secret))
	hashString := hex.EncodeToString(crc[:])
	return hashString == passhash
}

// HandleApiAuthPostTesting
// сервисный обработчик для проверки ключа через Bearer Token
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
// middleware функция для тестирования через браузер
//
//	( jwt = cookie.Value)
//	Со сгенерированным токеном.
func AuthCookie(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// смотрим наличие пароля
		pass := os.Getenv("TODO_PASSWORD")
		if len(pass) > 0 {
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
