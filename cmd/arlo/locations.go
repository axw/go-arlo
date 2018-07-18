package main

import (
	"context"

	"github.com/spf13/cobra"
)

var cmdLocations = &cobra.Command{
	Use:   "locations",
	Short: "Display registered locations",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		locations, err := arloClient.Locations(context.Background())
		if err != nil {
			return err
		}
		return dumpJSON(locations)
	},
}
