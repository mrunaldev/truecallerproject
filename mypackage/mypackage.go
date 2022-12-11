package mypackage

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
	prefixList []string
)

// Function to contruct the list of prefixes from input file
func GetLongestCommonPrefix(inputFileName string) string {

	// Open input file
	reader, err := os.Open(inputFileName)
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
	lcp := FindLongestCommonPrefix(prefixList)
	// PrefixList = nil
	return lcp
}

// Function to find the longest common prefix
func FindLongestCommonPrefix(strList []string) string {

	var longestPrefix string = ""
	var endPrefix = false

	if len(strList) > 0 {
		sort.Strings(strList)
		first := string(strList[0])
		last := string(strList[len(strList)-1])

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

// Function to add each line to the prefix list
func processLine(line string) {

	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	// log.Println("line:", line)
	prefixList = append(prefixList, line)
	wg.Done()

}
