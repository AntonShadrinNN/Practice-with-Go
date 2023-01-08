package main

import "fmt"

type Passenger struct{
	Name string
	TicketNumber int
	Boarded bool
}

type Bus struct{
	FrontSeat Passenger
}

func main(){
	casey := Passenger{"Casey", 1, false}
	fmt.Println(casey)

	var (
		bill = Passenger{Name: "Bill", TicketNumber: 2}
		ella = Passenger{Name: "Ella", TicketNumber: 3}
	)
	fmt.Printf("%s ticket number is %d\n", bill.Name, bill.TicketNumber)
	fmt.Printf("%s ticket number is %d\n", ella.Name, ella.TicketNumber)
}