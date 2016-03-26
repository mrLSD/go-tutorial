package utils

import "fmt"

/** Functions automaticaly invoke after package loaded
 */
func init() {
    
}

/** Test IF control
 */ 
func testIf() {
    fmt.Printf("================'nTest IF control")
    i := 10
    if i > 3 {
        fmt.Printf("\n\tTest if: %d\n", i)
    }
    if i < 3 {
        // ...
    } else {
        fmt.Printf("\n\tTest if-else: %d\n", i)
    }
    if j := 3; i + j < 10 {
        // ...
    } else if i + j >10 {
        fmt.Printf("\n\tTest if-else-if and pre-define values: %d\n", i)
    }
}