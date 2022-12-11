package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"sync"
	"time"
)

var wg sync.WaitGroup
var PrefixList []string

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
	// fmt.Println("line:", line)
	PrefixList = append(PrefixList, line)
	wg.Done()
}
func main() {
	// output := findLongestCommonPrefix([]string{"fayjn", "fayut", "fayume", "fayutom", "dslkgjkls"})

	file := "sample_prefixes.txt"
	reader, _ := os.Open(file)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		wg.Add(1)
		go processLine(scanner.Text())

	}
	wg.Wait()
	// sort.Strings(PrefixList)
	fmt.Println(FindLongestCommonPrefix(PrefixList))
}
