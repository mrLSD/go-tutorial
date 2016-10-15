package gbe

import (
	"./arrays"
	"./atomic-counters"
	"./channel-buffering"
	"./channel-directions"
	"./channel-synchronization"
	"./channels"
	"./closing-channels"
	"./closures"
	"./collection-functions"
	"./constants"
	"./defer-exec"
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
	"./mutexes"
	"./non-blocking-channel-operations"
	"./panic"
	"./pointers"
	"./range-flow"
	"./range-over-channels"
	"./rate-limiting"
	"./recursion"
	"./select_go"
	"./slices"
	"./sorting"
	"./sorting-by-functions"
	"./stateful-goroutines"
	"./structs"
	"./switch-control"
	"./tickers"
	"./timeouts"
	"./timers"
	"./values"
	"./variables"
	"./variadic-functions"
	"./worker-pools"
	"fmt"
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
	timers.Main()
	tickers.Main()
	worker_pools.Main()
	rate_limiting.Main()
	atomic_counters.Main()
	mutexes.Main()
	stateful_goroutines.Main()
	sorting.Main()
	sorting_by_functions.Main()

	panic.Main()
	defer_exec.Main()
	collection_functions.Main()
}
