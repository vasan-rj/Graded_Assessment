package controller

import (
	"blogmanager/entities"
	"blogmanager/service"
	"fmt"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

type BlogController struct {
	BlogService *service.BlogManager
}

func NewBlogController(blogService *service.BlogManager) *BlogController {
	return &BlogController{BlogService: blogService}
}

// CreateBlog handles the creation of a new blog
func (controller *BlogController) CreateBlog(ctx *gin.Context) {
	fmt.Println("CreateBlog: Received request to create a blog")

	var blog entities.BlogPost
	if err := ctx.ShouldBindJSON(&blog); err != nil {
		fmt.Println("CreateBlog: Invalid JSON:", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	fmt.Printf("CreateBlog: Parsed blog details: %+v\n", blog)

	createdBlog, err := controller.BlogService.AddBlog(&blog)
	if err != nil {
		fmt.Printf("CreateBlog: Error creating blog in service layer: %v\n", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create blog"})
		return
	}

	fmt.Printf("CreateBlog: Successfully created blog: %+v\n", createdBlog)
	ctx.JSON(http.StatusOK, createdBlog)
}

// GetBlog retrieves a blog by its ID
func (controller *BlogController) GetBlog(ctx *gin.Context) {
	id := ctx.Param("id")
	blogID, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("GetBlog: Invalid blog ID:", id)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	blog, err := controller.BlogService.FetchBlog(blogID)
	if err != nil {
		fmt.Printf("GetBlog: Blog with ID %d not found: %v\n", blogID, err)
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Blog not found"})
		return
	}

	ctx.JSON(http.StatusOK, blog)
}

// GetAllBlogs retrieves all blogs
func (controller *BlogController) GetAllBlogs(ctx *gin.Context) {
	blogs, err := controller.BlogService.FetchAllBlogs()
	if err != nil {
		fmt.Println("GetAllBlogs: Error retrieving blogs:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve blogs"})
		return
	}

	ctx.JSON(http.StatusOK, blogs)
}

// UpdateBlog updates an existing blog by its ID
func (controller *BlogController) UpdateBlog(ctx *gin.Context) {
	id := ctx.Param("id")
	blogID, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("UpdateBlog: Invalid blog ID:", id)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var blog entities.BlogPost
	if err := ctx.ShouldBindJSON(&blog); err != nil {
		fmt.Println("UpdateBlog: Invalid JSON:", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	blog.PostID = blogID
	updatedBlog, err := controller.BlogService.ModifyBlog(&blog)
	if err != nil {
		fmt.Printf("UpdateBlog: Error updating blog with ID %d: %v\n", blogID, err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update blog"})
		return
	}

	ctx.JSON(http.StatusOK, updatedBlog)
}

// DeleteBlog deletes a blog by its ID
func (controller *BlogController) DeleteBlog(ctx *gin.Context) {
	id := ctx.Param("id")
	blogID, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("DeleteBlog: Invalid blog ID:", id)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = controller.BlogService.RemoveBlog(blogID)
	if err != nil {
		fmt.Printf("DeleteBlog: Error deleting blog with ID %d: %v\n", blogID, err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete blog"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Blog deleted successfully"})
}
