package gbe

import (
	"fmt"
	"./hello-world"
	"./values"
	"./variables"
	"./constants"
)

func Run() {
	fmt.Println("=====\nGo by Examples:")
	hello_world.Main()
	values.Main()
	variables.Main()
	constants.Main()
}