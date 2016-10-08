package main

import (
    "fmt"
    "flag"
    "./utils"
    "./gbe"
)

func main() {
    useGBE := flag.Bool("gbe", false, "Use Go by Examples")
    flag.Parse()

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
    utils.TestChannels()

    utils.TestStructs()

    if *useGBE {
        gbe.Run()
    }
}

// Test simple arrays
func arrays() [3]int {
    var ar1 [3]int
    // var ar2 [...]int
    //ar1[0] = 3

    // ar2[1] = 2
    return ar1
}