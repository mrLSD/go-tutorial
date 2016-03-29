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
        ch <- "Done"
    }()
    // Read / Write blocke and WAITING
    // before reading after writin
    fmt.Printf("\tRecived from channel: %v\n", <-ch)
}