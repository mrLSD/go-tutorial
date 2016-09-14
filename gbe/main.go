package gbe

import (
	"fmt"
	"./hello-world"
	"./values"
	"./variables"
	"./constants"
	"./for_loop"
	"./if-else"
	"./switch-control"
	"./arrays"
)

func Run() {
	fmt.Println("=====\nGo by Examples:")

	hello_world.Main()
	values.Main()
	variables.Main()
	constants.Main()
	for_loop.Main()
	if_else.Main()
	switch_control.Main()
	arrays.Main()
}