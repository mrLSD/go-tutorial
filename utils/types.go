// Types testing
package utils

import "fmt"

// Basic struct
type Vector struct {
	x, y int
	name string
}

/** Vector value copied - not mutable
 */
func testInlineVect(v Vector, x, y int) {
	v.x = x
	v.y = y
	fmt.Printf("\tTest Vector type %v\n", v)
}

/** Method for type Vector
  not mutable attributes
*/
func (v Vector) testInlineVect(x, y int) {
	v.x = x
	v.y = y
	fmt.Printf("\tTest Vector type %v\n", v)
}

func pointerVect(p *Vector, x, y int) {
	p.x = x
	p.y = y
	fmt.Printf("\tTest Vector type - mutable %v\n", p)
}

func TestVect() {
	println("=================\nTest types")
	v1 := Vector{10, 20, "Test1"}
	testInlineVect(v1, 1, 2)
	v1.testInlineVect(3, 4)
	fmt.Println(v1)

	// Test for pointer
	p1 := &v1
	pointerVect(p1, 5, 6)
	fmt.Println(v1)
}
