
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
	"./slices"
	"./maps"
	"./range-flow"
	"./functions"
	"./multiple-return-values"
	"./variadic-functions"

	"./recursion"

	"./pointers"
	"./structs"
	"./methods"
	"./interfaces"

	"./errors"
	"./goroutines"
	"./channels"
	"./channel-buffering"
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
	slices.Main()
	maps.Main()

	range_flow.Main()
	functions.Main()
	multiple_return_values.Main()
	variadic_functions.Main()

	recursion.Main()

	pointers.Main()
	structs.Main()
	methods.Main()
	interfaces.Main()

	errors.Main()
	goroutines.Main()
	channels.Main()
	channel_buffering.Main()
}