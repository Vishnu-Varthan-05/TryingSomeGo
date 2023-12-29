package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-bestshopBackend/database"
	"github.com/go-bestshopBackend/handler"
)

func main(){
	_, err := database.InitializeDB()
	if err != nil{
		fmt.Println("Failed to initialize the database:", err)
		return
	}	

	router := gin.Default()
	router.GET("/categories", handler.GetCategories)
	router.GET("/categoriesName", handler.GetCategoriesName)
	
	
	router.Run("localhost:8080")
	fmt.Println("Hey server is running...")
}
