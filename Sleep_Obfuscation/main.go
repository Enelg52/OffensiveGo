package main

import (
	"bufio"
	"debug/pe"
	"fmt"
	"os"
	"unsafe"

	"golang.org/x/sys/windows"
)

type IMAGE_DOS_HEADER struct {
	E_magic    uint16
	E_cblp     uint16
	E_cp       uint16
	E_crlc     uint16
	E_cparhdr  uint16
	E_minalloc uint16
	E_maxalloc uint16
	E_ss       uint16
	E_sp       uint16
	E_csum     uint16
	E_ip       uint16
	E_cs       uint16
	E_lfarlc   uint16
	E_ovno     uint16
	E_res      [4]uint16
	E_oemid    uint16
	E_oeminfo  uint16
	E_res2     [10]uint16
	E_lfanew   int32
}

type IMAGE_NT_HEADERS64 struct {
	Signature               uint32
	IMAGE_FILE_HEADER       pe.FileHeader
	IMAGE_OPTIONAL_HEADER64 pe.OptionalHeader64
}

// https://docs.microsoft.com/en-us/windows/win32/api/winnt/ns-winnt-context
type M128A struct {
	Low  uint64
	High int64
}

// https://docs.microsoft.com/en-us/windows/win32/api/winnt/ns-winnt-context
type XMM_SAVE_AREA32 struct {
	ControlWord    uint16
	StatusWord     uint16
	TagWord        byte
	Reserved1      byte
	ErrorOpcode    uint16
	ErrorOffset    uint32
	ErrorSelector  uint16
	Reserved2      uint16
	DataOffset     uint32
	DataSelector   uint16
	Reserved3      uint16
	MxCsr          uint32
	MxCsr_Mask     uint32
	FloatRegisters [8]M128A
	XmmRegisters   [256]byte
	Reserved4      [96]byte
}

// https://docs.microsoft.com/en-us/windows/win32/api/winnt/ns-winnt-context
type CONTEXT struct {
	P1Home uint64
	P2Home uint64
	P3Home uint64
	P4Home uint64
	P5Home uint64
	P6Home uint64

	ContextFlags uint32
	MxCsr        uint32

	SegCs  uint16
	SegDs  uint16
	SegEs  uint16
	SegFs  uint16
	SegGs  uint16
	SegSs  uint16
	EFlags uint32

	Dr0 uint64
	Dr1 uint64
	Dr2 uint64
	Dr3 uint64
	Dr6 uint64
	Dr7 uint64

	Rax uint64
	Rcx uint64
	Rdx uint64
	Rbx uint64
	Rsp uint64
	Rbp uint64
	Rsi uint64
	Rdi uint64
	R8  uint64
	R9  uint64
	R10 uint64
	R11 uint64
	R12 uint64
	R13 uint64
	R14 uint64
	R15 uint64

	Rip uint64

	FltSave XMM_SAVE_AREA32

	VectorRegister [26]M128A
	VectorControl  uint64

	DebugControl         uint64
	LastBranchToRip      uint64
	LastBranchFromRip    uint64
	LastExceptionToRip   uint64
	LastExceptionFromRip uint64
}

type Ustring struct {
	Length        uint32
	MaximumLength uint32
	Buffer        []byte
}

const (
	WT_EXECUTEINTIMERTHREAD = 0x00000020
)

var (
	nt                    = windows.NewLazyDLL("ntdll.dll")
	k32                   = windows.NewLazyDLL("kernel32.dll")
	adv                   = windows.NewLazyDLL("advapi32.dll")
	createTimerQueue      = k32.NewProc("CreateTimerQueue")
	createTimerQueueTimer = k32.NewProc("CreateTimerQueueTimer")
	deleteTimerQueue      = k32.NewProc("DeleteTimerQueue")
	virtualProtect        = k32.NewProc("VirtualProtect").Addr()
	setEvent              = k32.NewProc("SetEvent").Addr()
	waitForSingleObject   = k32.NewProc("WaitForSingleObject").Addr()
	ntContinue            = nt.NewProc("NtContinue").Addr()
	rtlCaptureContext     = nt.NewProc("RtlCaptureContext").Addr()
	systemFunction032     = adv.NewProc("SystemFunction032").Addr()
)

func main() {

	for {
		ObfuscateAndSleep(4 * 1000)
	}
}

