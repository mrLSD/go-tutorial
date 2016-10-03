package gbe

import (
	"./arrays"
	"./atomic-counters"
	"./base64-encoding"
	"./channel-buffering"
	"./channel-directions"
	"./channel-synchronization"
	"./channels"
	"./closing-channels"
	"./closures"
	"./collection-functions"
	"./constants"
	"./defer-exec"
	"./epoch"
	"./errors"
	"./for_loop"
	"./functions"
	"./goroutines"
	"./hello-world"
	"./if-else"
	"./interfaces"
	"./json"
	"./maps"
	"./methods"
	"./multiple-return-values"
	"./mutexes"
	"./non-blocking-channel-operations"
	"./number-parsing"
	"./panic"
	"./pointers"
	"./random-numbers"
	"./range-flow"
	"./range-over-channels"
	"./rate-limiting"
	"./reading-files"
	"./recursion"
	"./regular-expressions"
	"./select_go"
	"./sha1-hashes"
	"./slices"
	"./sorting"
	"./sorting-by-functions"
	"./stateful-goroutines"
	"./string-formatting"
	"./string-functions"
	"./structs"
	"./switch-control"
	"./tickers"
	"./time-fn"
	"./time-formatting-parsing"
	"./timeouts"
	"./timers"
	"./url-parsing"
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
	string_functions.Main()
	string_formatting.Main()
	regular_expressions.Main()
	json.Main()
	time_fn.Main()
	epoch.Main()
	time_formatting_parsing.Main()

	random_numbers.Main()
	number_parsing.Main()
	url_parsing.Main()
	sha1_hashes.Main()
	base64_encoding.Main()
	reading_files.Main()
}
