/*
Copyright © 2023 Miguel Angel Ajo Pelayo <majopela@redhat.com
*/
package cmd

import (
	"github.com/redhat-et/jumpstarter/pkg/harness"
	"github.com/spf13/cobra"
)

// powerCmd represents the listDevices command
var detachStorage = &cobra.Command{
	Use:   "detach-storage",
	Short: "Detaches storage from the device",

	Run: func(cmd *cobra.Command, args []string) {
		if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
			handleErrorAsFatal(err)
		}

		driver := cmd.Flag("driver").Value.String()
		device, err := harness.FindDevice(driver, args[0])
		handleErrorAsFatal(err)

		err = device.AttachStorage(false)
		handleErrorAsFatal(err)

	},
}

func init() {
	rootCmd.AddCommand(detachStorage)
	detachStorage.Flags().StringP("driver", "d", "", "Only list devices for the specified driver")
}