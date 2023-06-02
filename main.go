package main

import (
	"golang-practise/GPT13"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

func main() {
	//fmt.Println(GPT07.WordSorting("/Users/pianweiwan/Documents/github/golang-practise/GPT07/test.txt"))
	//fmt.Println(GPT01.WordCount("/Users/pianweiwan/Documents/github/golang-practise/GPT01/test.txt"))
	//fmt.Println(GPT02.WordReverse("/Users/pianweiwan/Documents/github/golang-practise/GPT01/test.txt"))
	//GPT05.HttpServer()
	//GPT08.ServerRun()
	//res := GPT09.CountLinesConcurrently("/Users/pianweiwan/Documents/github/golang-practise/GPT09", 4)
	//fmt.Println("res:", res)
	GPT13.Test()
}
