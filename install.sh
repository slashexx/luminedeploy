#!/bin/bash

installLumine() {
    echo "Installing Lumine..."

    INSTALL_DIR="/usr/local/bin"
    TEMP_DIR="/tmp/lumine"

    if ! command -v git &> /dev/null; then
        echo "Git is not installed. Please install Git and try again."
        exit 1
    fi

    if ! command -v go &> /dev/null; then
        echo "Go (Golang) is not installed. Please install Go and try again."
        exit 1
    fi

    echo "Cloning Lumine repository..."
    git clone https://github.com/nexusrex18/lumine.git "$TEMP_DIR" || {
        echo "Failed to clone the Lumine repository."
        exit 1
    }

    cd "$TEMP_DIR" || exit

    echo "Building Lumine CLI..."
    go build -o lumine main.go || {
        echo "Failed to build Lumine CLI."
        exit 1
    }

    echo "Moving Lumine binary to $INSTALL_DIR..."
    sudo mv lumine "$INSTALL_DIR/lumine" || {
        echo "Failed to move Lumine CLI to $INSTALL_DIR. Please check your permissions."
        exit 1
    }

    echo "Cleaning up temporary files..."
    rm -rf "$TEMP_DIR"

    if command -v lumine &> /dev/null; then
        echo "Lumine installed successfully! 🎉"
        echo "You can now run Lumine using the command: lumine"
    else
        echo "Something went wrong during the installation. Please check the steps above."
    fi
}

installLumine
