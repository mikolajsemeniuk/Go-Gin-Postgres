package post

import (
	"database/sql"
	"errors"
	"go-gin-postgres/program/database"
	postEntity "go-gin-postgres/program/entities/post"
)

const (
	GET_POST_QUERY = "SELECT Id, Title, Description FROM Posts WHERE Id = $1"
)

func GetPost(id int64) (postEntity.Post, error) {
	var post postEntity.Post

	if error := database.Client.
		QueryRow(GET_POST_QUERY, id).
		Scan(&post.Id, &post.Title, &post.Description); error != nil {
		if error == sql.ErrNoRows {
			return postEntity.Post{}, errors.New("no rows :(")
		} else {
			return postEntity.Post{}, error
		}
	}
	return post, nil
}
