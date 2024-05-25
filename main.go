package revision

import (
	//"os"
	//"revision"
)

func LastRune(s string) rune {
	//args := os.Args[1:]

	// if len(args) == 0 {
	// 	return 0
	// }
	first := []rune(s)

	return first[len(s)-1]

}

