package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v3"
	jwt "github.com/golang-jwt/jwt"
)

func AuthenticateJWT() fiber.Handler {
	return func(c fiber.Ctx) error {
		const BearerSchema = "Bearer "
		header := c.Get("Authorization")
		if header == "" {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Missing Authorization Header"})
		}

		if !strings.HasPrefix(header, BearerSchema) {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid authorization header"})
		}

		tokenStr := header[len(BearerSchema):]
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("there was an error")
			}
			return []byte(os.Getenv("SECRET")), nil
		})

		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"Status": "False",
				"Error":  "Error occurred while token generation",
			})
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if float64(time.Now().Unix()) > claims["exp"].(float64) {
				return c.SendStatus(http.StatusUnauthorized)
			}

			c.Locals("userid", claims["sub"])
			return c.Next()
		} else {
			return c.SendStatus(http.StatusUnauthorized)
		}
	}
}
