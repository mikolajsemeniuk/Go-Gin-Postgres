package main

import "go-gin-postgres/program/application"

func main() {
	application.Route().Listen()
}
