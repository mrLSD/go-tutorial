package utils

import "fmt"

/** Basic interfaces usage
 */
func TestInterfaces() {
	fmt.Printf("=============================\nBasic usage for interfaces\n")
	basicInterface()

	d := dog{name: "A dog", msg: "Gafff~~"}
	c := cat{name: "A caaat", msg: "r-r-Myauuu... ))"}
	interfaceAsTypePAram(d)

	testTypeAssertion(d)
	testTypeAssertion(c)

	i := 100500
	testTypeAssertion2(i)

	res := implicitInterface()
	fmt.Printf("\tImplicit Interface msg: %v\n", res)

	emptyInterface()
	interfaceTypeSwitching()
}

/** Test interface for types that implement
  interface method
*/
func basicInterface() {
	var animal []animals
	animal = []animals{
		dog{"A dog", "Gaff"},
		cat{"A cat", "Myau"},
		man{"A man", "Oh, hi!"}}

	for _, str := range animal {
		fmt.Printf("\tInterface msg: %v\n", str.about())
	}
}

/** Test interface as param for Func
  common interface for types thay implement interface
*/
func interfaceAsTypePAram(param animals) {
	fmt.Printf("\tInterface as param msg: %v\n", param.about())
}

// Basic interface
type animals interface {
	about() string
}

type dog struct {
	name, msg string
}

type cat struct {
	name, msg string
}

type man struct {
	name, msg string
}

func (d dog) about() string {
	return d.name + " - " + d.msg
}

func (c cat) about() string {
	return c.name + " - " + c.msg
}

func (m man) about() string {
	return m.name + " - " + m.msg
}

// Type for pointer testing
type panimal struct {
	name, msg string
}

// Pointer reciver
// And invoked via interface with not init reciver
// we should check ptr - is it NIL
func (p *panimal) about() string {
	fmt.Printf("\tNil interface: %#v %T\n", p)
	if p == nil {
		return "our result - <nil>"
	}
	return p.name + " - " + p.msg
}

/** Type assertions make possible checke type for
  value of interface type
*/
func testTypeAssertion(val animals) {
	if v, ok := val.(dog); ok {
		fmt.Printf("\tTest Type Assertion. It's valid type: %#v\n", v)
	} else {
		fmt.Printf("\tTest Type Assertion. Isn't valid type: %#v\n", val)
	}
}

/** Empty interface sutisfy for any type
  All types satisfy the empty interface
  and val could be any type
*/
func testTypeAssertion2(val interface{}) {
	// Without `ok` -if not valid type
	// program exit with `panic`!!
	if v, ok := val.(dog); ok {
		fmt.Printf("\tTest Type Assertion. It's valid type: %#v\n", v)
	} else {
		fmt.Printf("\tTest Type Assertion. Isn't valid type: %#v\n", val)
	}
}

/** Test implicit interface init
 */
func implicitInterface() string {
	var a animals = dog{"Dog", "Coold"}
	return a.about()
}

/** Test empty interface and value init via that
  And test <nil> interfaces and invoking thier methods
*/
func emptyInterface() {
	// Empty interface
	var i interface{}

	i = "it's string"
	fmt.Printf("\tTest empty interface - valuse association: %v <%T>\n", i, i)

	i = 43
	fmt.Printf("\tTest empty interface - valuse association: %v <%T>\n", i, i)

	// Test `nil` interface
	// that case only for <ptr> cause not ptr fill struct
	// fields as empty values
	var ia animals
	var p *panimal
	ia = p
	i = ia.about()
	fmt.Printf("\tTest <nil> interface - valuse not filled: %v <%T>\n", i, i)

	// Fill interface via pointer init
	p = &panimal{"Ptr", "Msg..."}
	ia = p
	i = ia.about()
	fmt.Printf("\tTest <nil> interface - valuse filled: %v <%T>\n", i, i)

}

/** Test type switching via interface declaration
  The declaration in a type switch has the same
  syntax as a type assertion i.(T), but the
  specific type T is replaced with the keyword type.
*/
func interfaceTypeSwitching() {
	// Empty interface
	var i interface{}
	var d dog = dog{"The dog", "G-aff"}
	i = d
	// Interface type switch via constuction i.(type)
	// that's it type assertion variations
	switch t := i.(type) {
	case cat:
		fmt.Printf("\tInterface type switch: %#v\n", t)
	case dog:
		fmt.Printf("\tInterface type switch: %#v\n", t)
	}
}
