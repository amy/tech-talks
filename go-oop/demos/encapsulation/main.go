package main

import "github.com/amy/tech-talks/go-oop/encapsulation"

func main() {

	e := encapsulation.Encapsulation{}

	e.Expose()
	e.hide()
	e.Unhide()

}
