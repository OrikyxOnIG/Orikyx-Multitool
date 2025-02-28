package cmd

import (
	"github.com/spf13/cobra"
	"orikyx-multitool/pkg/ssh"
)

var SSHCmd = &cobra.Command{
	Use:   "ssh",
	Short: "SSH client commands",
}

var connectCmd = &cobra.Command{
	Use:   "connect [user]@[host]",
	Short: "Connect to an SSH server",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cmd.Help()
			return
		}
		ssh.Connect(args[0])  // Call SSH connection logic
	},
}

func init() {
	SSHCmd.AddCommand(connectCmd)  // Add the connect command to SSH commands
}
