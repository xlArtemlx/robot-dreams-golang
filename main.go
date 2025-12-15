package main

import (
	"fmt"
	"lesson_5/documentstore"
	"lesson_5/users"
)

func main() {
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

	userService, err := users.UserService(store)
	if err != nil {
		panic(err)
	}

	if _, err := userService.CreateUser(firstUser.ID, firstUser.Name); err != nil {
		fmt.Printf(
			"user %s (%s) created: %v\n",
			firstUser.Name,
			firstUser.ID,
			err,
		)
	}
	if _, err := userService.CreateUser(secondUser.ID, secondUser.Name); err != nil {
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

	if err := userService.DeleteUser("1"); err != nil {
		fmt.Println("delete user failed:", err)
	} else {
		fmt.Println("user 1 deleted")
	}

	_, err = userService.GetUser("1")
	fmt.Println("get 1 after delete:", err)
}
