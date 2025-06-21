package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

type M map[string]any
type contextKey string

const userInfoKey = contextKey("userInfo")

var (
	APPLICATION_NAME          = "JWT App"
	LOGIN_EXPIRATION_DURATION = time.Duration(1) * time.Hour
	JWT_SIGNING_METHOD        = jwt.SigningMethodHS256
	JWT_SIGNATURE_KEY         = []byte("secret")
)

func main() {
	mux := new(CustomMux)
	mux.RegisterMiddleware(MiddlewareJWTAuthorization)

	mux.HandleFunc("/index", HandlerIndex)
	mux.HandleFunc("/login", HandlerLogin)

	server := new(http.Server)
	server.Handler = mux
	server.Addr = ":3000"

	fmt.Println("Starting server at", server.Addr)
	server.ListenAndServe()
}

func HandlerIndex(w http.ResponseWriter, r *http.Request) {
	userInfo := r.Context().Value(userInfoKey).(jwt.MapClaims)
	message := fmt.Sprintf("hello %s (%s)", userInfo["Username"], userInfo["Group"])
	w.Write([]byte(message))
}

func HandlerLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	username, password, ok := r.BasicAuth()
	if !ok {
		http.Error(w, "Invalid username or password", http.StatusBadRequest)
		return
	}

	ok, userInfo := authenticateUser(username, password)
	if !ok {
		http.Error(w, "Invalid username or password", http.StatusBadRequest)
		return
	}

	type MyCustomClaims struct {
		jwt.RegisteredClaims
		Username string `json:"Username"`
		Email    string `json:"Email"`
		Group    string `json:"Group"`
	}

	claims := MyCustomClaims{
		jwt.RegisteredClaims{
			Issuer:    APPLICATION_NAME,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(LOGIN_EXPIRATION_DURATION)),
		},
		userInfo["username"].(string),
		userInfo["email"].(string),
		userInfo["group"].(string),
	}

	token := jwt.NewWithClaims(JWT_SIGNING_METHOD, claims)
	signedToken, err := token.SignedString(JWT_SIGNATURE_KEY)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	tokenString, _ := json.Marshal(M{"token": signedToken})
	w.Write([]byte(tokenString))
}

func authenticateUser(username, password string) (bool, M) {
	basePath, _ := os.Getwd()
	dbPath := filepath.Join(basePath, "users.json")
	buf, _ := os.ReadFile(dbPath)

	users := make([]M, 0)
	err := json.Unmarshal(buf, &users)
	if err != nil {
		return false, nil
	}
	var res M
	for _, user := range users {
		if user["username"] == username && user["password"] == password {
			res = user
		}
	}
	if res != nil {
		delete(res, "password")
		return true, res
	}
	return false, nil
}

func MiddlewareJWTAuthorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.URL.Path == "/login" {
			next.ServeHTTP(w, r)
			return
		}

		authorizationHeader := r.Header.Get("Authorization")
		if !strings.Contains(authorizationHeader, "Bearer") {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		tokenString := strings.ReplaceAll(authorizationHeader, "Bearer ", "")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
			return JWT_SIGNATURE_KEY, nil
		}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			http.Error(w, "Invalid Token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(context.Background(), userInfoKey, claims)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
