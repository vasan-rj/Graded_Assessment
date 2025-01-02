package main


import (
	db "blogmanager/config"
	"blogmanager/controller"
	"blogmanager/middleware"
	"blogmanager/repository"
	"blogmanager/service"
	
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the database
	db.SetupDatabase()

	// Set up the repository, service, and controller for blogs
	blogRepo := repository.NewBlogStore(db.GetDatabaseConnection())
	blogService := service.NewBlogManager(blogRepo)
	blogController := controller.NewBlogController(blogService)

	// Initialize Gin router
	router := gin.Default()

	// Apply global middleware
	router.Use(middleware.LoggingMiddleware()) // Logging middleware

	// API route group with authentication middleware
	apiRoutes := router.Group("/api")
	apiRoutes.Use(middleware.AuthMiddleware(db.GetDatabaseConnection()))

	// Blog-related routes
	apiRoutes.POST("/blog", blogController.CreateBlog)      // Create a blog
	apiRoutes.GET("/blog/:id", blogController.GetBlog)      // Get a specific blog by ID
	apiRoutes.GET("/blog", blogController.GetAllBlogs)      // Get all blogs
	apiRoutes.PUT("/blog/:id", blogController.UpdateBlog)   // Update a blog by ID
	apiRoutes.DELETE("/blog/:id", blogController.DeleteBlog) // Delete a blog by ID

	// Start the server
	const port = ":8080"
	if err := router.Run(port); err != nil {
		panic("Failed to start the server: " + err.Error())
	}
}

// Run the server
// go run main.go
// curl -X GET http://localhost:8080/api/blog
// curl -X GET http://localhost:8080/api/blog/1
// curl -X GET http://localhost:8080/api/blog/2
// curl -X GET http://localhost:8080/api/blog/3
// curl -X GET http://localhost:8080/api/blog/4	

