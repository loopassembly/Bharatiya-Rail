package middleware

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"railway-bac/initializers"
	"railway-bac/models"
)

func DeserializeUser(c *fiber.Ctx) error {
	authorizationHeader := c.Get("Authorization")

	// Check if the Authorization header is present and has a Bearer token
	if authorizationHeader == "" || !strings.HasPrefix(authorizationHeader, "Bearer ") {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "fail",
			"message": "Missing or invalid token",
		})
	}

	// Extract the token string
	tokenString := strings.TrimPrefix(authorizationHeader, "Bearer ")

	// Validate the token
	config, _ := initializers.LoadConfig(".")
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(config.JwtSecret), nil
	})

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "fail",
			"message": "Invalid token",
		})
	}

	// Extract claims and validate
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "fail",
			"message": "Invalid token claims",
		})
	}

	// Retrieve the user from the database
	var user models.User
	initializers.DB.First(&user, "id = ?", fmt.Sprint(claims["sub"]))
	if user.ID == 0 {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"status":  "fail",
			"message": "User not found",
		})
	}

	// Store the user in the context
	c.Locals("user", user)
	return c.Next()
}
