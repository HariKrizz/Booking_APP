package main

import (
	"fmt"
	"strings"
	"github.com/HariKrizz/Booking_APP"
)

const conferenceTickets int = 250

var conferenceName = "GDG India"
var remainingTickets uint = 250
var bookings = []string{}

func main() {

	greetUsers()

	for {

		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(userTickets, firstName, lastName, email)

			f_names := getFirstName()
			fmt.Printf("List of successful bookings are: %v\n", f_names)

			if remainingTickets == 0 {
				fmt.Println("Our conference tickets are sold out! Come back next year.")
				break
			}

			fmt.Println("Do you wish to continue?")
			var choice string
			fmt.Scan(&choice)

			if !(choice == "yes" || choice == "Yes" || choice == "y" || choice == "Y") {
				break
			}
		} else {
			if !isValidName {
				fmt.Println("Please enter a valid name.")
			}
			if !isValidEmail {
				fmt.Println("Please enter a valid email address.")
			}
			if !isValidTicketNumber {
				fmt.Println("Please enter a valid ticket number.")
			}
		}
	}
}

func greetUsers() {
	fmt.Println("Welcome to BookMyConference")
	fmt.Printf("Get your tickets for %v now!\n", conferenceName)
	fmt.Printf("We have total of %d tickets and %d tickets are available.\n", conferenceTickets, remainingTickets)
}

func getFirstName() []string { // input parameters and output parameters
	firstNames := []string{}

	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// User inputs
	fmt.Println("Enter your First name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your Last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your Email: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of Tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %d tickets for the conference. You will receive confirmation email at %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}


// if we need to run multiple files in go use "go run ."