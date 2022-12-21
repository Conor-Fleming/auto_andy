package main

import (
	"fmt"
)

func main() {
	fmt.Println("The Morse Machine Bot v0.01")

	AddRules()
	listRules()

	//initiate stream will handle gathering tweets that satisfy rules, authentications,
	//and sending of reply tweets
	initiateStream()

	deleteRules()

}
