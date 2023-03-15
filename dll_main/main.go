package main

/*
https://gist.github.com/NaniteFactory/7a82b68e822b7d2de44037d6e7511734
*/

import "C"

import (
	"unsafe"

	"github.com/nanitefactory/winmb"
)

//export Test
func Test() {
	winmb.MessageBoxPlain("export Test", "export Test")
}

// OnProcessAttach is an async callback (hook).
//
//export OnProcessAttach
func OnProcessAttach(
	hinstDLL unsafe.Pointer, // handle to DLL module
	fdwReason uint32, // reason for calling function
	lpReserved unsafe.Pointer, // reserved
) {
	winmb.MessageBoxPlain("DLL_PROCESS_ATTACH", "DLL_PROCESS_ATTACH")
}

func main() {
}
