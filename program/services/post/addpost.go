package post

import (
	"go-gin-postgres/program/database"
	"go-gin-postgres/program/inputs"
)

const (
	INSERT_POST_QUERY = "INSERT INTO Posts (Title, Description) VALUES ($1, $2);"
)

func AddPost(input inputs.PostInput) error {
	result, error := database.Client.Exec(INSERT_POST_QUERY, input.Title, input.Description)
	if error != nil {
		return error
	}

	rows, error := result.RowsAffected()
	if error != nil || rows == 0 {
		return error
	}

	return nil

	// ONLY FOR PING DATABASE TO CHECK IF CONNECTION WAS SUCCESSFUL

	// if error := datasources.Client.Ping(); error != nil {
	// 	panic(error)
	// }
	// statement, error := database.Client.Prepare(INSERT_QUERY)
	// if error != nil {
	// 	return error
	// }
	// defer statement.Close()

	// result, error := statement.Exec("hehs", "hehss")
	// if error != nil {
	// 	return error
	// }

	// // result.LastInsertId()
	// rows, error := result.RowsAffected()
	// if error != nil {
	// 	return error
	// }

	// if rows == 0 {
	// 	return errors.New("something went wrong")
	// }
	// fmt.Println("success")

	// return nil
}
