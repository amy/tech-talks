package encapsulation

import "fmt"

// Encapsulation blah
type Encapsulation struct{}

// Expose can be exported to other packages
func (e *Encapsulation) Expose() {
	fmt.Println("AHHHH! I'M EXPOSED!")
}

// hidden can only be used by functions in this package
func (e *Encapsulation) hide() {
	fmt.Println("Shhhh... this is super secret.")
}

// Unhide the private function
func (e *Encapsulation) Unhide() {

	e.hide()
	fmt.Println("Jk...")

}
