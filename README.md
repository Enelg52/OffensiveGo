# OffensiveGo - Golang Weaponization for your malwares.

This repo that contains some examples of offensives tools & utilities that can be used in a red team engagement rewrote in Golang. This project is also a way to train me to golang.

## Previous work :

These repo inspires me to make [OffensiveGo](https://github.com/RistBS/OffensiveGo)

- https://github.com/trickster0/OffensiveRust : Made by [trickster012](https://twitter.com/trickster012), this projects contains a bunch of example made in Rust
- https://github.com/byt3bl33d3r/OffensiveNim : 


## About Golang

- **Simpler syntax**: Go's syntax is simpler and easier to learn compared to Rust, which has a steeper learning curve.
- **Garbage collection**: Go uses garbage collection, which makes memory management easier for developers. Rust, on the other hand, uses a borrow checker to enforce memory safety, which can be more difficult to work with.
- **Cross-platform support**: Go has excellent cross-platform support and can be compiled to run on a wide range of platforms, including Windows, Linux, and macOS. Rust also has good cross-platform support, but its compilation process can be more complex.



## Examples 

| File                                                                                                   | Description                                                                                                                                                                              |
|--------------------------------------------------------------------------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| [Allocate_With_Syscalls](../master/Allocate_With_Syscalls/src/main.rs)                                 | It uses NTDLL functions directly with the ntapi Library                                                                                                                                  |
| [Create_DLL](../master/Create_DLL/src/lib.rs)                                                          | Creates DLL and pops up a msgbox, Rust does not fully support this so things might get weird since Rust DLL do not have a main function                                                  |
| [DeviceIoControl](../master/DeviceIoControl/src/main.rs)                                               | Opens driver handle and executing DeviceIoControl                                                                                                                                        |
| [EnableDebugPrivileges](../master/EnableDebugPrivileges/src/main.rs)                                   | Enable SeDebugPrivilege in the current process                                                                                                                                           |
| [Shellcode_Local_inject](../master/Shellcode_Local_inject/src/main.rs)                                 | Executes shellcode directly in local process by casting pointer                                                                                                                          |
| [Execute_With_CMD](../master/Execute_Without_Create_Process/src/main.rs)                               | Executes cmd by passing a command via Rust                                                                                                                                               |
| [ImportedFunctionCall](../master/ImportedFunctionCall/src/main.rs)                                     | It imports minidump from dbghelp and executes it                                                                                                                                         |
| [Kernel_Driver_Exploit](../master/Kernel_Driver_Exploit/src/main.rs)                                   | Kernel Driver exploit for a simple buffer overflow                                                                                                                                       |
| [Named_Pipe_Client](../master/Named_Pipe_Client/src/main.rs)                                           | Named Pipe Client                                                                                                                                                                        |
| [Named_Pipe_Server](../master/Named_Pipe_Server/src/main.rs)                                           | Named Pipe Server                                                                                                                                                                        |
| [PEB_Walk](../master/PEB_Walk/src/main.rs)                                                             | Dynamically resolve and invoke Windows APIs                                                                                                                                              |
| [Process_Injection_CreateThread](../master/Process_Injection_CreateThread/src/main.rs)                 | Process Injection in running process with CreateThread                                                                                                                                   |
| [Process_Injection_CreateRemoteThread](../master/Process_Injection_CreateRemoteThread/src/main.rs)     | Process Injection in remote process with CreateRemoteThread                                                                                                                              |
| [Process_Injection_Self_EnumSystemGeoID](../master/Process_Injection_Self_EnumSystemGeoID/src/main.rs) | Self injector that uses the EnumSystemsGeoID API call to run shellcode.                                                                                                                  |
| [Unhooking](../master/Unhooking/src/main.rs)                                                           | Unhooking calls                                                                                                                                                                          |
| [asm_syscall](../master/asm_syscall/src/main.rs)                                                       | Obtaining PEB address via asm                                                                                                                                                            |
| [base64_system_enum](../master/base64_system_enum/src/main.rs)                                         | Base64 encoding/decoding strings                                                                                                                                                         |
| [http-https-requests](../master/http-https-requests/src/main.rs)                                       | HTTP/S requests by ignoring cert check for GET/POST                                                                                                                                      |
| [patch_etw](../master/patch_etw/src/main.rs)                                                           | Patch ETW                                                                                                                                                                                |
| [ppid_spoof](../master/ppid_spoof/src/main.rs)                                                         | Spoof parent process for created process                                                                                                                                                 |
| [tcp_ssl_client](../master/tcp_ssl_client/src/main.rs)                                                 | TCP client with SSL that ignores cert check (Requires openssl and perl to be installed for compiling)                                                                                    |
| [tcp_ssl_server](../master/tcp_ssl_server/src/main.rs)                                                 | TCP Server, with port parameter(Requires openssl and perl to be installed for compiling)                                                                                                 |
| [wmi_execute](../master/wmi_execute/src/main.rs)                                                       | Executes WMI query to obtain the AV/EDRs in the host                                                                                                                                     |
| [Windows.h+ Bindings](../master/bindings.rs)                                                           | This file contains structures of Windows.h plus complete customized LDR,PEB,etc.. that are undocumented officially by Microsoft, add at the top of your file include!("../bindings.rs"); |
| [UUID_Shellcode_Execution](../master/UUID_Shellcode_Execution/src/main.rs)                             | Plants shellcode from UUID array into heap space and uses `EnumSystemLocalesA` Callback in order to execute the shellcode.                                                               |
| [AMSI Bypass](../master/amsi_bypass/src/main.rs)                                                       | AMSI Bypass on Local Process                                                                                                                                                             |
| [Injection_AES_Loader](../master/Injection_AES_Loader/src/main.rs)                                     | NtTestAlert Injection with AES decryption                                                                                                                                                |
| [Litcrypt_String_Encryption](../master/Litcrypt_String_Encryption/src/main.rs)                         | Using the [Litcrypt](https://github.com/anvie/litcrypt.rs) crate to encrypt literal strings at rest and in memory to defeat static AV.                                                   |
| [Api Hooking](../master/apihooking/src/main.rs)                                                        | Api Hooking using detour library                                                                                                                                                         |
| [memfd_create](../master/memfd_create/src/main.rs)                                                     | Execute payloads from memory using the memfd_create technique (For Linux)                                                                                                                |
| [RC4_Encryption](../master/Injection_Rc4_Loader/src/main.rs)                                           | RC4 Decrypted shellcode                                                                                                                                                                  |
| [Steal Token](../master/token_manipulation/src/main.rs) | Steal Token From Process|
| [Keyboard Hooking](../master/keyboard_hooking/src/main.rs) | Keylogging by hooking keyboard with SetWindowsHookEx |
| [memN0ps arsenal: shellcode_runner_classic-rs](https://github.com/memN0ps/arsenal-rs/blob/main/shellcode_runner_classic-rs/src/main.rs) | Classic shellcode runner/injector using `ntapi`                                                                                                                                                       |
| [memN0ps arsenal: dll_injector_classic-rs](https://github.com/memN0ps/arsenal-rs/blob/main/dll_injector_classic-rs/inject/src/main.rs)  | Classic DLL Injection using `windows-sys`                                                                                                                                                          |
| [memN0ps arsenal: module_stomping-rs](https://github.com/memN0ps/arsenal-rs/blob/main/module_stomping-rs/src/main.rs)                   | Module Stomping / Module Overloading / DLL Hollowing using `windows-sys`                                                                                                                                                              |
| [memN0ps arsenal: obfuscate_shellcode-rs](https://github.com/memN0ps/arsenal-rs/blob/main/obfuscate_shellcode-rs/src/main.rs)           | Simple shellcode XOR and AES obfuscator                                                                                                                                                                 |
| [memN0ps arsenal: process_hollowing-rs](https://github.com/memN0ps/arsenal-rs/blob/main/process_hollowing-rs/src/main.rs)               | Process Hollowing using `ntapi`                                                                                                                                                      |
| [memN0ps arsenal: rdi-rs](https://github.com/memN0ps/arsenal-rs/blob/main/rdi-rs/reflective_loader/src/loader.rs)                                  | Reflective DLL Injection using `windows-sys`                                                                                                                                                         |
| [memN0ps: eagle-rs](https://github.com/memN0ps/eagle-rs/blob/master/driver/src/lib.rs)                                                    | Rusty Rootkit: Windows Kernel Driver in Rust for Red Teamers using `winapi` and `ntapi`                                                                                                                                                       |
| [memN0ps: psyscalls-rs](https://github.com/memN0ps/psyscalls-rs/blob/main/parallel_syscalls/src/parallel_syscalls.rs)                                   | Rusty Parallel Syscalls library using `winapi`                                                                                                                                                      |
| [memN0ps: mmapper-rs](https://github.com/memN0ps/mmapper-rs/blob/main/loader/src/lib.rs)                                               | Rusty Manual Mapper using `winapi`                                                                                                                                   |
| [memN0ps: srdi-rs](https://github.com/memN0ps/srdi-rs/blob/main/reflective_loader/src/lib.rs)                                           | Rusty Shellcode Reflective DLL Injection using `windows-sys`                                                                                                                                               |
| [memN0ps: mordor-rs - freshycalls_syswhispers](https://github.com/memN0ps/mordor-rs/blob/main/freshycalls_syswhispers/tests/syscaller.rs)                 | Rusty FreshyCalls / SysWhispers1 / SysWhispers2 / SysWhispers3 library using `windows-sys`                                                                                                                                               |
| [memN0ps: mordor-rs - hells_halos_tartarus_gate](https://github.com/memN0ps/mordor-rs/blob/main/hells_halos_tartarus_gate/src/lib.rs)             | Rusty Hell's Gate / Halo's Gate / Tartarus' Gate Library using `windows-sys`                                                                                                                                           |
| [memN0ps: pemadness-rs](https://github.com/memN0ps/pemadness-rs/blob/main/pemadness/src/lib.rs)                                                   | Rusty Portable Executable Parsing Library (PE Parsing Library) using `windows-sys`                                                                                                                                                       |
| [memN0ps: mimiRust](https://github.com/memN0ps/mimiRust/blob/main/src/main.rs)                                                          | Mimikatz made in Rust by @ThottySploit. The original author deleted their GitHub account, so it's been uploaded for community use.                                                                                                                                                   |



## Interesting Tools in Golang

## OPSEC Considerations

## Credits

- 


### Golang Installation


`go build` for compilation 

Go binaries generally have no installation dependencies, compiler statically links Go runtime and needed packages. Static linking results in larger binaries
