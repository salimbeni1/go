package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

var conference_name = "Go Conference"

const conference_tickets uint = 50

var remaining_tickets uint = conference_tickets
var bookings []string
var userData = make(map[string]uint)

var wg = sync.WaitGroup{}

func main() {

	greetUser()

	PrintStuff()
	PrintWierdStuff()

	for remaining_tickets > 0 {
		var userName string
		var userTickets uint
		// ask user for their name
		fmt.Printf("Enter Name\n")
		fmt.Scan(&userName)
		fmt.Printf("Enter number of tickets\n")
		fmt.Scan(&userTickets)

		var noTicketRemaining bool = userTickets > remaining_tickets
		if noTicketRemaining {
			fmt.Printf("No enough remaining tickets\n")
			continue
		}

		bookTickets(userName, userTickets)
		wg.Add(1)
		go sendTicket(userName, userTickets)
	}

	printUpperCase()

	wg.Wait()
}

func sendTicket(userName string, userTickets uint) {
	time.Sleep(10 * time.Second)
	fmt.Printf("----> %v Tickets sent for %v\n", userTickets, userName)
	wg.Done()
}

func bookTickets(userName string, userTickets uint) {
	remaining_tickets = remaining_tickets - userTickets

	bookings = append(bookings, userName)
	userData[userName] = userTickets

	fmt.Printf("User %v booking %v tickets\n", userName, userTickets)
	fmt.Printf("Remainin %v tickets\n", remaining_tickets)

}

func greetUser() {
	fmt.Printf("Welcome to our %v booking experience \n", conference_name)
	fmt.Printf("We have a remaining of %v tickets \n", remaining_tickets)
	fmt.Printf("Get you tickets here to attend\n")

}

func printUpperCase() {
	for _, booking := range bookings {

		var upper_name = strings.ToUpper(booking)
		switch upper_name {
		case "ETIENNE":
			upper_name = "<ETIENNE>"
		default:
			upper_name = "{" + upper_name + "}"
		}
		fmt.Printf("%v %v\n", upper_name, userData[booking])

	}
}
