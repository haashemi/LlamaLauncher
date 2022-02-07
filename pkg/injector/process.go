package injector

import (
	"fmt"

	"golang.org/x/sys/windows"
)

func SuspendProcess(pid uint32) error {
	handle, err := windows.OpenProcess(windows.PROCESS_SUSPEND_RESUME, false, pid)
	if err != nil {
		return err
	}
	defer windows.CloseHandle(handle)

	if r1, _, _ := procNtSuspendProcess.Call(uintptr(handle)); r1 != 0 {
		return fmt.Errorf("NtStatus='0x%.8X'", r1)
	}
	return nil
}

func ResumeProcess(pid uint32) error {
	handle, err := windows.OpenProcess(windows.PROCESS_SUSPEND_RESUME, false, pid)
	if err != nil {
		return err
	}
	defer windows.CloseHandle(handle)

	r1, _, _ := procNtResumeProcess.Call(uintptr(handle))
	if r1 != 0 {
		return fmt.Errorf("NtStatus='0x%.8X'", r1)
	}
	return nil
}
