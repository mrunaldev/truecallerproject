package main

import (
	"fmt"
	"testing"
	mypkg "truecallerproject/mypackage"
)

func TestFindLongestCommonPrefixBasic(t *testing.T) {
	myList := []string{"mruntest", "mrunal", "mruddakdjsj", "mruet3ud"}
	res := mypkg.FindLongestCommonPrefix(myList)
	if res != "mru" {
		t.Errorf("FindLongestCommonPrefix(myList) = %s; want mru", res)
	}
}

func TestFindLongestCommonPrefixTableDriven(t *testing.T) {
	input1 := []string{"mruntest", "mrunal", "mruddakdjsj", "mruet3ud"}
	input2 := []string{"xyz", "xyzmrunal", "xyzmruddakdjsj", "xyzmruet3ud"}
	input3 := []string{"gdss", "mrunfdsfal", "fds", "mruet3sdfud", "mrbbfgdgwerjejor", "mruet3sdfud"}
	input4 := []string{"wwwdfh", "wwwdfgd", "wwfff", "wgww"}
	input5 := []string{"vs3355", "55", "ghn@d", "hff$Ffg77"}
	var tests = []struct {
		inputProvided  []string
		outputExpected string
	}{
		{input1, "mru"},
		{input2, "xyz"},
		{input3, ""},
		{input4, "w"},
		{input5, ""},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("My App Test : %s", tt.inputProvided)
		t.Run(testname, func(t *testing.T) {
			res := mypkg.FindLongestCommonPrefix(tt.inputProvided)
			if res != tt.outputExpected {
				t.Errorf("got %s, want %s", res, tt.outputExpected)
			}
		})
	}
}
