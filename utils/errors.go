package utils

import ( 
    "fmt"
    "errors"
)

func TestErrors() {
    fmt.Printf("=======================\nTest errors\n")
    if _, err := testErr(-1); err != nil {
        fmt.Printf("\tSome error: %v\n", err)
    }
    if _, err := testErr(21); err == nil {
        fmt.Printf("\tNo errors: %v\n", err)
    }
    fmt.Printf("\tTest errors - no error\n")
}

/** Test errors
    By convention, errors are the last return value 
    and have type error, a built-in interface
 */
func testErr(i int) (int, error) {
    if i < 0 {
        // errors.New constructs a basic error value 
        // with the given error message
        return i*0, errors.New("Should > 0")
    }
    // A nil value in the error position 
    // indicates that there was no error
    return i*10, nil
}