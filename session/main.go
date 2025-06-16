package main

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gorilla/context"
	"github.com/gorilla/sessions"
)

func cookieStore() *sessions.CookieStore {
	authKey := []byte("my-auth-key-very-secret")
	encryptionKey := []byte("my-encription-key-very-secret123")

	store := sessions.NewCookieStore(authKey, encryptionKey)
	store.Options.Path = "/"
	store.Options.MaxAge = 86400 * 7
	store.Options.HttpOnly = true

	return store
}

const SESSION_ID = "id"

func main() {
	store := cookieStore()

	app := fiber.New()

	app.Use(adaptor.HTTPMiddleware(context.ClearHandler))

	app.Get("/set", adaptor.HTTPHandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, SESSION_ID)
		session.Values["message1"] = "hello"
		session.Values["message2"] = "world"
		session.Save(r, w)
		http.Redirect(w, r, "/get", http.StatusTemporaryRedirect)
	}))

	app.Get("/get", func(c *fiber.Ctx) error {
		httpReq, err := adaptor.ConvertRequest(c, false)
		if err != nil {
			return err
		}
		session, _ := store.Get(httpReq, SESSION_ID)
		if len(session.Values) == 0 {
			return c.SendString("empty result")
		}
		return c.SendString(fmt.Sprintf("%s %s", session.Values["message1"], session.Values["message2"]))
	})

	app.Get("/delete", adaptor.HTTPHandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, SESSION_ID)
		session.Options.MaxAge = -1
		session.Save(r, w)
		http.Redirect(w, r, "/get", http.StatusTemporaryRedirect)
	}))

	app.Listen(":3000")
}
