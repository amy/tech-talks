package main

import "github.com/amy/tech-talks/go-oop/abstraction"

func main() {

	hillary := new(abstraction.Hillary)
	trump := new(abstraction.Trump)

	h := abstraction.Campaign{hillary}
	t := abstraction.Campaign{trump}

	h.Slogan()
	t.Slogan()

}
