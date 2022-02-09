package cmd

import (
	"llamalauncher/pkg/system"
	"os/exec"
	"time"
)

func launchGame(gamePath string, launchArgs, dllPaths []string) {
	launcherExe := gamePath + "\\FortniteGame\\Binaries\\Win64\\FortniteLauncher.exe"
	shippingExe := gamePath + "\\FortniteGame\\Binaries\\Win64\\FortniteClient-Win64-Shipping.exe"
	eacShippingExe := gamePath + "\\FortniteGame\\Binaries\\Win64\\FortniteClient-Win64-Shipping_EAC.exe"

	log.Info("Running Fortnite Launcher...")
	fnLauncher := exec.Command(launcherExe, launchArgs...)
	checkErr(fnLauncher.Start())
	system.SuspendProcess(uint32(fnLauncher.Process.Pid))

	log.Info("Running Fortnite EAC Process...")
	fnEacProcess := exec.Command(eacShippingExe, launchArgs...)
	checkErr(fnEacProcess.Start())
	system.SuspendProcess(uint32(fnEacProcess.Process.Pid))

	log.Info("Running Fortnite Process...")
	fnProcess := exec.Command(shippingExe, launchArgs...)
	checkErr(fnProcess.Start())
	system.SuspendProcess(uint32(fnProcess.Process.Pid))

	time.Sleep(time.Second * 1)
	if len(dllPaths) > 0 {
		log.Info("Injecting Dlls...")
		for _, dllPath := range dllPaths {
			checkErr(system.InjectDLL(uint32(fnProcess.Process.Pid), dllPath))
			log.Info(dllPath + " injected!")
		}
		system.ResumeProcess(uint32(fnProcess.Process.Pid))
		log.Info("All Dlls injected!")
	}

	// If the game is closed kill the launcher and the eac shipping
	fnProcess.Wait()
	log.Warn("fnProcess CLOSED")
	fnLauncher.Process.Kill()
	log.Warn("fnLauncher CLOSED")
	fnEacProcess.Process.Kill()
	log.Warn("fnEacProcess CLOSED")
}
