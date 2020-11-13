package main

import (
	"fmt"
	"github.com/ChimeraCoder/gojson"
	"log"
	"strings"
)

func main() {
	json, err := gojson.ParseJson(strings.NewReader(`{"test":1}`))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(json)
}
