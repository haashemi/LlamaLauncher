package egl

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type FortniteVersion struct {
	CLN        string `json:"cln"`
	Build      string `json:"build"`
	Branch     string `json:"branch"`
	Version    string `json:"version"`
	ModuleName string `json:"moduleName"`
}

func GetFortniteVersion() (*FortniteVersion, error) {
	resp, err := http.Get("https://fortnite-public-service-prod11.ol.epicgames.com/fortnite/api/version")
	if err != nil {
		return nil, err
	} else if resp.StatusCode != http.StatusOK {
		return nil, FindError(resp.Body)
	}
	defer resp.Body.Close()

	raw, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	data := &FortniteVersion{}
	return data, json.Unmarshal(raw, data)
}

func (data *FortniteVersion) String(platform Client) string {
	return "Fortnite/++Fortnite+Release-" + data.Version + "-CL-" + data.CLN + "-" + platform.OS + " " + platform.OS + "/" + platform.OSVersion
}
