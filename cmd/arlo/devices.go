package main

import (
	"context"

	"github.com/spf13/cobra"
)

var cmdDevices = &cobra.Command{
	Use:   "devices",
	Short: "Display registered devices",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		devices, err := arloClient.Devices(context.Background())
		if err != nil {
			return err
		}
		return dumpJSON(devices)
	},
}
