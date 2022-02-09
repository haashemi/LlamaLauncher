package system

import (
	"golang.org/x/sys/windows"
)

func EnableWindowsTerminalSequences() error {
	stdoutHandle, err := windows.GetStdHandle(uint32(windows.STD_OUTPUT_HANDLE))
	if err != nil {
		return err
	}

	var terminalMode uint32
	err = windows.GetConsoleMode(stdoutHandle, &terminalMode)
	if err != nil {
		return err
	}
	terminalMode |= windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING
	return windows.SetConsoleMode(stdoutHandle, terminalMode)
}
