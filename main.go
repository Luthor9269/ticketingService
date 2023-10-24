package main

import (
	"fmt"
	"sync"
	"time"
)

const conferenceTickets uint = 50

var conferenceName string = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName   string
	lastName    string
	email       string
	userTickets uint
}

// creating a wait group for gorouteines
var wg = sync.WaitGroup{}

// main function is the entry point for the programme
func main() {

	greetUser()
	for remainingTickets > 0 {

		// getUserInput()
		// assign whatever user inputs into userName
		firstName, lastName, email, userTickets := getUserInput()
		// getting the validated booleans from validateuserinput
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)
		if isValidEmail && isValidName && isValidTicketNumber {
			remainingTickets, bookings = bookTicket(userTickets, firstName, lastName, email)
			wg.Add(1)
			// creating and sending a ticket
			// make this concurrent using go
			go sendTicket(userTickets, firstName, lastName, email)
			// Calling the first name function here
			fmt.Printf("Slice is %v\n", bookings)
			var firstNamesSlice []string = printFirstName()
			fmt.Printf("First names are %v\n", firstNamesSlice)

			if remainingTickets == 0 {
				fmt.Printf("Our conference is booked out. Try again next year")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First or last name entered is too short")
			}
			if !isValidEmail {
				fmt.Println("Email addres does not contain@ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets enteres is wrong")
			}
		}
	}
	wg.Wait()
}

func greetUser() {
	fmt.Printf("Welcome to %v. This conference has %v tickets in total and %v tickets remaining\n", conferenceName, conferenceTickets, remainingTickets)
}

func printFirstName() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstName := booking.firstName
		firstNames = append(firstNames, firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your First Name")
	fmt.Scan(&firstName)

	fmt.Println("Enter your Last name")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) (uint, []UserData) {
	remainingTickets = remainingTickets - userTickets
	//creating a map
	var userData = UserData{
		firstName:   firstName,
		lastName:    lastName,
		email:       email,
		userTickets: userTickets,
	}

	// passing in the userData struct into the bookings list
	bookings = append(bookings, userData)
	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("Only left %v tickets for %v \n", remainingTickets, conferenceName)
	return remainingTickets, bookings
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	// simlating delay in creating and sending ticket
	// stops the execution for 10s
	time.Sleep(10 * time.Second)
	fmt.Println("=============================")
	var ticket = fmt.Sprintf("%v tickets for %v %v\n", userTickets, firstName, lastName)
	fmt.Printf("Sending ticket: \n%v \nto email address %v\n", ticket, email)
	fmt.Println("==============================")
	wg.Done()
}
