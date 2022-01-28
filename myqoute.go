package myqoute

import (
	"math/rand"

	"time"

	"rsc.io/quote"
)

func GetQuote() string {
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 2
	var randNum = min + rand.Intn(max-min+1)
	newquote := "Failed from start."
	switch randNum {
	case 1:
		newquote = quote.Glass()
		newquote = newquote
	case 2:
		newquote = quote.Go()
		newquote = newquote
	case 3:
		newquote = quote.Opt()
		newquote = newquote
	case 4:
		newquote = quote.Hello()
		newquote = newquote
	default:
		newquote = "Failed to find number.."
		newquote = newquote
	}
	return newquote

}
