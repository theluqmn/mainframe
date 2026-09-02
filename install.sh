#!/bin/bash

set -e

echo ""
echo "Mainframe Installer"
echo ""
echo "Website: https://theluqmn.hackclub.app/mainframe"
echo "Repository: https://github.com/theluqmn/mainframe"
echo ""

read -p "Where should /mainframe be located? (./):      " mainframe_folder
if [ -z "$mainframe_folder" ]; then
    mainframe_folder="./"
fi

read -p "Where to create a data folder? (./):           " data_folder
if [ -z "$data_folder" ]; then
    data_folder="./data"
fi

read -p "What port should Mainframe listen on? (8080):  " port
if [ -z "$port" ]; then
    port="8080"
fi

echo ""
echo "Performing installation..."
# mkdir -p "$mainframe_folder"
# cd "$mainframe_folder"

# curl -L -o mainframe https://github.com/theluqmn/mainframe/releases/latest/download/mainframe-linux
# chmod +x mainframe

# mkdir -p "../$data_folder"
# echo "{\"port\": $port, \"data\": \"$data_folder\"}" > config.json

echo ""
echo "Installation complete!"
echo "You can now run mainframe by navigating to '$mainframe_folder' and running './mainframe'."