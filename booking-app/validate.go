package main

import "strings"

func infoValidation(firstName string, lastName string, email string, numberofTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) > 2 && len(lastName) > 2
	isvalidEmail := strings.Contains(email, "@")
	isValidTicket := numberofTickets > 0 && numberofTickets <= remainingTicket

	return isValidName, isvalidEmail, isValidTicket
}
