package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"orikyx-multitool/pkg/banner"
	"orikyx-multitool/pkg/networking"
	"orikyx-multitool/pkg/ssh"
)

var rootCmd = &cobra.Command{
	Use:   "orikyx-multitool",
	Short: "A powerful multitool for networking, SSH, and more",
	Run: func(cmd *cobra.Command, args []string) {
		banner.PrintBanner()  // Display the banner when the app starts
		fmt.Println("Type --help for available commands")
	},
}

func init() {
	rootCmd.AddCommand(networking.NetworkCmd)  // Add network commands to CLI
	rootCmd.AddCommand(ssh.SSHCmd)             // Add SSH client commands to CLI
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
