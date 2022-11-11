package cmd

import (
	"fmt"
	"llamalauncher/pkg/database"
	"llamalauncher/pkg/egl"
	"llamalauncher/pkg/system"
	"llamalauncher/pkg/util"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/LlamaNite/llamalog"
)

var UserAgent = "LlamaLauncher/v0.1.1-dev"
var log = llamalog.NewLogger("LlamaLauncher")
var antiCheatRegex = regexp.MustCompile(`[A-Z]`)

func init() {
	checkErr(database.AutoMigrateDB("./config/db.sqlite3"))
}

func Run() {
	checkErr(system.EnableWindowsTerminalSequences())

	account, err := selectAccount()
	checkErr(err)

	fortniteVersion, err := egl.GetFortniteVersion()
	checkErr(err)
	UserAgent = fortniteVersion.String(egl.ClientLauncherApp2)

	exchangeCode, err := egl.GetExchange(account.AccessToken, egl.ClientLauncherApp2, UserAgent)
	checkErr(err)

	caledra, err := egl.GetCaldera(account.AccountID, exchangeCode.Code)
	checkErr(err)

	gamePath, err := util.FindGamePath()
	checkErr(err)

	launchCommand, err := util.GetLaunchCommand(gamePath)
	checkErr(err)

	antiCheat := strings.ToLower(strings.Join(antiCheatRegex.FindAllString(caledra.Provider, -1), ""))
	var noAntiCheat string
	switch antiCheat {
	case "eac":
		noAntiCheat = "-nobe"
	case "be":
		noAntiCheat = "-noeac"
	}

	launchArgs := append(strings.Split(launchCommand[1:], " "),
		"-AUTH_LOGIN=unused",
		"-AUTH_PASSWORD="+exchangeCode.Code,
		"-AUTH_TYPE=exchangecode",
		"-epicapp=Fortnite",
		"-epicenv=Prod",
		"-EpicPortal",
		"-epicusername="+account.DisplayName,
		"-epicuserid="+account.AccountID,
		"-epiclocale=en",
		"-epicsandboxid=fn",
		noAntiCheat,
		"-fromfl="+antiCheat,
		"-caldera="+caledra.Jwt,
	)

	// fmt.Println(launchArgs)
	clearCMD()
	fmt.Println(llamalog.Magenta(">> LlamaLauncher >> Part of LlamaNite projects"))
	fmt.Println(llamalog.Magenta(">> LlamaNite projects >> https://llamanite.com"))
	fmt.Println(llamalog.Magenta(">> LlamaLauncher source code >> https://github.com/MR-AliHaashemi/LlamaLauncher\n"))
	fmt.Println(llamalog.Magenta(">> STARTING FORTNITE <<"))
	launchGame(gamePath, strings.ToUpper(antiCheat), launchArgs, []string{".\\config\\PlataniumV2.dll"})
}

func checkErr(err error) {
	if err != nil {
		log.Error(err.Error())
		log.Warn("LlamaLauncher will exit in 5 seconds...")
		time.Sleep(time.Second * 5)
		os.Exit(1)
	}
}

func clearCMD() {
	// \033[H  => change cursor position to default
	// \x1b[2J => Blank screen
	fmt.Println("\033[H\x1b[2J")
}
