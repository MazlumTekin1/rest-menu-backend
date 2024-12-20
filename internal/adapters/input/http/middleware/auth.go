package middleware

//PATH: internal/adapters/input/http/middleware/auth.go

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID   int      `json:"user_id"`
	Username string   `json:"username"`
	Roles    []string `json:"roles"`
	jwt.RegisteredClaims
}

func RequireAuth(jwtSecret string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get token from header
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(401).JSON(fiber.Map{
				"error": "authorization header required",
			})
		}

		// Parse token
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			return c.Status(401).JSON(fiber.Map{
				"error": "invalid authorization header format",
			})
		}

		// Validate token
		token, err := jwt.ParseWithClaims(tokenParts[1], &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtSecret), nil
		})

		if err != nil || !token.Valid {
			return c.Status(401).JSON(fiber.Map{
				"error": "invalid token",
			})
		}

		// Set claims to context
		if claims, ok := token.Claims.(*Claims); ok {
			c.Locals("user", claims)
		}

		return c.Next()
	}
}

func RequireRoles(roles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := c.Locals("user").(*Claims)

		hasRole := false
		for _, requiredRole := range roles {
			for _, userRole := range user.Roles {
				if requiredRole == userRole {
					hasRole = true
					break
				}
			}
		}

		if !hasRole {
			return c.Status(403).JSON(fiber.Map{
				"error": "insufficient permissions",
			})
		}

		return c.Next()
	}
}
