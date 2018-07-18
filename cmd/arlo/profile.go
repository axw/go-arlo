package main

import (
	"context"

	"github.com/spf13/cobra"
)

var cmdProfile = &cobra.Command{
	Use:   "profile",
	Short: "Display your profile information",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		profile, err := arloClient.Profile(context.Background())
		if err != nil {
			return err
		}
		return dumpJSON(profile)
	},
}
