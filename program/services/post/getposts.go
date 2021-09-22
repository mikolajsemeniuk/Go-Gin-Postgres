package post

import (
	"go-gin-postgres/program/database"
	postEntity "go-gin-postgres/program/entities/post"
)

const (
	GET_POSTS_QUERY = "SELECT Id, Title, Description FROM Posts"
)

func GetPosts() ([]postEntity.Post, error) {
	var posts []postEntity.Post

	rows, error := database.Client.Query(GET_POSTS_QUERY)
	if error != nil {
		return posts, error
	}
	defer rows.Close()

	for rows.Next() {
		var post postEntity.Post
		if error := rows.Scan(&post.Id, &post.Title, &post.Description); error != nil {
			return posts, error
		}
		posts = append(posts, post)
	}

	if error = rows.Err(); error != nil {
		return posts, error
	}

	return posts, nil
}
