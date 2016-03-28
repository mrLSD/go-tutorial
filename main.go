package main

import (
    "fmt"
    "./utils"
)

func main() {
    fmt.Printf("Arrays: %v\n", arrays())
    
    utils.ForTest()
    utils.TestArrays()
    utils.TestSLices()
    utils.TestMaps()
    utils.TestVect()
    utils.TestFunc()
    utils.TestStruct()
    utils.TestInterfaces()
    utils.TestErrors()
}

// Test simple arrays
func arrays() [3]int {
    var ar1 [3]int
    // var ar2 [...]int
    //ar1[0] = 3

    // ar2[1] = 2
    return ar1
}