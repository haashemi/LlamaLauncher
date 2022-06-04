package cmd

import (
	"context"
	"fmt"
	"llamalauncher/pkg/database"
	"llamalauncher/pkg/database/ent"
	"llamalauncher/pkg/egl"

	"github.com/LlamaNite/llamalog"
)

func addNewAccount() (*ent.User, error) {
	clearCMD()
	log.Info("Open one of these URLs and enter the code here")
	fmt.Printf("   => %s\n   => %s\n\n", egl.ClientLauncherApp2.AuthCodeURL(), egl.ClientLauncherApp2.AuthCodeLoginURL())
	fmt.Print(llamalog.Magenta("Enter AuthCode => "))

	var authCode string
	fmt.Scanln(&authCode)

	data, err := egl.OAuthFromAuthorizationCode(authCode, egl.ClientLauncherApp2, UserAgent)
	if err != nil {
		return nil, err
	}

	db, err := database.GetConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	return db.User.Create().
		SetAccountID(data.AccountID).
		SetDisplayName(data.DisplayName).
		SetAccessToken(data.AccessToken).
		SetAccessTokenExpiresAt(data.ExpiresAt).
		SetRefreshToken(data.RefreshToken).
		SetRefreshTokenExpiresAt(data.RefreshExpiresAt).
		Save(context.Background())
}
