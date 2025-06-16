package main

import (
	"math/rand"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gorilla/securecookie"
)

type M map[string]any

const CookieName = "data"

var sc = securecookie.New([]byte("very-secret"), []byte("qfEiDYKhqQyCIKFBJmMFOrLgWQhxkzOj"))

func setCookie(c *fiber.Ctx, name string, data M) error {
	encoded, err := sc.Encode(name, data)
	if err != nil {
		return err
	}

	cookie := new(fiber.Cookie)
	cookie.Name = name
	cookie.Value = encoded
	cookie.Path = "/"
	cookie.Secure = false
	cookie.HTTPOnly = true
	cookie.Expires = time.Now().Add(1 * time.Hour)

	c.Cookie(cookie)
	return nil
}

func getCookie(c *fiber.Ctx, name string) (M, error) {
	cookie := c.Cookies(name, "")
	data := M{}
	err := sc.Decode(name, cookie, &data)
	if err == nil {
		return data, nil
	}
	return nil, err
}

func RandomString(length int) string {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		data, err := getCookie(c, CookieName)
		if err != nil && c.Cookies(CookieName) == "" && err != securecookie.ErrMacInvalid {
			return err
		}
		if data == nil {
			data = M{"Message": "Hello", "ID": RandomString(32)}

			err = setCookie(c, CookieName, data)
			if err != nil {
				return err
			}
		}
		return c.JSON(data)
	})

	log.Fatal(app.Listen(":3000"))
}
