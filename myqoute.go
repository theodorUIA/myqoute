package myqoute

import (
	"math/rand"

	"strconv"

	"rsc.io/quote"
)

func GetQuote() string {

	var randNum = rand.Intn(4-1) + 1
	newquote := "Fant ikke en ny quote"
	switch randNum {
	case 1:
		newquote := quote.Glass()
		newquote = newquote + ". Num: " + strconv.Itoa(randNum)
	case 2:
		newquote := quote.Go()
		newquote = newquote + ". Num: " + strconv.Itoa(randNum)
	case 3:
		newquote := quote.Opt()
		newquote = newquote + ". Num: " + strconv.Itoa(randNum)
	case 4:
		newquote := quote.Hello()
		newquote = newquote + ". Num: " + strconv.Itoa(randNum)
	default:
		newquote := "Fant ikke en ny quote"
		newquote = newquote + ". Num: " + strconv.Itoa(randNum)
	}
	return newquote

}
