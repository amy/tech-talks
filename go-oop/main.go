package main

import (
	"fmt"

	"github.com/amy/tech-talks/go-oop/abstraction"
)

func main() {

	fmt.Println("This works")

	hillary := new(abstraction.Hillary)

	campaign := abstraction.Campaign{
		SloganSayer: hillary,
	}

	campaign.Slogan()

}
