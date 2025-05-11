package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"golang.org/x/crypto/ssh"
)

func main() {
	// Set up SSH client configuration
	config := &ssh.ClientConfig{
		User: "admin",
		Auth: []ssh.AuthMethod{
			ssh.Password("admin"),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         5 * time.Second,
	}

	// Connect to the remote SSH server
	conn, err := ssh.Dial("tcp", "10.0.0.10:22", config)
	if err != nil {
		log.Fatalf("Failed to dial: %s", err)
	}
	defer conn.Close()

	// Open a new SSH session
	session, err := conn.NewSession()
	if err != nil {
		log.Fatalf("Failed to create session: %s", err)
	}
	defer session.Close()

	// Run the command to export the configuration
	output, err := session.CombinedOutput("ls /")
	if err != nil {
		log.Fatalf("Failed to run command: %s", err)
	}

	// Write the output to a file
	err = os.WriteFile(".\\src\\ssh\\config.txt", output, 0644)
	if err != nil {
		log.Fatalf("Failed to write file: %s", err)
	}

	fmt.Println("Configuration exported successfully!")
}
