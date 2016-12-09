package main

import (
	"fmt"
	"github.com/alexflint/go-arg"
	"github.com/mrlsd/go-tutor/gbe"
	"github.com/mrlsd/go-tutor/utils"
)

type args struct {
	Rbe bool `arg:"-r,help:use Go By Examples"`
}

func (*args) Version() string {
	return "Go Tutorials v0.1"
}

func main() {
	var args args

	//args.Rbe = true
	arg.MustParse(&args)

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
	//utils.TestFiles()
	utils.TestWrkDependencies()
	utils.CheckLoopForLinkList()

	if args.Rbe {
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
