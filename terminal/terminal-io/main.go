package main

import (
	"fmt"
	"log"

	"github.com/sethgrid/curse"
)

func main() {

	c, err := curse.New()
	if err != nil {
		log.Fatal(err)
	}

	c.SetColorBold(curse.RED).SetBackgroundColor(curse.BLACK)
	fmt.Println("Position: ", c.Position)
	c.SetDefaultStyle()
	fmt.Println("something to be erased")
	c.MoveUp(1).EraseCurrentLine().MoveDown(1)
}
