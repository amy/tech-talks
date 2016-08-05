package composition

import (
	"fmt"
)

type Candidate struct {
	FirstName   string
	LastName    string
	Party       string
	IsQualified bool
}

type Hillary struct {
	Candidate
}

type Trump struct {
	Candidate
}

func (c *Candidate) Name() {

	fmt.Printf("Hello! My name is %v %v\n", c.FirstName, c.LastName)

}

func (c *Candidate) Qualified() {

	if c.IsQualified == true {
		fmt.Printf("I am qualified to be POTUS.\n")
	} else {
		fmt.Printf("I am not qualified to be POTUS.\n")
	}

}
