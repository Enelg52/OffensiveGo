package main

import (
    "fmt"
    // "windows"
	"unsafe"
	"golang.org/x/sys/windows"
	"time"
)

// var (
// 	u32           = windows.NewLazyDLL("user32.dll")
// 	k32           = windows.NewLazyDLL("kernel32.dll")
// 	GetCurrentProcess = k32.NewProc("GetCurrentProcess")
// 	ReadProcessMemory = u32.NewProc("ReadProcessMemory")
// )

func main() {

	// Get a handle to the current process
	hProcess, err := windows.GetCurrentProcess()
    if err != nil {
        panic(err)
    }

    // Get the base address of the process module
    var mod windows.ModuleInfo
	size := uint32(unsafe.Sizeof(mod))
    err = windows.GetModuleInformation(hProcess, 0, &mod, size)
    if err != nil {
        panic(err)
    }
    baseAddress := uintptr(mod.EntryPoint)

    // Read the PEB address from the process memory
    
	var pebAddress uintptr
    var bytesRead uintptr
	// 0x30 for 32-bit, 0x60 for 64-bit
    err = windows.ReadProcessMemory(hProcess, uintptr(baseAddress+0x60), (*byte)(unsafe.Pointer(&pebAddress)), unsafe.Sizeof(pebAddress), &bytesRead)
    if err != nil {
        panic(err)
    }

	fmt.Printf("PEB 1: 0x%X\n", pebAddress)
	fmt.Printf("PEB 1: 0x%X\n", baseAddress)
    fmt.Printf("PEB offset: 0x%X\n", pebAddress-baseAddress)
	time.Sleep(10000 * time.Second)
}



// https://github.com/shenwei356/rush/blob/master/process/process_windows.go
// ;===========FUNCTIONS=============
// ;=======Function : Get Kernel32 base address============
// ;Technique : PEB InMemoryOrderModuleList
// sub esp, 0x4             ; reserve stack space for called functions
// push esi
// xor eax, eax             ; clear eax
// xor ebx, ebx
// mov bl,0x30
// mov eax, [fs:ebx ]      ; get a pointer to the PEB
// mov eax, [ eax + 0x0C ]  ; get PEB->Ldr
// mov eax, [ eax + 0x14 ]  ; get PEB->Ldr.InMemoryOrderModuleList.Flink (1st entry)
// push eax
// pop esi
// mov eax, [ esi ]         ; get the next entry (2nd entry)
// push eax
// pop esi
// mov eax, [ esi ]         ; get the next entry (3rd entry)
// mov eax, [ eax + 0x10 ]  ; get the 3rd entries base address (kernel32.dll)
// pop esi


// fn get_module_base_addr(module_name: &str) -> HINSTANCE {
//     unsafe {
//         let peb_offset: *const u64 = __readgsqword(0x60)  as *const u64;
//         let rf_peb: *const PEB = peb_offset as * const PEB;
//         let peb = *rf_peb;

//         let mut p_ldr_data_table_entry: *const LDR_DATA_TABLE_ENTRY = (*peb.Ldr).InMemoryOrderModuleList.Flink as *const LDR_DATA_TABLE_ENTRY;
//         let mut p_list_entry = &(*peb.Ldr).InMemoryOrderModuleList as *const LIST_ENTRY;

//         loop {
//             let buffer = std::slice::from_raw_parts(
//                 (*p_ldr_data_table_entry).FullDllName.Buffer,
//                 (*p_ldr_data_table_entry).FullDllName.Length as usize / 2);
//             let dll_name = String::from_utf16_lossy(buffer);
//             if dll_name.to_lowercase().starts_with(module_name) {
//                 let module_base: HINSTANCE = (*p_ldr_data_table_entry).Reserved2[0] as HINSTANCE;
//                 return module_base;
//             }
//             if p_list_entry == (*peb.Ldr).InMemoryOrderModuleList.Blink {
//                 println!("Module not found!");
//                 return 0;
//             }
//             p_list_entry = (*p_list_entry).Flink;
//             p_ldr_data_table_entry = (*p_list_entry).Flink as *const LDR_DATA_TABLE_ENTRY;
//         }
//     }
// }
