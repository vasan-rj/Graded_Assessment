package repository


import (
	"blogmanager/entities"
	"database/sql"
	"fmt"
	"log"
	"time"
)

type BlogStore struct {
	Connection *sql.DB
}

func NewBlogStore(connection *sql.DB) *BlogStore {
	return &BlogStore{Connection: connection}
}

func (store *BlogStore) AddBlog(post *entities.BlogPost) (*entities.BlogPost, error) {
	stmt, prepErr := store.Connection.Prepare("INSERT INTO blogs (title, content, author, timestamp) VALUES (?, ?, ?, ?)")
	if prepErr != nil {
		return nil, prepErr
	}
	defer stmt.Close()

	result, execErr := stmt.Exec(post.Heading, post.Body, post.Writer, time.Now().String())
	if execErr != nil {
		return nil, execErr
	}

	insertedID, idErr := result.LastInsertId()
	if idErr != nil {
		return nil, idErr
	}

	post.PostID = int(insertedID)
	return post, nil
}

func (store *BlogStore) FetchBlog(postID int) (*entities.BlogPost, error) {
	row := store.Connection.QueryRow("SELECT id, title, content, author, timestamp FROM blogs WHERE id = ?", postID)
	post := &entities.BlogPost{}
	scanErr := row.Scan(&post.PostID, &post.Heading, &post.Body, &post.Writer, &post.CreatedAt)
	if scanErr != nil {
		return nil, scanErr
	}
	return post, nil
}

func (store *BlogStore) FetchAllBlogs() ([]entities.BlogPost, error) {
	rows, queryErr := store.Connection.Query("SELECT id, title, content, author, timestamp FROM blogs")
	if queryErr != nil {
		return nil, queryErr
	}
	defer rows.Close()

	var posts []entities.BlogPost
	for rows.Next() {
		var post entities.BlogPost
		scanErr := rows.Scan(&post.PostID, &post.Heading, &post.Body, &post.Writer, &post.CreatedAt)
		if scanErr != nil {
			log.Fatal(scanErr)
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func (store *BlogStore) ModifyBlog(post *entities.BlogPost) (*entities.BlogPost, error) {
	stmt, prepErr := store.Connection.Prepare("UPDATE blogs SET title = ?, content = ?, author = ?, timestamp = ? WHERE id = ?")
	if prepErr != nil {
		return nil, prepErr
	}
	defer stmt.Close()

	_, execErr := stmt.Exec(post.Heading, post.Body, post.Writer, time.Now().String(), post.PostID)
	if execErr != nil {
		return nil, execErr
	}
	fmt.Println("Blog post updated successfully with ID:", post.PostID)
	return post, nil
}

func (store *BlogStore) RemoveBlog(postID int) error {
	stmt, prepErr := store.Connection.Prepare("DELETE FROM blogs WHERE id = ?")
	if prepErr != nil {
		return prepErr
	}
	defer stmt.Close()

	_, execErr := stmt.Exec(postID)
	if execErr != nil {
		return execErr
	}
	fmt.Println("Blog post deleted successfully with ID:", postID)
	return nil
}
