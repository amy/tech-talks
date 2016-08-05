package methodReceiver

import "fmt"

// Voter information
type Voter struct {
	Name  string
	Party string
}

// Endorse candidate
func (v *Voter) Endorse() {
	fmt.Println("I'm with gop-her.")
}
