package main

import (
	"tools/app/console"
	_ "tools/bootstrap"
	_ "tools/routes"
)

func main() {
	console.InitConsole()
	//bytes, err := ioutil.ReadFile("test.txt")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//strService.DoGenModel(string(bytes[:]))
}
