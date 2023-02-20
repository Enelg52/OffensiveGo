# OffensiveGo - Golang Weaponization for your malwares.

![image](https://user-images.githubusercontent.com/75935486/220174425-bc827a51-ffdb-4da2-bbd6-161c5e46c340.png)


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




## Interesting Tools in Golang

## OPSEC Considerations

## Credits

- 


### Golang Installation


`go build` for compilation 

Go binaries generally have no installation dependencies, compiler statically links Go runtime and needed packages. Static linking results in larger binaries
