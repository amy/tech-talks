package abstraction

import "fmt"

type SloganSayer interface {
	Slogan()
}

type Campaign struct {
	SloganSayer
}

// Hillary implements the SloganSayer interface
type Hillary struct{}

func (h *Hillary) Slogan() {

	fmt.Println("Stronger Together.")

}

// Trump implements the SloganSayer interface
type Trump struct{}

func (t *Trump) Slogan() {

	fmt.Println("Make America Great Again.")

}
