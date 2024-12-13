package testingExample

import (
	"context"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testingExample/database"
	"testingExample/service"
)

func main() {

	ctx := context.Background()

	g, err := gorm.Open(mysql.Open(""), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// rules of DI
	// 1. no object should be created inside the function
	// 2. all objects should be passed as arguments, those objects should be interfaces

	// Create a new database instance
	db := database.NewDatabase(g)
	userService := service.NewService(db)

	// Get all users
	users, err := userService.GetAll(ctx)
	if err != nil {
		panic(err)
	}

	for _, user := range users {
		println(user.Name)
	}

}
