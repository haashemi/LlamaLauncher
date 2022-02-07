package injector

import (
	"errors"
	"os"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

const (
	VirtualAllocMemCommit  = 0x00001000
	VirtualAllocMemReserve = 0x00002000
	PageReadWrite          = 0x04
	SizeOfChar             = 4
)

var (
	ntdllModule          = windows.NewLazySystemDLL("ntdll.dll")
	procNtResumeProcess  = ntdllModule.NewProc("NtResumeProcess")
	procNtSuspendProcess = ntdllModule.NewProc("NtSuspendProcess")

	kernel32Module         = windows.NewLazySystemDLL("kernel32.dll")
	procGetModuleHandle    = kernel32Module.NewProc("GetModuleHandleW")
	procVirtualAllocEx     = kernel32Module.NewProc("VirtualAllocEx")
	procCreateRemoteThread = kernel32Module.NewProc("CreateRemoteThread")
)

/*
HMODULE GetModuleHandleW(
    [in, optional] LPCWSTR lpModuleName
);
*/
func GetModuleHandle(lpModuleName string) (windows.Handle, error) {
	ptr, err := syscall.UTF16PtrFromString(lpModuleName)
	if err != nil {
		return 0, err
	}

	ret, _, _ := syscall.Syscall(procGetModuleHandle.Addr(), 1, uintptr(unsafe.Pointer(ptr)), 0, 0)

	return windows.Handle(ret), nil
}

/*
Original Function:
LPVOID VirtualAllocEx(
    [in]           HANDLE hProcess,
    [in, optional] LPVOID lpAddress,
    [in]           SIZE_T dwSize,
    [in]           DWORD  flAllocationType,
    [in]           DWORD  flProtect
);
*/
func VirtualAllocEx(hProcess windows.Handle, lpAddress uintptr, dwSize uint, flAllocationType, flProtect uint32) uintptr {
	ret, _, _ := procVirtualAllocEx.Call(
		uintptr(hProcess),
		lpAddress,
		uintptr(dwSize),
		uintptr(flAllocationType),
		uintptr(flProtect),
	)

	return ret
}

/*
Original Function:
HANDLE CreateRemoteThread(
    [in]  HANDLE                 hProcess,
    [in]  LPSECURITY_ATTRIBUTES  lpThreadAttributes,
    [in]  SIZE_T                 dwStackSize,
    [in]  LPTHREAD_START_ROUTINE lpStartAddress,
    [in]  LPVOID                 lpParameter,
    [in]  DWORD                  dwCreationFlags,
    [out] LPDWORD                lpThreadId
);
*/
func CreateRemoteThread(hProcess windows.Handle, lpThreadAttributes uintptr, dwStackSize uint, lpStartAddress, lpParameter uintptr, dwCreationFlags uint32, lpThreadId *uintptr) uintptr {
	ret, _, _ := procCreateRemoteThread.Call(
		uintptr(hProcess),
		lpThreadAttributes,
		uintptr(dwStackSize),
		lpStartAddress,
		lpParameter,
		uintptr(dwCreationFlags),
		uintptr(unsafe.Pointer(lpThreadId)),
	)

	return ret
}

func InjectDLL(pid uint32, dll string) error {
	if f, err := os.Stat(dll); err != nil || f.IsDir() {
		return errors.New("invalid DLL file")
	}

	processHandle, err := windows.OpenProcess(
		windows.PROCESS_QUERY_INFORMATION|
			windows.PROCESS_CREATE_THREAD|
			windows.PROCESS_VM_OPERATION|
			windows.PROCESS_VM_WRITE|
			windows.PROCESS_VM_READ,
		false, pid,
	)
	if err != nil {
		return err
	}
	defer windows.Close(processHandle)

	kernel32Handle, err := GetModuleHandle("kernel32.dll")
	if err != nil {
		return err
	}

	loadLibraryA, err := windows.GetProcAddress(kernel32Handle, "LoadLibraryA")
	if err != nil {
		return err
	}

	size := (uint)((len(dll) + 1) * SizeOfChar)
	memAddress := VirtualAllocEx(processHandle, uintptr(0), size, VirtualAllocMemCommit|VirtualAllocMemReserve, PageReadWrite)
	ptr, err := syscall.BytePtrFromString(dll)
	if err != nil {
		return err
	}

	var bytesWritten uintptr
	windows.WriteProcessMemory(processHandle, memAddress, ptr, uintptr(size), &bytesWritten)
	// fmt.Println("written", bytesWritten)

	var lpThreadId uintptr
	CreateRemoteThread(processHandle, uintptr(0), 0, loadLibraryA, memAddress, 0, &lpThreadId)
	// fmt.Println("created thread", lpThreadId)
	return nil
}
