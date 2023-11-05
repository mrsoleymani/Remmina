package main

import (
	"fmt"
	"os"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

func main() {
	// Remote server details
	host := "192.168.2.34"
	port := 2290
	username := "mrs"
	password := "your_password" // Replace with your actual password

	// Local and remote directories
	// localDir := "/path/to/local/directory"
	remoteDir := ".local/share/remmina/"

	// Create SSH client configuration
	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// Connect to the remote server over SSH
	client, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", host, port), config)
	if err != nil {
		fmt.Println("Failed to connect:", err)
		return
	}
	defer client.Close()

	// Open an SFTP session on the SSH connection
	sftpClient, err := sftp.NewClient(client)
	if err != nil {
		fmt.Println("Failed to open SFTP session:", err)
		return
	}
	defer sftpClient.Close()

	// Use sftp.Walk to list files in the remote directory
	err = sftp.wa(sftpClient, remoteDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("Error:", err)
			return nil
		}

		// Print or process the file information
		fmt.Println("File:", path, "Size:", info.Size())

		// Continue with additional processing if needed

		return nil
	})

	if err != nil {
		fmt.Println("Failed to walk remote directory:", err)
		return
	}

	fmt.Println("SFTP directory listing completed.")
}
