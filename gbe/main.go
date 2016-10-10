package gbe

import (
	"./arrays"
	"./channel-buffering"
	"./channel-directions"
	"./channel-synchronization"
	"./channels"
	"./closures"
	"./constants"
	"./errors"
	"./for_loop"
	"./functions"
	"./goroutines"
	"./hello-world"
	"./if-else"
	"./interfaces"
	"./maps"
	"./methods"
	"./multiple-return-values"
	"./non-blocking-channel-operations"
	"./pointers"
	"./range-flow"
	"./recursion"
	"./select_go"
	"./slices"
	"./structs"
	"./switch-control"
	"./timeouts"
	"./values"
	"./variables"
	"./variadic-functions"
	"./closing-channels"
	"fmt"
	"./range-over-channels"
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
	closures.Main()
	recursion.Main()
	pointers.Main()
	structs.Main()
	methods.Main()
	interfaces.Main()

	errors.Main()
	goroutines.Main()
	channels.Main()
	channel_buffering.Main()
	channel_synchronization.Main()
	channel_directions.Main()
	select_go.Main()
	timeouts.Main()
	non_blocking_channel_operations.Main()
	closing_channels.Main()

	range_over_channels.Main()
}
