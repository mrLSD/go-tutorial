// Basic testing for slices/arrays/maps
package utils

import "fmt"

/** Array testing
 */
func TestArrays() {
    print("======\nTest arrays\n")
    var (
        arr1 [2]int // different types
        arr2 [3]int
    )
    arr1[0] = 0xFA1
    arr1[1] = 23
    arr2[0] = 12
    arr2[1] = 24
    arr2[2] = 35
    arr3 := [...]int{-1, 0, 3}
    fmt.Printf("\t %#v \n\t %#v \n\t %#v \n", arr1, arr2, arr3)
    
    // Autocreate array [11]capacity 
    // and fill it `0` items
    // exclude last element - that pre-init
    arr4 := [...]int{10:123}
    fmt.Printf("Aitocreate array:\n\t %#v \n", arr4)
}

/** Slices testing
    Slices is common case for arrays
    Data structures that must be initialized before use
 */
func TestSLices() {
    print("======\nTest slices\n")
    var arr1, arr3 []int
    arr1 = make([]int, 2)
    arr1[0] = 10
    arr1[1] = 20
    
    arr2 := []int{3, 6, 7}
    // Append examples
    arr3 = append(arr3, 1, 4, 5)
    arr3 = append(arr3, arr1...)
    arr3 = append(arr3, []int{7}...)
    fmt.Printf("Slices: \n\t %#v \n\t %#v \n\t %#v \n", arr1, arr2, arr3)

    fmt.Printf("Slice: \n\t length:\t %d \n\t capacity:\t %d \n ", 
        len(arr3), cap(arr3))

    // From index (including)
    startIndex := 4
    // get items count, including start index
    sliceLength := 2
    fmt.Printf("Slice [%d:%d]: \n\t %#v \n", 
        startIndex, sliceLength,
        arr3[startIndex:startIndex+sliceLength])
        
    // From index (including)
    startIndex = 4
    // get items count, including start index
    // Slice lenth get via capacity of array
    sliceLength = cap(arr3) - startIndex
    fmt.Printf("Slice [%d:%d]: \n\t %#v \n", 
        startIndex, sliceLength,
        arr3[startIndex:startIndex+sliceLength])
    testSliceReturn(arr3, arr2)
    fmt.Printf("Changed slice: \n\t %v \n\t %v \n", arr3, arr2)
}

/** Map that like associative arrays, hash, dict
 */
func TestMaps() {
    print("======\nTest maps\n")
    m1 := map[int]string{7: "seven"}
    m1[1] = "one"
    m1[3] = "three"
    
    var m2 map[string]int
    m2 = make(map[string]int, 3)
    m2["one"] = 100
    
    // Check is inde exist via `ok`
    if three, ok := m1[3]; ok {
        println("Test map index:", three)
    }
    fmt.Printf("Maps:\n\t %#v \n\t %#v \n", m1, m2)
}

func testSliceReturn(s1, s2 []int) {
    // slice changed only at curreny func
    s1 = append(s1, []int{6, 5, 4}...)
    s1[0] = 100
    
    // slice changed outsid of curreny func
    s2[0] = 100
    
    // Check index of slice
    // index should checked via `len` - not `cap` or `ok`
    s2 = append(s2, []int{4}...)
    fmt.Printf("%v <-", cap(s2))
    ind := 8
    if ind < len(s1) {
        fmt.Print("Check slice index: \n\t [", ind, "] = ", s1[len(s1)-1], "; // check [ok]\n")
    }

    fmt.Printf("Change slice: \n\t %v \n\t %v \n", s1, s2)
}