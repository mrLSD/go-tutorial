package utils

import "fmt"

// Basic struct
type testStruct struct {
    val1, val2  int
    val3        float32
    val4        bool
}

// Struct with func
type testStructWithFn struct {
    val1, val2  int
    fn          func (v1 int) float32
}

/** Test structs
 */
func TestStruct() {
    fmt.Printf("\n======================\nStructs\n")  
    fnStruct()
}

/** Basic test for structs
 */
func fnStruct() {
    // Create new struct and fill values
    // all struct values shoud be set
    s1 := testStruct{1, 3, 4.1, true}
    fmt.Printf("\tInit structs: %#v\n", s1)   
    
    sf1 := testStructWithFn{4, 5, func(x int) float32 {return float32(x)+.1}}
    fmt.Printf("\tInit structs with Func: %#v\n", sf1)
    fmt.Printf("\tStructs with Func - result: %#v\n", sf1.fn(5))
    sf1.fn = func (res int) float32 {
        return float32(res) * .2 / 0.9
    }
    fmt.Printf("\tStructs with Func - result: %#v\n", sf1.fn(5))
    
    var s2, s3, s4 testStruct
    // Set struct value by name and optionally one value
    s2 = testStruct{val1: 4}
    s2.val2 = 10
    fmt.Printf("\tInit struct by value: %#v\n", s2)   
    
    // Set struct value by name and optionally one value
    // init position doesn't metter
    s3 = testStruct{val4: true, val2: 20}
    fmt.Printf("\tInit struct by value: %#v\n", s3)
    
    // Set optional values explicitly
    s4.val1 = 10
    s4.val3 = 1.2
    fmt.Printf("\tInit struct by value: %#v\n", s4)
    
    // Test struct as ptr
    var p1 *testStruct
    p1 = &testStruct{val3: 4.1, val2: 23}
    // it's alias for (*o1).val1 = ...
    p1.val1 = 30
    (*p1).val4 = true
    fmt.Printf("\tInit struct as Ptr: %#v\n", *p1)
    
    // Test struct methods
    fmt.Printf("\tStruct method1: %#v\n", s1.testMethod1())
    // Test metho with pointer reciever
    // s1.testMethod2() - alias (&s1).testMethod2()
    res1 := s1.testMethod2()
    res2 := (&s1).testMethod2()
    // For pointer reciever with don't need set *p1
    res3 := p1.testMethod2()
    fmt.Printf("\tStruct method2 - pointer reciever: %#v %#v %#v\n", res1, res2, res3)
    
    fmt.Printf("\tStruct method3 mutable - befort: %#v\n", s1)
    // s1.testMethod3() alias for - (&s1).testMethod3()
    fmt.Printf("\tStruct method3 - pointer reciever and mutable values: %#v\n", s1.testMethod3())
    fmt.Printf("\tStruct method3 mutable - after: %#v\n", s1)
    
    fmt.Printf("\tStruct method3 mutable - befort: %#v\n", p1)
    fmt.Printf("\tStruct method3 - pointer reciever and mutable values: %#v\n", p1.testMethod3())
    fmt.Printf("\tStruct method3 mutable - after: %#v\n", p1)
}

func (s testStruct) testMethod1() int {
    return s.val1 + s.val2
}

func (p *testStruct) testMethod2() int {
    return p.val1 + p.val2
}

func (p *testStruct) testMethod3() int {
    p.val1 += 30
    return p.val1 + p.val2
}