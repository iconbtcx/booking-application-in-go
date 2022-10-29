package main

import "fmt"

func main() {
	var matchName = "Wrestlamania"
	const AvailableTickets = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome to %v booking application\n", matchName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", AvailableTickets, remainingTickets)
	fmt.Println("You can book your ticket here")

	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// ask user for thier name
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email name: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, matchName)
}
