package cmd

import (
	"github.com/spf13/cobra"
	"orikyx-multitool/pkg/networking"
)

var NetworkCmd = &cobra.Command{
	Use:   "network",
	Short: "Networking related commands",
}

var pingCmd = &cobra.Command{
	Use:   "ping [target]",
	Short: "Ping a target IP or hostname",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cmd.Help()
			return
		}
		networking.Ping(args[0])
	},
}

func init() {
	NetworkCmd.AddCommand(pingCmd)  // Add the ping command to the network command
}
