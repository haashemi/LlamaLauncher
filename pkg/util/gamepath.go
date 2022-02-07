package util

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

const EGLProgramDataPath = "C:\\ProgramData\\Epic\\UnrealEngineLauncher\\LauncherInstalled.dat"

var ErrGameNotFound = errors.New("game not found. make sure to install game from epic-games launcher")

type LauncherInstalled struct {
	InstallationList []struct {
		InstallLocation string `json:"InstallLocation"`
		NamespaceID     string `json:"NamespaceId"`
		ItemID          string `json:"ItemId"`
		ArtifactID      string `json:"ArtifactId"`
		AppVersion      string `json:"AppVersion"`
		AppName         string `json:"AppName"`
	} `json:"InstallationList"`
}

func FindGamePath() (string, error) {
	f, err := ioutil.ReadFile(EGLProgramDataPath)
	if err != nil {
		return "", err
	}

	data := LauncherInstalled{}
	if err = json.Unmarshal(f, &data); err != nil {
		return "", err
	}

	for _, game := range data.InstallationList {
		if game.NamespaceID == "fn" {
			return game.InstallLocation, nil
		}
	}
	return "", ErrGameNotFound
}
