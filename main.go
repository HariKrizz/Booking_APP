package main

import (
	"Booking_APP/common"
	"fmt"
)

const conferenceTickets int = 250

var conferenceName = "GDG India"
var remainingTickets uint = 250
//var bookings = make([]map[string]string, 0)  //Initialized a list of maps
var bookings = make([]userData, 0)

type userData struct  {
	firstName 			string 
	lastName 			string 
	email 				string 
	number_of_tickets 	uint
}

func main() {

	greetUsers()

	for {

		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := common.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(userTickets, firstName, lastName, email)
			sendTicket(userTickets, firstName, lastName, email)

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
		firstNames = append(firstNames, booking.firstName)
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

	var userData = userData {
		firstName : firstName,
		lastName : lastName,
		email : email,
		number_of_tickets: userTickets,
	}
	
	//Mapping user  // if we need to run multiple files in go use "go run ."
                      //keys  //values
	// var userData = make(map[string]string)
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["number_of_tickets"] = strconv.FormatUint(uint64(userTickets), 10)
	
	bookings = append(bookings, userData)
	fmt.Printf("List of bookings: %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %d tickets for the conference. You will receive confirmation email at %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	var ticket = fmt.Sprintf("%v tickets for %v %v",  userTickets, firstName, lastName)
	fmt.Println("********************")
	fmt.Printf("Sending ticket:\n %v\n To email address: %v\n", ticket, email)
	fmt.Println("********************")
}
