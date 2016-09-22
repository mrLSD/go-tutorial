package gbe

import (
	"fmt"
	"./hello-world"
	"./values"
	"./variables"
	"./constants"
	"./for_loop"
	"./if-else"
)

func Run() {
	fmt.Println("=====\nGo by Examples:")
	hello_world.Main()
	values.Main()
	variables.Main()
	constants.Main()
	for_loop.Main()
	if_else.Main()
}