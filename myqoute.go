package myqoute

import (
	"math/rand"

	"strconv"

	"time"

	"rsc.io/quote"
)

func GetQuote() string {
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 2
	var randNum = min + rand.Intn(max-min+1)
	newquote := "Failed from start." + " Num: " + strconv.Itoa(randNum)
	switch randNum {
	case 1:
		newquote = quote.Glass()
		newquote = newquote + ". Num: " + strconv.Itoa(randNum)
	case 2:
		newquote = quote.Go()
		newquote = newquote + ". Num: " + strconv.Itoa(randNum)
	case 3:
		newquote = quote.Opt()
		newquote = newquote + ". Num: " + strconv.Itoa(randNum)
	case 4:
		newquote = quote.Hello()
		newquote = newquote + ". Num: " + strconv.Itoa(randNum)
	default:
		newquote = "Failed to find number.."
		newquote = newquote + ". Num: " + strconv.Itoa(randNum)
	}
	return newquote

}
