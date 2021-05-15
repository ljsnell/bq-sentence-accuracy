package main

import (
	"fmt"

	"github.com/masatana/go-textdistance"
)

func main() {
	s1 := "this is a test"
	s2 := "that is a test"
	fmt.Println(textdistance.LevenshteinDistance(s1, s2))
	fmt.Println(textdistance.DamerauLevenshteinDistance(s1, s2))
	fmt.Println(textdistance.JaroDistance(s1, s2))
	fmt.Println(textdistance.JaroWinklerDistance(s1, s2))
}