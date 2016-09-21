package utils

import "fmt"

/** Functions automaticaly invoke after package loaded
 */
func init() {
    testIf()
    testSwitch()
}

/** Test IF control
 */
func testIf() {
    fmt.Printf("\n================\nTest IF control\n")
    i := 10
    if i > 3 {
        fmt.Printf("\tTest if: %d\n", i)
    }
    if i < 3 {
        // ...
    } else {
        fmt.Printf("\tTest if-else: %d\n", i)
    }
    if j := 3; i + j < 10 {
        // ...
    } else if i + j > 10 {
        fmt.Printf("\tTest if-else-if and pre-define values: %d %d\n", i, j)
    }
    // !!! afte if block - j variable not available
}

/** Test switch control
 */
func testSwitch() {
    fmt.Printf("\n================\nTest Switch control\n")
    var i int = 10
    switch i {
        case 1, 5, 10, 20:
            fmt.Printf("\tCase for 1, 5, 10, 20: %d\n", i)
    }
    switch i {
        case 1, 4, 5:
            // ...
        default:
            fmt.Printf("\tdefault case for: %d\n", i)
    }
    switch {
        case i > 0 && 20 > i:
            fmt.Printf("\tcase for i > 0 && 20 < i: %d\n", i)
            fallthrough
        case i > 5 && 15 > i:
            fmt.Printf("\tcase without break for i > 5 && 15 < i: %d\n", i)
            fallthrough
        default:
            fmt.Printf("\tdefault case without break for: %d\n", i)
    }
}