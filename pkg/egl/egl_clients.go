package egl

import (
	"encoding/base64"
	"net/url"
)

type Client struct{ ClientID, Secret, OS, OSVersion string }

var (
	ClientPC           Client = Client{"ec684b8c687f479fadea3cb2ad83f5c6", "e1f31c211f28413186262d37a13fc84d", "Windows", "10.0.22000.1.256.64bit"}
	ClientLauncherApp2 Client = Client{"34a02cf8f4414e29b15921876da36f9a", "daafbccc737745039dffe53d94fc76cf", "Windows", "10.0.22000.1.256.64bit"}
)

func (c *Client) Encode() string {
	return base64.StdEncoding.EncodeToString([]byte(c.ClientID + ":" + c.Secret))
}

func (c *Client) AuthCodeURL() string {
	return `https://www.epicgames.com/id/api/redirect?clientId=` + c.ClientID + `&responseType=code`
}

func (c *Client) AuthCodeLoginURL() string {
	return "https://www.epicgames.com/id/login?redirectUrl=" + url.QueryEscape(c.AuthCodeURL())
}
