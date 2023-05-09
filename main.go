package main

import (
	"fmt"
	"strings"
)

const conferenceTickets int = 50
var remainingTickets uint = 50
var conferenceName string = "Go Conference"
var bookings []string

func main() {
	greetUsers()

	for {
		firstName, lastName, email, userTickets := getUserInputs()
		isValidName, isValidEmail, isValidTickets := validateUserInputs(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTickets {
			bookTicket(firstName, lastName, userTickets, email)

			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings are: %v\n", firstNames)
		} else {
			if !isValidName {
				fmt.Println("Invalid first name!")
			}

			if !isValidEmail {
				fmt.Println("Invalid email!")
			}

			if !isValidTickets {
				fmt.Println("Invalid number of tickets!")
			}
		}
	}
}

func greetUsers() {
	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available!\n", conferenceTickets, remainingTickets)
}

func getFirstNames() []string {
	// print only first names
	firstNames := []string{}
	for _, booking := range bookings {
		name := strings.Fields(booking)[0]
		firstNames = append(firstNames, name)
	}

	return firstNames
}

func validateUserInputs(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTickets := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTickets
}

func getUserInputs() (string, string, string, uint){
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// asking for user input
	fmt.Println("Enter Your First Name: ")
	fmt.Scanln(&firstName)

	fmt.Println("Enter Your Last Name: ")
	fmt.Scanln(&lastName)

	fmt.Println("Enter Your Email: ")
	fmt.Scanln(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scanln(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(firstName string, lastName string, userTickets uint, email string){
	// book ticket in system
	remainingTickets -= userTickets
	bookings = append(bookings, firstName+" "+lastName)

	
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
