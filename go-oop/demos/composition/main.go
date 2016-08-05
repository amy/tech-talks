package main

import "github.com/amy/tech-talks/go-oop/composition"

func main() {

	hillary := composition.Hillary{
		Candidate: composition.Candidate{
			FirstName:   "Hillary",
			LastName:    "Clinton",
			Party:       "Democratic",
			IsQualified: true,
		},
	}

	trump := composition.Trump{
		Candidate: composition.Candidate{
			FirstName:   "Donald",
			LastName:    "Trump",
			Party:       "Republican",
			IsQualified: false,
		},
	}

	hillary.Name()
	hillary.Qualified()

	trump.Name()
	trump.Qualified()

}
