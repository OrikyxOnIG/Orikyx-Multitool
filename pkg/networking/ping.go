package networking

import (
	"fmt"
	"os/exec"
)

func Ping(target string) {
	cmd := exec.Command("ping", "-c", "4", target)  // For Linux or Mac, adjust for Windows as needed
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println(string(output))
}
