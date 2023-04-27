# OffensiveGo - Golang Weaponization for red teamers.

![image](https://user-images.githubusercontent.com/75935486/220217814-242de1ba-1f62-4b0b-a1be-6cf8b82ab0da.png)


**This repo is made by [@RistBS](https://twitter.com/RistBs), [@Enelg](https://twitter.com/Enelg_) & [@dreamkinn](https://twitter.com/dreamkinn) and contains some examples of offensives tools & utilities rewrote in Golang that can be used in a red team engagement.**

## Table of Content

- [Previous work](#previous-work)
- [About Golang](#about-golang)  
  - [Installation](#installation)
  - [Workspace Setup](#setup)
  - [Compilation](#compilation)
  - [Assembly in Go]()
- [Examples](#examples)
- [Interesting Tools in Golang](#interesting-tools-in-golang)
- [Golang Binary](#golang-binary)
- [Credits](#credits)

## Previous works

These repo inspires me to make [OffensiveGo](https://github.com/RistBS/OffensiveGo)

- [OffensiveRust](https://github.com/trickster0/OffensiveRust) : this project contains a bunch of examples made in [Rust](https://www.rust-lang.org/).
- [OffensiveNim](https://github.com/byt3bl33d3r/OffensiveNim) : this one contains examples written in [Nim](https://nim-lang.org/).
- [OffensiveCSharp](https://github.com/matterpreter/OffensiveCSharp) : A Collection of Offensive C# Tooling.
- [OffensiveDLR](https://github.com/byt3bl33d3r/OffensiveDLR) : Toolbox containing research notes & PoC code for weaponizing .NET's [DLR](https://learn.microsoft.com/en-us/dotnet/framework/reflection-and-codedom/dynamic-language-runtime-overview).
- [OffensiveVBA](https://github.com/S3cur3Th1sSh1t/OffensiveVBA) : This repo covers some code execution and AV Evasion methods for Macros in Office documents.
- [OffensiveZig](https://github.com/darkr4y/OffensiveZig) : Some attempts at using [Zig](https://ziglang.org/) in penetration testing.


## About Golang

- **Simpler syntax**: Go's syntax is simpler and easier to learn.
- **Garbage collection**: Go uses garbage collection, which makes memory management easier for developers.
- **Cross-platform support**: Go has excellent cross-platform support and can be compiled to run on a wide range of platforms, including Windows, Linux, and macOS. Rust also has good cross-platform support, but its compilation process can be more complex.
- **Goroutine**:  Goroutines are lightweight threads of execution that enable concurrent programming in Go, making it easy to write efficient, safe, and scalable concurrent programs, allowing for non-blocking concurrent execution and communication via channels.

**OPSEC Consideration & Caveats of Golang**

Go binaries generally have no installation dependencies, compiler statically links Go runtime and needed packages. Static linking results in larger binaries. 2 MB for "Hello World" compared to 54 KB in C.

### Workspace Setup


- create a `go.mod` file in your project directory, which will be used to manage dependencies :
```ps
> go mod init offensivego 
```

- ensure that all your project's dependencies are up to date :
```ps
> go mod tidy 
```

### Compilation

- Use `go build file.go` for compilation.
- Omit debug symbols and strip the symbol table. it can also reduce binary size by about 30% : `go build -ldflags="-s -w" file.go`
- Hide console, to avoid Go program displaying console windows on execution : `go build -ldflags -H=windowsgui rshell.go`

### Assembly in Go

The [assembly language](https://go.dev/doc/asm) used with Go is a high-level and is based on [Plan9 (P9)](https://9p.io/sys/doc/asm.html), architecture-independent language that includes mnemonics like `CALL` and `RET`, as well as higher-level constructs like loops and conditionals, which are implemented using lower-level assembly instructions by the assembler.

- That's how you declare function :
  ![image](https://user-images.githubusercontent.com/75935486/234733526-bacaf609-b59a-4b99-a4f2-f708f65a0909.png)
  - **NOSPLIT** : Don't insert the preamble to check if the stack must be split. The frame for the routine, plus anything it calls, must fit in the spare space remaining in the current stack segment. Used to protect routines such as the stack splitting code itself, which can improve performance.
  - **NOFRAME** : skip the generation of a function prologue and epilogue, even if this is not a leaf function, which can also improve performance by reducing the overhead of setting up and tearing down the stack frame for each call.
  
> **Note** : It can be useful to use Assembly in Go for your loaders if you want to make a custom VEH or implementing direct/indirect syscall stub.

## Examples 

| File                                                              | Description                                                                                                                                |
|-------------------------------------------------------------------|--------------------------------------------------------------------------------------------------------------------------------------------|
| [Process Injection - APC](../main/injection_native_apc/main.go)   | Execute a shellcode with `NtQueueApcThread`                                                                                                |
| [Process Injection - CreateThread](../main/injection_thread)      | Execute a shellcode with `NtCreateThreadEx` and `CreateThread`                                                                             |
| [Whoami](../main/whoami/main.go)                                  | rebuilt whoami process to show current user, groups & privileges                                                                           |
| [EnableDebugPrivileges](../main/enable_debug_priv/main.go)        | Enable SeDebugPrivilege in the current process                                                                                             |
| [ACG + BlockDll](../main/acg_blockdll_process/main.go)            | Apply Arbitrary Code Guard (ACG) & BlockDll policy on your process                                                                         |                                                                             
| [Module Refreshing]()                                             | Refresh the .text section of a module in memory (NTDLL here)                                                                               | 
| [Detect Hooks](../main/detect_hooks/main.go)                      | Detect Hooks set by AV/EDR on NTDLL                                                                                                        |
| [Sleep Obfuscation](../main/sleep_obfuscation/main.go)            | Perform Sleep Obfuscation with Timer Queue Timers                                                                                                |
| [AMSI Patching](../main/amsi_bypasses/)               | bypass AMSI by patching in memory `AmsiScanBuffer`                       |
| [ETW Patching](../main/etw_bypasses/)                 | bypass ETW, by patching in memory with `ret` on `NtTraceControl`                                   |
| [Network](../main/network)                                        | TCP, HTTP, and named pipes servers and clients for different communication channels.                                                       |
| [WMI Query](../main/wmi/wmi.go)                                         | List the av/edr solution with a wmi query                                                                                                  |
| [sRDI](../main/srdi/srdi.go)                                      | Convert DLL files to position independent shellcode                                                                                        |
| [Cryptography](../main/crypto)                                    | Encryption algorithms for various usage. Contains AES, RC4, chacha20 and xor.                                                              |
| [Self Remove](../main/self_remove/self_remove.go)                 | Self remove a executable. Golang implementation of [delete-self-poc](https://github.com/LloydLabs/delete-self-poc)                                          |
| [Process Dump](../main/process_dump/process_dump.go)              | Dump any process with `MiniDumpWriteDump`. In this example, it dumps LSASS                                                                     |
| [Dllmain](../main/dll_main)                                       | `DllMain()` entrypoint in Golang from [this](https://gist.github.com/NaniteFactory/7a82b68e822b7d2de44037d6e7511734). Can be used for dll hijacking. |
| [Token Manipulation](../main/token)                                            | Play with windows token. List all the token, `ImpersonateLoggedOnUser` and `CreateProcessWithToken`.                                       |
| [Sandbox detection/evasion](../main/sandbox)| Sandbox detection and evasion techniques |
| [Callback Injection](../main/callback_injection)| Callback shellcode injection using `GrayStringA`, `EnumFonts` and more... |
| [Instrumentation Callback](../main/instrumentation_callback)| Disable Instrumentation Callback on your process to reduce any potential direct syscall detection  |

> **Note** : The misc contains some scripts that can be written in other languages with a link with Go language


## Interesting Tools in Golang

- [Acheron](https://github.com/f1zm0/acheron) : Indirect syscalls for AV/EDR evasion in Go assembly.
- [Sliver](https://github.com/BishopFox/sliver) : An Adversary Emulation Framework fully written in Golang with advanced evasion capabilities.
- [BananaPhone](https://github.com/C-Sto/BananaPhone) : An easy to use GO variant of [Hells gate](https://github.com/am0nsec/HellsGate) with automatic SSN parsing.
- [SourcePoint](https://github.com/Tylous/SourcePoint) : C2 profile generator for Cobalt Strike command and control servers designed to ensure evasion by reducing the Indicators of Compromise IoCs.
- [ScareCrow](https://github.com/optiv/ScareCrow) : Payload creation framework designed around EDR bypass such as AMSI & ETW Bypass, Encryption, Stealth Process Injections, ect.
- [RedGuard](https://github.com/wikiZ/RedGuard) : a C2 front flow control tool, Can avoid Blue Teams, AVs, EDRs check.
- [Freeze](https://github.com/optiv/Freeze) : Payload toolkit for bypassing EDRs using suspended processes, direct syscalls, and alternative execution methods.
- [Mangle](https://github.com/optiv/Mangle) : A tool that manipulates aspects of compiled executables (.exe or DLL) to avoid detection from EDRs.
- [Dent](https://github.com/optiv/Dent) : A framework for creating COM-based bypasses utilizing vulnerabilities in Microsoft's WDAPT sensors.
- [Ivy](https://github.com/optiv/Ivy) : Payload creation framework for the execution of arbitrary VBA (macro) source code directly in memory. Ivy’s loader does this by utilizing programmatical access in the VBA object environment to load, decrypt and execute shellcode.


## Golang Binary


- The `symtab` (symbol table) section contains information about the symbols defined in an executable file or shared library. The symtab section is generated by the Go linker.
![image](https://user-images.githubusercontent.com/75935486/233974681-182d7770-c709-4035-ad23-637d3d1c136f.png)



## Credits
- [@joff_thyer](https://twitter.com/joff_thyer) - https://www.youtube.com/watch?v=gH9qyHVc9-M&t=1131s&ab_channel=BlackHillsInformationSecurity
- [@BlueSentinelSec](https://twitter.com/BlueSentinelSec) - https://github.com/bluesentinelsec/OffensiveGoLang/blob/master/Offensive%20GoLang%202.0%20-%20SANS%20Pen%20Test%20HackFest%202021.pdf
