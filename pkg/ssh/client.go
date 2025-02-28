package ssh

import (
	"fmt"
	"golang.org/x/crypto/ssh"
)

func Connect(target string) {
	// SSH connection logic (e.g., user@host)
	client, err := ssh.Dial("tcp", target, &ssh.ClientConfig{
		User: "your-username",
		Auth: []ssh.AuthMethod{ssh.Password("your-password")},
	})
	if err != nil {
		fmt.Printf("Failed to connect: %s\n", err)
		return
	}
	defer client.Close()

	// Execute a command after connection, for example
	session, err := client.NewSession()
	if err != nil {
		fmt.Printf("Failed to create session: %s\n", err)
		return
	}
	defer session.Close()

	output, err := session.CombinedOutput("uptime")
	if err != nil {
		fmt.Printf("Failed to run command: %s\n", err)
		return
	}
	fmt.Println(string(output))
}
