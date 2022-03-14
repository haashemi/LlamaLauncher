package cmd

import (
	"llamalauncher/pkg/system"
	"os/exec"
	"time"
)

func launchGame(gamePath string, antiCheat string, launchArgs, dllPaths []string) {
	launcherExe := gamePath + "\\FortniteGame\\Binaries\\Win64\\FortniteLauncher.exe"
	shippingExe := gamePath + "\\FortniteGame\\Binaries\\Win64\\FortniteClient-Win64-Shipping.exe"
	antiCheatShippingExe := gamePath + "\\FortniteGame\\Binaries\\Win64\\FortniteClient-Win64-Shipping_"+antiCheat+".exe"

	log.Info("Running Fortnite Launcher...")
	fnLauncher := exec.Command(launcherExe, launchArgs...)
	checkErr(fnLauncher.Start())
	system.SuspendProcess(uint32(fnLauncher.Process.Pid))

	log.Info("Running Fortnite Anti-Cheat Process...")
	fnAntiCheatProcess := exec.Command(antiCheatShippingExe, launchArgs...)
	checkErr(fnAntiCheatProcess.Start())
	system.SuspendProcess(uint32(fnAntiCheatProcess.Process.Pid))

	log.Info("Running Fortnite Process...")
	fnProcess := exec.Command(shippingExe, launchArgs...)
	checkErr(fnProcess.Start())
	system.SuspendProcess(uint32(fnProcess.Process.Pid))

	time.Sleep(time.Second)
	if len(dllPaths) > 0 {
		log.Info("Injecting Dlls...")
		for _, dllPath := range dllPaths {
			checkErr(system.InjectDLL(uint32(fnProcess.Process.Pid), dllPath))
			log.Info(dllPath + " injected!")
		}
		system.ResumeProcess(uint32(fnProcess.Process.Pid))
		// log.Info("All Dlls injected!")
	}

	// If the game is closed kill the launcher and the eac shipping
	fnProcess.Wait()
	log.Warn("Fortnite Process CLOSED")
	fnLauncher.Process.Kill()
	log.Warn("Fortnite Launcher CLOSED")
	fnAntiCheatProcess.Process.Kill()
	log.Warn("Fortnite Anti-Cheat CLOSED")
}
