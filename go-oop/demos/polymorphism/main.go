package main

import "github.com/amy/tech-talks/go-oop/polymorphism"

func main() {
	hillary := new(polymorphism.Hillary)

	hillary.Slogan()

	polymorphism.SaySlogan(hillary)
}
