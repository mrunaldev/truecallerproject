package main

import (
	"log"
	mypkg "truecallerproject/mypackage"
)

const (
	inputFileName = "sample_prefixes.txt"
)

func main() {

	// Find the Longest Common Prefix from input list
	lcp := mypkg.GetLongestCommonPrefix(inputFileName)

	// Process the result. In this case we would just cehck and print it.
	if lcp != "" {
		log.Println("Longest common prefix found is : ", lcp)
	} else {
		log.Println("There was no longest common prefix found with the provided list of inputs")
	}

}
