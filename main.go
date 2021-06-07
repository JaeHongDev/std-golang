package main

import "fmt"

func PrintWords(words ...string) (returnWords []string) {
	returnWords = words
	return
}
func main() {

	str := PrintWords("a", "b", "c")

	for i := 0; i < len(str); i++ {
		fmt.Println(str[i])
	}

	for index, item := range str {
		fmt.Println(index, item)
	}
}
