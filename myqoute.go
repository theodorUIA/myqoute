package myqoute

import (
	"math/rand"

	"rsc.io/quote"
)

func GetQuote() string {

	var randNum = rand.Intn(4) + 1
	switch randNum {
	case 1:
		newquote := quote.Glass()
		return newquote
	case 2:
		newquote := quote.Go()
		return newquote
	case 3:
		newquote := quote.Opt()
		return newquote
	case 4:
		newquote := quote.Hello()
		return newquote
	default:
		newquote := "Fant ikke en ny quote"
		return newquote
	}

}
