package main

// TODO : Fixing img base reloc

import (
	"fmt"
	"strings"
	"syscall"
	"unsafe"
	"golang.org/x/sys/windows"
)

func StringToCharPtr(str string) *uint8 {
	chars := append([]byte(str), 0) // null terminated
	return &chars[0]
}


func main() {
	var in string
	var kern32Handle, _ = windows.LoadLibrary("kernel32.dll")
	var psapiHandle, _ = windows.LoadLibrary("psapi.dll")

	getModuleHandle, _ := windows.GetProcAddress(kern32Handle, "GetModuleHandleA")
	getCurrentProcessHandle, _ := windows.GetProcAddress(kern32Handle, "GetCurrentProcess")
	getModuleInfoHandle, _ := windows.GetProcAddress(psapiHandle, "GetModuleInformation")
	createFileAHandle, _ := windows.GetProcAddress(kern32Handle, "CreateFileA")
	createFileMappingHandle, _ := windows.GetProcAddress(kern32Handle, "CreateFileMappingA")
	mapViewOfFileHandle, _ := windows.GetProcAddress(kern32Handle, "MapViewOfFile")
	virtualProtectHandle, _ := windows.GetProcAddress(kern32Handle, "VirtualProtect")
	rtlCopyMemoryHandle, _ := windows.GetProcAddress(kern32Handle, "RtlCopyMemory")

	ntdllName := StringToCharPtr("C:\\windows\\system32\\ntdll.dll")
	ntdllModuleHandle, _, err := syscall.Syscall(getModuleHandle, 1, uintptr(unsafe.Pointer(ntdllName)), 0, 0)

	processHandle, _, err := syscall.Syscall(getCurrentProcessHandle, 0, 0, 0, 0)

	var modInfo windows.ModuleInfo
	_, _, err = syscall.Syscall6(getModuleInfoHandle, 4, processHandle, ntdllModuleHandle, uintptr(unsafe.Pointer(&modInfo)), unsafe.Sizeof(modInfo), 0, 0)

	fileName := StringToCharPtr("C:\\windows\\system32\\ntdll.dll")
	ntdllHandle, _, err := syscall.Syscall9(createFileAHandle, 7, uintptr(unsafe.Pointer(fileName)), windows.GENERIC_READ, windows.FILE_SHARE_READ, 0, windows.OPEN_EXISTING, 0, 0, 0, 0)                       
	if syscall.Handle(ntdllHandle) == syscall.InvalidHandle {
		fmt.Println(err)
		return
	}

	// create file mapping object
	fileMapObj, _, err := syscall.Syscall6(createFileMappingHandle, 6, ntdllHandle, 0, windows.PAGE_READONLY|0x01000000, 0, 0, 0)
	if syscall.Handle(fileMapObj) == 0 {
		fmt.Println(err)
		return
	}

	ntdllMapping, _, err := syscall.Syscall6(mapViewOfFileHandle, 5, fileMapObj, windows.FILE_MAP_READ, 0, 0, 0, 0)
	if syscall.Handle(ntdllMapping) == 0 {
		fmt.Println(err)
		return
	}

	dosHeaderELfanewOffset := uintptr(0x3c)     
	ntHeaderNumOfSectionsOffset := uintptr(0x06) 
	dosHeaderAddr := modInfo.BaseOfDll
	ntHeaderAddr := uintptr(modInfo.BaseOfDll + uintptr(*((*uint32)(unsafe.Pointer(dosHeaderAddr + dosHeaderELfanewOffset)))))

	sizeSectionHeader := uintptr(40)
	numOfSections := *(*uint16)(unsafe.Pointer(ntHeaderAddr + ntHeaderNumOfSectionsOffset))
	sizeOfOptionalheader := uintptr(*(*uint16)(unsafe.Pointer(ntHeaderAddr + 0x4 + 0x10))) 
	sectionHeadersAddr := ntHeaderAddr + 0x04 + 0x14 + sizeOfOptionalheader

	for i := uintptr(0); i < uintptr(numOfSections); i++ {
		curSectionHeaderAddr := sectionHeadersAddr + uintptr(unsafe.Pointer(sizeSectionHeader*i))
		sectionName := unsafe.Slice((*byte)(unsafe.Pointer(curSectionHeaderAddr)), 8)

		if strings.Contains(string(sectionName), ".text") {
			var oldProtect uint32 = 0

			ntdllVirtualAddress := uintptr(*(*int32)(unsafe.Pointer(curSectionHeaderAddr + 0x0C))) 
			ntdllVirtualSize := uintptr(*(*int32)(unsafe.Pointer(curSectionHeaderAddr + 0x08)))    
			syscall.Syscall6(virtualProtectHandle, 4, ntdllModuleHandle+ntdllVirtualAddress, ntdllVirtualSize, windows.PAGE_EXECUTE_READWRITE, uintptr(unsafe.Pointer(&oldProtect)), 0, 0)

			fmt.Scanln(&in)
			r1, _, err := syscall.Syscall(rtlCopyMemoryHandle, 3, ntdllModuleHandle+ntdllVirtualAddress, ntdllMapping+ntdllVirtualAddress, ntdllVirtualSize)

			syscall.Syscall6(virtualProtectHandle, 4, ntdllModuleHandle+ntdllVirtualAddress, ntdllVirtualSize, uintptr(oldProtect), uintptr(unsafe.Pointer(&oldProtect)), 0, 0)

		}
	}
	fmt.Scanln(&in)
}

