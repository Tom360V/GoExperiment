package main

import (
	"GoExperiment/pkg/foo"
	"log"
)

func main() {
	log.Println("Hello world")

	f := foo.NewFoo("myFoo")

	f.Show()

	log.Println("bye")
}
