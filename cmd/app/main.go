package main

import (
	"context"
	"fmt"

	"github.com/xlArtemlx/robot-dreams-golang/internal/documentstore"
	"github.com/xlArtemlx/robot-dreams-golang/users"
)

func main() {
	ctx := context.Background()

	store := documentstore.NewStore()

	firstUser := struct {
		ID   string
		Name string
	}{
		ID:   "1",
		Name: "Simpson",
	}

	secondUser := struct {
		ID   string
		Name string
	}{
		ID:   "2",
		Name: "Bart",
	}

	repo, err := users.NewDocumentStoreUserRepository(ctx, store)
	if err != nil {
		panic(err)
	}

	userService := users.NewService(repo)

	if _, err := userService.CreateUser(ctx, firstUser.ID, firstUser.Name); err != nil {
		fmt.Printf(
			"user %s (%s) created: %v\n",
			firstUser.Name,
			firstUser.ID,
			err,
		)
	}
	if _, err := userService.CreateUser(ctx, secondUser.ID, secondUser.Name); err != nil {
		fmt.Printf(
			"user %s (%s) created: %v\n",
			secondUser.Name,
			secondUser.ID,
			err,
		)
	}

	list, err := userService.ListUsers()
	if err != nil {
		fmt.Println("list users failed:", err)
	} else {
		fmt.Println("users list:", list)
	}

	user, err := userService.GetUser("1")
	if err != nil {
		fmt.Println("get user failed:", err)
	} else {
		fmt.Println("user 1:", user)
	}

	if err := userService.DeleteUser(ctx, "1"); err != nil {
		fmt.Println("delete user failed:", err)
	} else {
		fmt.Println("user 1 deleted")
	}

	_, err = userService.GetUser("1")
	fmt.Println("get 1 after delete:", err)
}
