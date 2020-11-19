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
	lettersMap := make(map[string]int)

	//Iterate through each letter of the first string and place them in a map with the key being the char and value being number of occurances
	for _, i := range firstSplit {
		_, ok := lettersMap[i]
		if ok {
			lettersMap[i] += 1
		} else {
			lettersMap[i] = 1
		}
	}
  //Iterate through the second string and lookup to see if the lettersMap contains all the the letters within the second string
	for _, i := range secondSplit {
    //If there are any that do not exist within lettersMap, return false
    if lettersMap[i] <= 0{
      fmt.Println(lettersMap)
      fmt.Println("false")
      return false
    } else {
      lettersMap[i] -= 1
    }
	}
	fmt.Println("true")
	return true
}


func main() {
	validAnagram("rat", "rar")
}
