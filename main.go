package main

import (
	"log"

	"github.com/capybara-alt/diolibgen/diolibgen"
)

func main() {
	if err := diolibgen.Run(); err != nil {
		log.Fatal(err)
	}
}
