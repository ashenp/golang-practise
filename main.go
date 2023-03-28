package main

import (
	"fmt"
	"golang-practise/GPT01"
	"golang-practise/GPT02"
)

func main() {
	fmt.Println(GPT01.WordCount("/Users/pianweiwan/Documents/github/golang-practise/GPT01/test.txt"))
	fmt.Println(GPT02.WordReverse("/Users/pianweiwan/Documents/github/golang-practise/GPT01/test.txt"))
}
