package main

import (
	"fmt"
	"structs.com/structs-example/user"
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	//Call the constructor to create a new user
	appUser, err := user.New(userFirstName, userLastName, userBirthdate)

	if err != nil {
		fmt.Println(err)
		return
	}
	admin := user.NewAdmin("kariuki@gmail", "1234")
	admin.PrintUserDetails()
	admin.ClearUserName()
	admin.PrintUserDetails()

	appUser.PrintUserDetails()
	appUser.ClearUserName()
	appUser.PrintUserDetails()

}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