func ObfuscateAndSleep(sleep uint64) {

	input := bufio.NewScanner(os.Stdin)
	input.Scan()

	var timerQueue uintptr

	var imageBase windows.Handle
	err := windows.GetModuleHandleEx(0, nil, &imageBase)
	if err != nil {
		panic(err)
	}

	imageSize := *&(*IMAGE_NT_HEADERS64)(unsafe.Pointer(uintptr(*&(*IMAGE_DOS_HEADER)(unsafe.Pointer(imageBase)).E_lfanew) + uintptr(imageBase))).IMAGE_OPTIONAL_HEADER64.SizeOfImage
	key := Ustring{Length: 16, MaximumLength: 16, Buffer: []byte{0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55}}
	img := Ustring{Length: imageSize, MaximumLength: imageSize, Buffer: unsafe.Slice((*byte)(unsafe.Pointer(imageBase)), imageSize)}

	eventHandle, err := windows.CreateEvent(nil, 0, 0, nil)
	if err != nil {
		panic(err)
	}

	ctxStruct := CONTEXT{}
	// ctxStruct := make([]uint8, 1232)

	var newTimer windows.Handle

	timerQueue, _, err = createTimerQueue.Call()
	if timerQueue == 0 {
		panic(err)
	}

	r1, _, _ := createTimerQueueTimer.Call(uintptr(unsafe.Pointer(&newTimer)), timerQueue, rtlCaptureContext, uintptr(unsafe.Pointer(&ctxStruct)), 0, 0, WT_EXECUTEINTIMERTHREAD)
	if r1 == 0 {
		_, _, err := deleteTimerQueue.Call(timerQueue)
		if err.Error() != "The operation completed successfully." {
			panic(err)
		}
		return
	}

	windows.WaitForSingleObject(windows.Handle(eventHandle), 0x32)

	ContextVirtualProtect := ctxStruct
	ContextWaitForSingleObject := ctxStruct
	ContextSystemFunction032 := ctxStruct
	ContextSystemFunction032D := ctxStruct
	ContextVirtualProtectX := ctxStruct
	ContextSetEvent := ctxStruct

	var oldProtect uint32

	ContextVirtualProtect.Rsp -= 8
	ContextVirtualProtect.Rip = *(*uint64)(unsafe.Pointer(&virtualProtect))
	ContextVirtualProtect.Rcx = *(*uint64)(unsafe.Pointer(&imageBase))
	ContextVirtualProtect.Rdx = uint64(imageSize)
	ContextVirtualProtect.R8 = windows.PAGE_READWRITE
	ContextVirtualProtect.R9 = uint64(uintptr(unsafe.Pointer(&oldProtect)))

	ContextSystemFunction032.Rsp -= 8
	ContextSystemFunction032.Rip = *(*uint64)(unsafe.Pointer(&systemFunction032))
	ContextSystemFunction032.Rcx = uint64(uintptr(unsafe.Pointer(&img)))
	ContextSystemFunction032.Rdx = uint64(uintptr(unsafe.Pointer(&key)))

	ContextWaitForSingleObject.Rsp -= 8
	ContextWaitForSingleObject.Rip = *(*uint64)(unsafe.Pointer(&waitForSingleObject))
	ContextWaitForSingleObject.Rcx = uint64(uintptr(unsafe.Pointer(windows.CurrentProcess())))
	ContextWaitForSingleObject.Rdx = sleep

	ContextSystemFunction032D.Rsp -= 8
	ContextSystemFunction032D.Rip = *(*uint64)(unsafe.Pointer(&systemFunction032))

	ContextSystemFunction032D.Rcx = uint64(uintptr(unsafe.Pointer(&img)))
	ContextSystemFunction032D.Rdx = uint64(uintptr(unsafe.Pointer(&key)))

	ContextVirtualProtectX.Rsp -= 8
	ContextVirtualProtectX.Rip = *(*uint64)(unsafe.Pointer(&virtualProtect))
	ContextVirtualProtectX.Rcx = *(*uint64)(unsafe.Pointer(&imageBase))
	ContextVirtualProtectX.Rdx = uint64(imageSize)
	ContextVirtualProtectX.R8 = windows.PAGE_EXECUTE_READWRITE
	ContextVirtualProtectX.R9 = uint64(uintptr(unsafe.Pointer(&oldProtect)))

	ContextSetEvent.Rsp -= 8
	ContextSetEvent.Rip = *(*uint64)(unsafe.Pointer((&setEvent)))
	ContextSetEvent.Rcx = *(*uint64)(unsafe.Pointer((&eventHandle)))

	r1, _, err = createTimerQueueTimer.Call(uintptr(unsafe.Pointer(&newTimer)), timerQueue, ntContinue, uintptr(unsafe.Pointer(&ContextVirtualProtect)), 100, 0, WT_EXECUTEINTIMERTHREAD)
	if r1 == 0 {
		panic(err)
	}

	r1, _, err = createTimerQueueTimer.Call(uintptr(unsafe.Pointer(&newTimer)), timerQueue, ntContinue, uintptr(unsafe.Pointer(&ContextSystemFunction032)), 200, 0, WT_EXECUTEINTIMERTHREAD)
	if r1 == 0 {
		panic(err)
	}
	r1, _, err = createTimerQueueTimer.Call(uintptr(unsafe.Pointer(&newTimer)), timerQueue, ntContinue, uintptr(unsafe.Pointer(&ContextWaitForSingleObject)), 300, 0, WT_EXECUTEINTIMERTHREAD)
	if r1 == 0 {
		panic(err)
	}
	r1, _, err = createTimerQueueTimer.Call(uintptr(unsafe.Pointer(&newTimer)), timerQueue, ntContinue, uintptr(unsafe.Pointer(&ContextSystemFunction032D)), 400, 0, WT_EXECUTEINTIMERTHREAD)
	if r1 == 0 {
		panic(err)
	}
	r1, _, err = createTimerQueueTimer.Call(uintptr(unsafe.Pointer(&newTimer)), timerQueue, ntContinue, uintptr(unsafe.Pointer(&ContextVirtualProtectX)), 500, 0, WT_EXECUTEINTIMERTHREAD)
	if r1 == 0 {
		panic(err)
	}
	r1, _, err = createTimerQueueTimer.Call(uintptr(unsafe.Pointer(&newTimer)), timerQueue, ntContinue, uintptr(unsafe.Pointer(&ContextSetEvent)), 600, 0, WT_EXECUTEINTIMERTHREAD)
	if r1 == 0 {
		panic(err)
	}

	_, err = windows.WaitForSingleObject(windows.Handle(eventHandle), windows.INFINITE)
	if err != nil {
		panic(err)
	}

	fmt.Println("End Wait")

}
