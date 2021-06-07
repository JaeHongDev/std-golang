package main

import "fmt"

func PrintWords(words ...string) []string {
	return words
}
func main() {
	fmt.Print(PrintWords("a", "b", "c"))

}
