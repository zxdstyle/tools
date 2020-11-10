package main

import (
	"io/ioutil"
	"log"
	"tools/app/service/strService"
	_ "tools/routes"
)

func main()  {
	//g.Server().Run()

	bytes, err := ioutil.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	strService.DoGenModel(string(bytes[:]))
}