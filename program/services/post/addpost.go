package post

import (
	"errors"
	"fmt"
	"go-gin-postgres/program/database"
)

const (
	INSERT_QUERY = "INSERT INTO Posts (Title, Description) VALUES ($1, $2);"
)

func AddPost() error {
	// ONLY FOR PING DATABASE TO CHECK IF CONNECTION WAS SUCCESSFUL

	// if error := datasources.Client.Ping(); error != nil {
	// 	panic(error)
	// }
	statement, error := database.Client.Prepare(INSERT_QUERY)
	if error != nil {
		return error
	}
	defer statement.Close()

	result, error := statement.Exec("hehs", "hehss")
	if error != nil {
		return error
	}

	// result.LastInsertId()
	rows, error := result.RowsAffected()
	if error != nil {
		return error
	}

	if rows == 0 {
		return errors.New("something went wrong")
	}
	fmt.Println("success")

	return nil
}
