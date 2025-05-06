package controllers

import (
	"github.com/gofiber/fiber/v2"
	"railway-bac/initializers"
	"railway-bac/models"
	// "railway-back/utils"
	"fmt"
	
)

// Create CBCMaterialRecord
func CreateRecord(c *fiber.Ctx) error {
	// Retrieve the user from context
	user := c.Locals("user").(models.User)

	// Parse request body into the record model
	var record models.CBCMaterialRecord
	if err := c.BodyParser(&record); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}

	// Associate the record with the user
	record.UserID = user.ID

	// Save the record to the database
	if err := initializers.DB.Create(&record).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to save record",
		})
	}

	// Respond with the created record
	return c.Status(fiber.StatusCreated).JSON(record)
}






// Get all records
func GetRecords(c *fiber.Ctx) error {
	// Retrieve the user from context
	user, ok := c.Locals("user").(models.User)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized: Invalid user context",
		})
	}

	fmt.Printf("Fetching records for user ID: %d\n", user.ID) // Debug log

	// Retrieve records associated with the user
	var records []models.CBCMaterialRecord
	err := initializers.DB.Where("user_id = ?", user.ID).Find(&records).Error
	if err != nil {
		fmt.Println("Database error:", err) // Debug log
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch records",
		})
	}

	// Check if records are empty
	if len(records) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "No records found for the user",
		})
	}

	// Respond with the retrieved records
	return c.JSON(fiber.Map{
		"status": "success",
		"data":   records,
	})
}





