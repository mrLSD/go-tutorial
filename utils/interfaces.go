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
    name, msg   string
}

type cat struct {
    name, msg   string
}

type man struct {
    name, msg   string
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