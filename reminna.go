package main

import (
	"fmt"
	"os"

	"golang.org/x/crypto/ssh"
)

func main() {
	// Remote server details
	host := "192.168.2.34"
	port := 2290
	username := "mrs"
	password := "your_password" // Replace with your actual password

	// Local and remote directories
	localDir := "/path/to/local/directory"
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

	// Change to the remote directory
	err = sftpClient.Chdir(remoteDir)
	if err != nil {
		fmt.Println("Failed to change directory:", err)
		return
	}

	// List remote files
	files, err := sftpClient.ReadDir(".")
	if err != nil {
		fmt.Println("Failed to list files:", err)
		return
	}

	for _, file := range files {
		// Download each file
		remoteFilePath := file.Name()
		localFilePath := localDir + "/" + remoteFilePath

		srcFile, err := sftpClient.Open(remoteFilePath)
		if err != nil {
			fmt.Println("Failed to open remote file:", err)
			return
		}
		defer srcFile.Close()

		dstFile, err := os.Create(localFilePath)
		if err != nil {
			fmt.Println("Failed to create local file:", err)
			return
		}
		defer dstFile.Close()

		_, err = srcFile.WriteTo(dstFile)
		if err != nil {
			fmt.Println("Failed to download file:", err)
			return
		}
		fmt.Println("Downloaded:", remoteFilePath)
	}
	fmt.Println("SFTP download completed.")
}
