package cmd

import (
	"llamalauncher/pkg/system"
	"os/exec"
)

func launchGame(gamePath string, antiCheat string, launchArgs, dllPaths []string) {
	launcherExe := gamePath + "\\FortniteGame\\Binaries\\Win64\\FortniteLauncher.exe"
	shippingExe := gamePath + "\\FortniteGame\\Binaries\\Win64\\FortniteClient-Win64-Shipping.exe"
	antiCheatShippingExe := gamePath + "\\FortniteGame\\Binaries\\Win64\\FortniteClient-Win64-Shipping_" + antiCheat + ".exe"

	log.Info("Running Fortnite Launcher...")
	fnLauncher := exec.Command(launcherExe)
	checkErr(fnLauncher.Start())
	system.SuspendProcess(uint32(fnLauncher.Process.Pid))

	log.Info("Running Fortnite Anti-Cheat Process...")
	fnAntiCheatProcess := exec.Command(antiCheatShippingExe)
	checkErr(fnAntiCheatProcess.Start())
	system.SuspendProcess(uint32(fnAntiCheatProcess.Process.Pid))

	log.Info("Running Fortnite Process...")
	fnProcess := exec.Command(shippingExe, launchArgs...)
	checkErr(fnProcess.Start())
	system.SuspendProcess(uint32(fnProcess.Process.Pid))

	if len(dllPaths) > 0 {
		log.Info("Injecting Dlls...")
		for _, dllPath := range dllPaths {
			checkErr(system.InjectDLL(uint32(fnProcess.Process.Pid), dllPath))
			log.Info(dllPath + " injected!")
		}
		system.ResumeProcess(uint32(fnProcess.Process.Pid))
	}

	// If the game is closed kill the launcher and the eac shipping
	defer fnAntiCheatProcess.Process.Kill()
	defer fnLauncher.Process.Kill()
	fnProcess.Wait()
	log.Warn("Fortnite Process CLOSED")
}
