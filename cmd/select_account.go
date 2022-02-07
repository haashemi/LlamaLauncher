package cmd

import (
	"context"
	"errors"
	"fmt"
	"llamalauncher/pkg/database"
	"llamalauncher/pkg/database/ent"
	"llamalauncher/pkg/database/ent/user"
	"llamalauncher/pkg/egl"
	"strconv"

	"github.com/LlamaNite/llamalog"
)

func selectAccount() (*ent.User, error) {
	clearCMD()

	fmt.Println(llamalog.Magenta(">> LlamaLauncher >> Part of LlamaNite projects"))
	fmt.Println(llamalog.Magenta(">> Loading Accounts..."))

	accounts, err := getRefreshedAccounts()
	if err != nil {
		return nil, err
	}

	clearCMD()
	accountsList := map[string]*ent.User{}
	fmt.Println(llamalog.Green("Your Accounts in LlamaLauncher"))
	fmt.Println("    0) Add New Account")
	for index, acc := range accounts {
		fmt.Printf("    %d) %s\n", index+1, acc.DisplayName)
		accountsList[strconv.Itoa(index+1)] = acc
	}

	accRange := fmt.Sprintf("(0 - %d)", len(accounts))
	if len(accounts) == 0 {
		accRange = "(0)"
	}
	fmt.Printf("\n%s %s %s ", llamalog.Magenta("Select your account"), accRange, llamalog.Magenta("=>"))

	var index string
	fmt.Scanln(&index)

	if index == "0" {
		return addNewAccount()
	} else if acc, ok := accountsList[index]; ok {
		return acc, nil
	}

	return nil, errors.New("index out of range")
}

func getRefreshedAccounts() ([]*ent.User, error) {
	db, err := database.GetConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	accounts, err := db.User.Query().All(context.Background())
	if err != nil {
		return nil, err
	}

	for _, acc := range accounts {
		newData, err := egl.OAuthFromRefreshToken(acc.RefreshToken, egl.ClientLauncherApp2, UserAgent)
		if err != nil {
			_, err = db.User.Delete().Where(user.AccountID(acc.AccountID)).Exec(context.Background())
			if err != nil {
				return nil, err
			}
			continue
		}

		err = db.User.Update().Where(user.AccountID(acc.AccountID)).
			SetDisplayName(newData.DisplayName).
			SetAccessToken(newData.AccessToken).
			SetAccessTokenExpiresAt(newData.ExpiresAt).
			SetRefreshToken(newData.RefreshToken).
			SetRefresTokenExpiresAt(newData.RefreshExpiresAt).
			Exec(context.Background())
		if err != nil {
			return nil, err
		}
	}

	accounts, err = db.User.Query().All(context.Background())
	if err != nil {
		return nil, err
	}

	return accounts, nil
}
