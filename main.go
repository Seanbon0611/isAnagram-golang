package main

import (
	"fmt"
	"strings"
)

func validAnagram(s1 string, s2 string) bool {
	//Place a check to see if lengths of each string are the same
	if len(s1) != len(s2) {
		return false
	}
	//Split both of the strings, so we can get an array of letters that composes the string.
	firstSplit := strings.Split(s1, "")
	secondSplit := strings.Split(s2, "")

	//Initiate an empty map that we'll use to store our characters in the first array.
	lettersMap := make(map[string]int)

	//Iterate through each letter of the first string and place them in the map with the key being the character and value being the number of occurrences
	for _, i := range firstSplit {
		_, ok := lettersMap[i]
		if ok {
			lettersMap[i] += 1
		} else {
			lettersMap[i] = 1
		}
	}
	//Iterate through the second string and lookup to see if the map contains all the letters within the second string.
	for _, i := range secondSplit {
		//If there are any that do not exist within the map, return false.
		if lettersMap[i] <= 0 {
			return false
		} else {
			//Because the count of that letter is greater than 0, we're gonna reduce the count of that element by 1.
			lettersMap[i] -= 1
		}
	}
	//return true because the number of occurrences for each letter within the second string matches up with the ones in the first string.
	return true
}

func main() {
	fmt.Println(validAnagram("helloworld", "oowldlrhle")) //true
	fmt.Println(validAnagram("rat", "ta"))                //false
	fmt.Println(validAnagram("rat", "cat"))               //false
}
