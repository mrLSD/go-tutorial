package gbe

import (
	"fmt"
	"./hello-world"
	"./values"
	"./variables"
	"./constants"
	"./for_loop"
)

func Run() {
	fmt.Println("=====\nGo by Examples:")
	hello_world.Main()
	values.Main()
	variables.Main()
	constants.Main()
	for_loop.Main()
}