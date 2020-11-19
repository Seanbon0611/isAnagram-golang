package main

import (
	"fmt"
	"strings"
)

/*
Given two strings, write a fucntion to determine if the second string is an anagram of the first.

ex:
validAnagram("awesome", "awesom") //false
validAnagram("anagram", "nagaram") //true
*/

func validAnagram(s1 string, s2 string) bool {
	//Place a check to see if lengths of each string are the same
	if len(s1) != len(s2) {
		fmt.Println("false")
		return false
	}
	//Split both strings
	firstSplit := strings.Split(s1, "")
	secondSplit := strings.Split(s2, "")
	firstStringMap := make(map[string]int)
	secondStringMap := make(map[string]int)
	//Iterate through each letter of the strings and place them in a map with the key being the char and value being number of occurances
	for _, i := range firstSplit {
		_, ok := firstStringMap[i]
		if ok {
			firstStringMap[i] += 1
		} else {
			firstStringMap[i] = 1
		}
	}
	for _, i := range secondSplit {
    _, ok := secondStringMap[i]
		if ok {
			secondStringMap[i] += 1
		} else {
			secondStringMap[i] = 1
		}
	}
	//If the number of occurences for each letter do not match it is not an anagram
	for key, _ := range firstStringMap {
		_, ok := secondStringMap[key]
		if !ok {
			fmt.Println("false")
			return false
		}
	}
	fmt.Println("true")
	return true
}


func main() {
	validAnagram("rat", "tar")
}
