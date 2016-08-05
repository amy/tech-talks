package main

import "github.com/amy/tech-talks/go-oop/methodReceiver"

func main() {

	voter = methodReceiver.Voter{
		Name:  "Amy Chen",
		Party: "Democrat",
	}

	voter.Endorse()
}
