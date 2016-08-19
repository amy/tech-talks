package polymorphism

import "fmt"

// SloganSayer blah
type SloganSayer interface {
	Slogan()
}

// SaySlogan blah
func SaySlogan(sloganSayer SloganSayer) {
	sloganSayer.Slogan()
}

// Hillary blah
type Hillary struct{}

// Slogan blah
func (h *Hillary) Slogan() {

	fmt.Println("Stronger Together.")

}

// Trump blah
type Trump struct{}

// Slogan blah
func (t *Trump) Slogan() {

	fmt.Println("Make America Great Again.")

}
