package util

import "sync"

// CPUUsed ... set to init
var CPPUused int

// WaitG ...
var WaitG syn.WaitGroup

// Mu ...
var Mu sync.Mutex

func CheckErr(err error) {
	if err != nil {
		// pannic err
		panic(err.Error())
	}
}
