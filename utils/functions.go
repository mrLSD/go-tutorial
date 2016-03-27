package utils

import "fmt"

func TestFunc() {
    fmt.Printf("\n===========================\nTest functions\n")
    testBasicFunc1(10, 20)
    testBasicFunc2(10, 20)
    
    _, _, c := testBasicFunc3(10, 20, 3.)
    fmt.Printf("\tFloat after transform type : %f\n", c)
    
    x, y := testBasicFunc4(4, 7)
    fmt.Printf("\tFinc after reverce multiple vars: %d %d\n", x, y)
    
    x, y = testBasicFunc5(4, 7, 6)
    fmt.Printf("\tFinc after variadic params: %d %d\n", x, y)
    
    x, y = testBasicFunc5(4, 7, 6, 3)
    fmt.Printf("\tFinc after variadic params: %d %d\n", x, y)
}

/** Test functions usage
    We can acept several input variables 
    separated by comma and grouped for type 
    by comma. For return - same logic
 */
func testBasicFunc1(a, b int) int {
    fmt.Printf("\tTest func params: %d\n", a + b)
    return a + b
}

/** Func multi results
 */
func testBasicFunc2(a, b int) (int, int) {
    fmt.Printf("\tTest func params: %d %d\n", a + b, b - a)
    return a + b, b - a
}

/** Different params for func with different types
    and different types result
 */
func testBasicFunc3(a, b int, c float32) (int, int, float32) {
    fmt.Printf("\tTest func with different params: %d %d %#v\n", a + int(c), b - a, float32(a) - c)
    return int(c), a, float32(b)
}

/** Named multiple results for func
 */
func testBasicFunc4(a, b int) (x int, y int) {
    x, y = b, a
    return
}

/** Variadic params for func
 */
func testBasicFunc5(a int, b ...int) (x, y int) {
    x = a
    y = x
    for _, res := range b {
        y += res
    }
    return
}