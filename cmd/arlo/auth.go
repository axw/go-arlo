package main

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
)

var cmdLogin = &cobra.Command{
	Use:   "login [email]",
	Short: "Login to the Netgear Arlo service",
	Args:  cobra.RangeArgs(0, 1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var email string
		if len(args) == 1 {
			email = args[0]
		}
		return login(email)
	},
}

var cmdLogout = &cobra.Command{
	Use:   "logout",
	Short: "Logout of the Netgear Arlo service",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		return arloClient.Logout(context.Background())
	},
}

func login(email string) error {
	oldState, err := terminal.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		return err
	}
	defer terminal.Restore(int(os.Stdin.Fd()), oldState)
	term := terminal.NewTerminal(os.Stdin, "")

	for email == "" {
		fmt.Fprint(term, "Email: ")
		email, err = term.ReadLine()
		if err != nil {
			return err
		}
	}
	var password string
	for password == "" {
		password, err = term.ReadPassword("Password for " + email + ": ")
		if err != nil {
			return err
		}
	}
	return arloClient.Login(context.Background(), email, password)
}

func promptAuth() (email, password string, err error) {
	oldState, err := terminal.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		return "", "", err
	}
	defer terminal.Restore(int(os.Stdin.Fd()), oldState)

	term := terminal.NewTerminal(os.Stdin, "")
	for email == "" {
		fmt.Fprint(term, "Email: ")
		email, err = term.ReadLine()
		if err != nil {
			return "", "", err
		}
	}
	for password == "" {
		password, err = term.ReadPassword("Password for " + email + ": ")
		if err != nil {
			return "", "", err
		}
	}
	return email, password, nil
}
