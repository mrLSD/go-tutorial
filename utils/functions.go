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
    fmt.Printf("\tFunc after variadic params [3]: %d %d\n", x, y)
    
    x, y = testBasicFunc5(4, 7, 6, 3)
    fmt.Printf("\tFunc after variadic params [4]: %d %d\n", x, y)
    
    // Test with predeclared slice
    arr := []int{1, 3, 5, 9}
    x, y = testBasicFunc5(4, arr...)
    fmt.Printf("\tFunc after variadic params with predeclared slice: %d %d\n", x, y)
    
    lambda := testLambda(10)
    fmt.Printf("\tTest lambda with predefined state: %d\n", lambda())
    fmt.Printf("\tTest lambda with predefined state: %d\n", lambda())
    fmt.Printf("\tTest lambda with predefined state: %d\n", lambda())
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

/** Test lambda func
    That func return func but
    statement/environment PREDEFINED!!!
 */
func testLambda(predefinded int) func() int {
    i := predefinded
    return func() int {
        i += 1
        return i
    }
}