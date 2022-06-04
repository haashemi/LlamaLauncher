package egl

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type CalderaData struct {
	Jwt      string `json:"jwt"`
	Provider string `json:"provider"`
}

func GetCaldera(accountID, exchangeCode string) (*CalderaData, error) {
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(map[string]interface{}{
		"account_id":    accountID,
		"exchange_code": exchangeCode,
		"test_mode":     false,
		"epic_app":      "fortnite",
		"nvidia":        false,
	})

	req, err := http.NewRequest(http.MethodPost, "https://caldera-service-prod.ecosec.on.epicgames.com/caldera/api/v1/launcher/racp", buf)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Caldera/UNKNOWN-UNKNOWN-UNKNOWN")

	resp, err := http.DefaultClient.Do(req)
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

	data := &CalderaData{}
	return data, json.Unmarshal(raw, data)
}
