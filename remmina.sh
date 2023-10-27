#!/bin/bash

# Remote server details
host="192.168.2.34"
port="22"
username="mrs"
password="mr$%0921202571"

# Local and remote directories
#local_dir="/path/to/local/directory"
remote_dir=".local/share/remmina/"

# SFTP connection
sftp -oPort=$port $username@$host <<EOF
  cd $remote_dir
  get *  # You can use 'get' to download files instead of 'put'
  bye
EOF
