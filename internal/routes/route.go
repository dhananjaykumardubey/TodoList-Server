package routes

import (
	"fmt"
	"log"
	"net/http"
	"time"

	db "github.com/dhananjaykumardubey/TodoList-Server/internal/db"
	"github.com/dhananjaykumardubey/TodoList-Server/internal/models"
	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	fmt.Println("router default got")
	// Enable request logging
	router.Use(gin.Logger())
	apiGroup := router.Group("/todo")
	apiGroup.GET("/all", GetAllTodoHandler)
	apiGroup.POST("/create", PostAllTodoHandler)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
	fmt.Println("Server is running on http://localhost:8080")

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}
	return router
}

func GetAllTodoHandler(c *gin.Context) {
	data, err := db.GetTodo()
	if err != nil {
		// Handle error appropriately
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch todos"})
		return
	}

	// jsonResponse, err := json.Marshal(data)

	if err != nil {
		// Handle error appropriately
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch todos"})
		return
	}
	c.JSON(http.StatusOK, data)
}

func PostAllTodoHandler(c *gin.Context) {
	var todoRequest models.TodoRequestModel
	err := c.ShouldBindJSON(&todoRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	id, err := db.AddTodo(todoRequest)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create a Todo."})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("TODO created successfully at Id %d", id)})
}
