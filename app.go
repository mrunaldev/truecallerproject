package main

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"sort"
	"sync"
	"time"
)

var (
	wg         sync.WaitGroup
	PrefixList []string
)

const (
	inputFile = "sample_prefixes.txt"
)

// Function to find the longest common prefix
func FindLongestCommonPrefix(strs []string) string {
	var longestPrefix string = ""
	var endPrefix = false

	if len(strs) > 0 {
		sort.Strings(strs)
		first := string(strs[0])
		last := string(strs[len(strs)-1])

		for i := 0; i < len(first); i++ {
			if !endPrefix && string(last[i]) == string(first[i]) {
				longestPrefix += string(last[i])
			} else {
				endPrefix = true
			}
		}
	}
	return longestPrefix
}
func processLine(line string) {
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	// log.Println("line:", line)
	PrefixList = append(PrefixList, line)
	wg.Done()
}
func main() {
	// Open input file
	reader, err := os.Open(inputFile)
	if err != nil {
		log.Panic(err)
	}

	// Close reader on exit and check for its returned error
	defer func() {
		if err := reader.Close(); err != nil {
			log.Panic(err)
		}
	}()

	// Read each line from the input config file using goroutines
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		wg.Add(1)
		go processLine(scanner.Text())
	}
	wg.Wait()

	// Find the Longest Common Prefix from input list
	lcp := FindLongestCommonPrefix(PrefixList)

	// Process the result. In this case we would just cehck and print it.
	if lcp != "" {
		log.Println("There was no longest common prefix found with the provided list of inputs")
	} else {
		log.Println("Longest common prefix found is : ", lcp)
	}
}
