// Test channels
package utils

import "fmt"

/** Test channels basics
 */
func TestChannels() {
	fmt.Printf("======================\nTest channels\n")
	ch := make(chan string)
	// Create goroutine via clojure
	// goroutine runned async
	go func() {
		// Send to channel
		ch <- "Done"
		// Recieve: <-ch
	}()
	// Read / Write blocke and WAITING
	// before reading after writin
	fmt.Printf("\tRecived from channel: %v\n", <-ch)

	bufferedChannels()
}

/** Test channgels buffer
  Extended logic. To chanel we can write messages
  <= buffer length. While buffer read/write - it's blocked
  We explicitly check index/length for channel for
  prevent runtime errors!
*/
func bufferedChannels() {
	// Channel bufer = 3
	ch := make(chan string, 3)
	// Send to channel
	ch <- "Will"
	ch <- "done"
	ch <- "well"
	msg := ""
	// Recieve channel buffer
	for len(ch) > 0 {
		msg += <-ch + " "
	}
	ch <- "! ! !"
	for len(ch) > 0 {
		msg += <-ch + " [x]"
	}
	fmt.Printf("\tRecived from channel buffer: %v\n", msg)
}

func testChannelsSync() {

}
