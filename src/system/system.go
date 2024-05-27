package system

// #cgo CFLAGS: -g -Wall
// #include "system.h"
import "C"
import (
	"fmt"
	"runtime"
)

var createWindow map[string]func()

func init() {
	createWindow = make(map[string]func())
	createWindow["linux"] = func() {
		C.connectToWayland()
	}
	createWindow["windows"] = func() {
	}
}

func CreateWindowForEnv() {
	value, ok := createWindow[runtime.GOOS]
	if ok {
		value()
	} else {
		fmt.Println("Error creating window")
	}
}
