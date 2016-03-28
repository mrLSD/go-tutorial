package utils

import "fmt"

/** Basic interfaces usage
 */
func TestInterfaces() {
    fmt.Printf("=============================\nBasic usage for interfaces\n")
    basicInterface()
    
    d := dog{name: "A dog", msg: "Gafff~~"}
    interfaceAsTypePAram(d)
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