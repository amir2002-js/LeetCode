// You are given two strings word1 and word2. Merge the strings by adding letters in alternating order
// , starting with word1. If a string is longer than the other,
// append the additional letters onto the end of the merged string

package main

import (
	"fmt"
	"strings"
)

func main() {
	var word1 string
	fmt.Print("word1 : ")
	fmt.Scanln(&word1)

	var word2 string
	print("word2 : ")
	fmt.Scanln(&word2)

	f(word1 , word2)
}

func f(word1 string, word2 string) string{
	arr1 := strings.Split(word1,"")
	arr2 := strings.Split(word2,"")

	var minLen int
	var myStrMerged string = ""

	var subArr []string
	var subStr string

	if len(arr1) > len(arr2){
		minLen = len(arr2)
		subArr = arr1[len(arr2) : len(arr1)]

	}else if len(arr1) < len(arr2){ 
		minLen = len(arr1)
		subArr = arr2[len(arr1):len(arr2)]

	}else {
		minLen = len(arr1)
		subArr = []string{""}
	}

	for i := range minLen {
		myStrMerged += arr1[i] 
		myStrMerged += arr2[i] 
	}

	subStr = strings.Join(subArr, "")

	myStrMerged += subStr

	return myStrMerged

}