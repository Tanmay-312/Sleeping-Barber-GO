package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

// variables
var seatingCapacity = 15
var arrivalRate = 100
var cutDuration = 1000 * time.Millisecond
var timeOpen = 10 * time.Second

func main() {
	// send our random number generator
	rand.Seed(time.Now().UnixNano())

	// print welcome message
	color.Yellow("Welcome to the barber shop!")
	color.Yellow("The sleeping barber problem")
	color.Yellow("---------------------------")

	// create channels if we need any
	clientChan := make(chan string, seatingCapacity)
	doneChan := make(chan bool)

	// create the barbershop
	shop := BarberShop{
		ShopCapacity:    seatingCapacity,
		HairCutDuration: cutDuration,
		NumberOfBarbers: 0,
		ClientChan:      clientChan,
		BarberDoneChan:  doneChan,
		Open:            true,
	}

	color.Green("The shop is open for the day!")

	// add barbers (though problem ask for 1 barber only, you can add as much as you wish in this code)
	shop.addBarber("Franklin")
	//shop.addBarber("Trevor")
	//shop.addBarber("Michael")
	//shop.addBarber("CJ")
	//shop.addBarber("Tommy")
	//shop.addBarber("Lucy")

	// start the barbershop as a goroutine
	shopClosing := make(chan bool)
	closed := make(chan bool)

	go func() {
		<-time.After(timeOpen)
		shopClosing <- true
		shop.closeShopForDay()
		closed <- true
	}()

	// add clients
	i := 1

	go func() {
		for {
			// get a random number with average arrival rate
			randomMilliseconds := rand.Int() % (2 * arrivalRate)

			select {
			case <-shopClosing:
				return
			case <-time.After(time.Millisecond * time.Duration(randomMilliseconds)):
				shop.addClient(fmt.Sprintf("Client #%d", i))
				i++
			}
		}
	}()

	// block until the barbershop is closed
	<-closed
}
