package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-bestshopBackend/database"
	"github.com/go-bestshopBackend/models"
)

func SayHello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello from handler.go",
	})
}
func GetCategories(c *gin.Context) {
	allCategories, err := fetchCategories("all")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK, allCategories)
}
func GetCategoriesName(c *gin.Context) {
	allCategories, err := fetchCategories("name")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK, allCategories)
}

func fetchCategories(what string) (interface{}, error) {
	db := database.DB
	switch what {
	case "all":
		var allCategories []models.Category
		if err := db.Table("category").Find(&allCategories).Error; err != nil {
			fmt.Println("Error fetching all categories:", err)
			return nil, err
		}
		return allCategories, nil
	case "name":
		var categoryName []models.CategoryIDWithName
		if err := db.Table("category").Find(&categoryName).Error; err != nil {
			fmt.Println("Error fetching category names:", err)
			return nil, err
		}
		return categoryName, nil
	default:
		return nil, fmt.Errorf("unsupported category fetch criteria: %s", what)
	}
}

