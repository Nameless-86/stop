package main

import (
	"fmt"
	"math/rand"
	"time"
)

var MAX_CHICKEN_PRICE float32 = 5
var MAX_EMPANADA_PRICE float32 = 3

func main() {
	var chickenChannel = make(chan string)
	var empanadaChannel = make(chan string)
	var websites = []string{"walmart.com", "cotsco.com", "wholefoods.com"}
	for i := range websites {
		go checkChickenPrices(websites[i], chickenChannel)
		go checkEmpanadaPrices(websites[i], empanadaChannel)
	}
	sendMessage(chickenChannel, empanadaChannel)
}

func checkEmpanadaPrices(website string, empanadaChannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		var empanadaPrice = rand.Float32() * 20
		if empanadaPrice <= MAX_CHICKEN_PRICE {
			empanadaChannel <- website
			break
		}
	}
}

func checkChickenPrices(website string, chickenChannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		var chickenPrice = rand.Float32() * 20
		if chickenPrice <= MAX_CHICKEN_PRICE {
			chickenChannel <- website
			break
		}
	}
}

func sendMessage(chickenChannel chan string, empanadaChannel chan string) {
	select {
	case website := <-chickenChannel:
		fmt.Printf("\nFound a deal at %v", website)
	case website := <-empanadaChannel:
		fmt.Printf("\nFound a deal at GO! NOW! %v", website)

	}

}
