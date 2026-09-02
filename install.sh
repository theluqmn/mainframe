#!/bin/bash

set -e

echo ""
echo "Mainframe Installer"
echo ""
echo "Welcome to a basic CLI-based somewhat interactive installer for Mainframe, a
web OS ahh server management software."
echo ""
echo "Website: https://theluqmn.hackclub.app/mainframe"
echo "Repository: https://github.com/theluqmn/mainframe"

echo ""
echo "Basic configuration"
echo ""

read -p "(1/3) Where should /mainframe be located? (./):      " mainframe_folder
if [ -z "$mainframe_folder" ]; then
    mainframe_folder="./"
fi

read -p "(2/3) Where to create a data folder? (./):           " data_folder
if [ -z "$data_folder" ]; then
    data_folder="./data"
fi

read -p "(3/3) What port should Mainframe listen on? (8080):  " port
if [ -z "$port" ]; then
    port="8080"
fi

echo ""
echo "Web configuration"
echo ""
echo "This will be required to access the web OS. The address is either
the domain that you have pointed to the server, or the static public IPv4 address"
read -p "Enter an address (heheheeee.com): " hostname
if [ -z "$hostname" ]; then
    hostname="127.0.0.1"
fi

echo ""
echo "Performing installation..."
mkdir -p "$mainframe_folder"
cd "$mainframe_folder"

curl -L -o mainframe https://github.com/theluqmn/mainframe/releases/latest/download/mainframe-linux
chmod +x mainframe

mkdir -p "../$data_folder"
echo "{\"port\": $port, \"data\": \"$data_folder\", \"hostname\": \"$hostname\"}" > config.json

echo ""
echo "Installation complete!"
echo "You can now run mainframe by navigating to '$mainframe_folder' and running './mainframe'."