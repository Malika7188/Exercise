package main

import (
	"fmt"
	"os"

	//"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	args := os.Args[1:]

	
	str1 := args[0]
	str2 := args[1]


	if check(str1, str2) {
		fmt.Println(str1)
	} else {
		return
	}


	// for _, char := range str1 {
	// 	output := false 
	// 	for i, char1 := range str2 {
	// 		if char == char1 {
	// 			result += string(char1)
	// 			str2 = str2[i:]

	// 			output = true
	// 			break
	// 		}
	// 	}
	// 	if !output {
	// 		return
	// 	}
	// 	// for _, char := range result {
	// 	// 	z01.PrintRune(char)
	// 	// }
	// }
	// fmt.Println(result)


}

func check(str1, str2 string) bool {
	var i,j int
	for i < len(str1) && j < len(str2) {
		if str1[i] == str2[j] {
			i++
		}
		j++
	}
	if i == len(str1) {
		return true
	} else {
		return false
	}
}
