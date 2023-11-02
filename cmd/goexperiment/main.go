package main

import (
	"log"

	"github.com/tom360v/GoExperiment/pkg/foo"
)

func main() {
	log.Println("Hello world")

	f := foo.NewFoo("myFoo")

	f.Show()

	log.Println("bye")
}
