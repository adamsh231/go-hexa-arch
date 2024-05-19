package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func BasicAuthMiddleware(apiKey string) fiber.Handler {
	return basicauth.New(basicauth.Config{
		Next: nil,
		Users: map[string]string{
			"john":  "doe",
			"admin": "123456",
		},
		Realm: "",
		Authorizer: func(user, pass string) bool {
			if user == "john" && pass == "doe" {
				return true
			}
			if user == "admin" && pass == "123456" {
				return true
			}
			return false
		},
		Unauthorized: func(c *fiber.Ctx) error {
			return c.SendFile("./unauthorized.html")
		},
		ContextUsername: "_user",
		ContextPassword: "_pass",
	})
}
