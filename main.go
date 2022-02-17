package main

import (
	"fmt"

	"github.com/kyokomi/emoji/v2"
)

func main() {

	emoji.Println("Hello :world_map:!")
	emoji.Println(":beer: Beer!!!")

	pizzaMessage := emoji.Sprint("I like a :pizza: and :sushi:!!")
	fmt.Println(pizzaMessage)
}
