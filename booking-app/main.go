package main

import (
	"fmt"
	"sync"
	"time"
)

var conferenceName string = "Go Conference"

const conferenceTicket int = 50

var remainingTicket uint = 50

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberofTickets uint
}

// var bookings []string  //slice
// var bookings = make([]map[string]string, 0)  //slicemap
var bookings = make([]UserData, 0) // slice of structs

var wg = sync.WaitGroup{}

func main() {

	greetings()

	// Ticket booking details

	for {
		firstName, lastName, email, numberofTickets := getUserInput()

		isvalidName, isvalidEmail, isValidTickets := infoValidation(firstName, lastName, email, numberofTickets)

		if isvalidName && isvalidEmail && isValidTickets {
			// storing the data in a slice

			bookingTicket(firstName, lastName, email, numberofTickets)
			wg.Add(1)
			go sendTicket(firstName, lastName, email, numberofTickets)

			// Making the list to show only first names

			firstNames := getFirstNames()
			fmt.Printf("First names of the bookings are %v\n", firstNames)

			// last check if tickets are available
			if remainingTicket == 0 {
				fmt.Printf("All tickets are sold.Comeback next year\n")
				break
			}
		} else {
			fmt.Printf("Invalid input. Try Again\n")
			if !isvalidName {
				fmt.Printf("Your first name or last name is too short\n")
			}
			if !isvalidEmail {
				fmt.Printf("Your email doesn't contain '@'. Check again and enter email correctly\n")
			}
			if !isValidTickets {
				fmt.Printf("Enter valid no. of tickets\n")
			}
		}

	}

	wg.Wait()
	//Output printing
	fmt.Printf("Bookings are: \n%v\n", bookings)

}

func greetings() {

	fmt.Printf("Welcome to our %v ticket booking application\n", conferenceName)
	fmt.Printf("We have %v conference ticket and only %v are left\n", conferenceTicket, remainingTicket)
	fmt.Printf("Get your tickets here\n")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {

		// // For stored data in slice
		// var names = strings.Fields(booking)
		// firstNames = append(firstNames, names[2])

		// //for saving in map
		// firstNames = append(firstNames, booking["firstName"])

		//using structs
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func bookingTicket(firstName string, lastName string, email string, numberofTickets uint) {

	// // saving Data in an slice
	// bookings = append(bookings, "\nName : "+firstName+" "+lastName+"  Tickets Booked : "+strconv.FormatUint(uint64(numberofTickets), 10)+"\n")

	// //saving data in map
	// var userData = make(map[string]string)
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["numberofTickets"] = strconv.FormatUint(uint64(numberofTickets), 10)

	// bookings = append(bookings, userData)

	//saving data in structures
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberofTickets: numberofTickets,
	}

	bookings = append(bookings, userData)
	remainingTicket = remainingTicket - numberofTickets

	fmt.Printf("Thank you %v %v .Your %v tickets are booked. You will get a confirmation email shortly in your registered email.\n", firstName, lastName, numberofTickets)
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var numberofTickets uint

	fmt.Print("Enter the your first name : ")
	fmt.Scan(&firstName)

	fmt.Print("Enter the your last name : ")
	fmt.Scan(&lastName)

	fmt.Print("Enter the your email address: ")
	fmt.Scan(&email)

	fmt.Print("Enter the number of tickets you want to book: ")
	fmt.Scan(&numberofTickets)
	return firstName, lastName, email, numberofTickets
}

func sendTicket(firstName string, lastName string, email string, numberofTickets uint) {
	time.Sleep(20 * time.Second)
	fmt.Println("*************************")
	var mail = fmt.Sprintf("%v tickets for %v %v", numberofTickets, firstName, lastName)
	fmt.Printf("Ticket:\n%v\n is sending to to email %v\n", mail, email)
	wg.Done()
}
