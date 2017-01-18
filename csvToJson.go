package main

import (
	"log"
	"os"
)

type Employee struct {
	Name string
	Age  int
	Job  string
}

func main() {
	file, err := os.Open("esha.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

}
