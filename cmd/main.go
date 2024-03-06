package main

import (
	"fmt"

	"miras/internal/repository"
	"miras/internal/services"
	transport "miras/internal/transport/rest"
)

func main() {

	db, err := repository.NewDb()
	if err != nil {
		panic(err)
	}

	repo := repository.NewRepository(db)

	service := services.NewService(repo)
	transport := transport.NewHandler(service)
	transport.Router()
	err = transport.Gin.Run("localhost:8000")
	if err != nil {
		panic(err)
	}
	fmt.Println("Server successfuly started")
}
