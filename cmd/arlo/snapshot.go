package main

import (
	"context"
	"fmt"

	arlo "github.com/axw/go-arlo"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var cmdSnapshot = func() *cobra.Command {
	var refresh bool
	cmd := &cobra.Command{
		Use:   "snapshot [camera-name]",
		Short: "Download a camera snapshot",
		Args:  cobra.RangeArgs(1, 1),
		RunE: func(cmd *cobra.Command, args []string) error {
			// TODO(axw) accept a basestation name for disambiguation?
			ctx := context.Background()
			devices, err := arloClient.Devices(ctx)
			if err != nil {
				return err
			}
			name := args[0]
			matching := make([]arlo.Device, 0, 1)
			for _, device := range devices {
				if device.DeviceType == "camera" && device.DeviceName == name {
					matching = append(matching, device)
				}
			}
			switch len(matching) {
			case 1:
			case 0:
				// TODO(axw) edit-distance
				return errors.Errorf("camera %q not found", name)
			default:
				return errors.Errorf("multiple cameras called %q found", name)
			}

			camera := matching[0]
			if refresh {
				fmt.Println("Refreshing snapshot on", name)
				if err := arloClient.FullFrameSnapshot(ctx, camera); err != nil {
					return err
				}
				// TODO(axw) wait for snapshot to complete
			}
			fmt.Println(camera.PresignedFullFrameSnapshotURL)
			return nil
		},
	}
	cmd.Flags().BoolVar(&refresh, "refresh", true, "Take a new snapshot")
	return cmd
}()
