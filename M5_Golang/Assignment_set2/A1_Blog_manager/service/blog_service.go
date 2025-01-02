package service

import (
	entities "blogmanager/entities"
	"blogmanager/repository"
)

type BlogManager struct {
	BlogStore *repository.BlogStore
}

func NewBlogManager(blogStore *repository.BlogStore) *BlogManager {
	return &BlogManager{BlogStore: blogStore}
}

func (manager *BlogManager) AddBlog(post *entities.BlogPost) (*entities.BlogPost, error) {
	return manager.BlogStore.AddBlog(post)
}

func (manager *BlogManager) FetchBlog(postID int) (*entities.BlogPost, error) {
	return manager.BlogStore.FetchBlog(postID)
}

func (manager *BlogManager) FetchAllBlogs() ([]entities.BlogPost, error) {
	return manager.BlogStore.FetchAllBlogs()
}

func (manager *BlogManager) ModifyBlog(post *entities.BlogPost) (*entities.BlogPost, error) {
	return manager.BlogStore.ModifyBlog(post)
}

func (manager *BlogManager) RemoveBlog(postID int) error {
	return manager.BlogStore.RemoveBlog(postID)
}
