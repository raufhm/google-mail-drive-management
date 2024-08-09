
# Google CLI

`googlecli` is a command-line tool for interacting with Gmail and Google Drive. It allows you to download and purge your emails and drive content based on various filters.

![20f71f07-4c70-4638-a79a-6a76c101088b](https://github.com/user-attachments/assets/b12e88ca-7b00-4010-beaa-c67d6a54edc8)


## Installation

To install `googlecli`, you can use the following command:

```bash
#!/bin/bash

# Build for Windows
GOOS=windows GOARCH=amd64 go build -o bin/googlecli-windows-amd64.exe

# Build for macOS (Intel)
GOOS=darwin GOARCH=amd64 go build -o bin/googlecli-darwin-amd64

# Build for macOS (Apple Silicon)
GOOS=darwin GOARCH=arm64 go build -o bin/googlecli-darwin-arm64

# Build for Linux (x86_64)
GOOS=linux GOARCH=amd64 go build -o bin/googlecli-linux-amd64

# Build for Linux (ARM64)
GOOS=linux GOARCH=arm64 go build -o bin/googlecli-linux-arm64

echo "Builds completed!"


Make sure to give the script execute permissions and run it:
chmod +x build.sh
./build.sh
```

## Usage

### General Usage

The `googlecli` CLI provides several commands and flags for interacting with Gmail and Google Drive. Below are the available commands and their respective flags.

### Commands

- **completion**: Generate the autocompletion script for the specified shell.
- **getGdriveContent**: Retrieve Google Drive content based on specified filters.
- **getGmailContent**: Retrieve Gmail content based on specified filters.
- **help**: Display help information for any command.

### Global Flags

- `-e, --email string`: Specify the email address for authentication.
- `-h, --help`: Display help information for the command.

### Example Usage

#### Get Gmail Content

To retrieve Gmail content, use the following command:

```bash
googlecli getGmailContent [flags]
```

Flags:
- `-d, --download`: Download the emails locally.
- `-e, --email string`: Specify the email address for authentication.
- `-p, --purge`: Purge the emails from Gmail after processing.

Example:

```bash
googlecli getGmailContent -e user@example.com -d
```

#### Get Google Drive Content

To retrieve Google Drive content, use the following command:

```bash
googlecli getGdriveContent [flags]
```

Flags:
- `-d, --download`: Download the drive content locally.
- `-e, --email string`: Specify the email address for authentication.
- `-p, --purge`: Purge the files from Google Drive after processing.

Example:

```bash
googlecli getGdriveContent -e user@example.com -d
```

### Autocompletion

To generate autocompletion scripts for your shell, use the `completion` command.

```bash
googlecli completion [bash|zsh|fish|powershell]
```

## Authentication

`googlecli` requires authentication with Google services. Make sure you have the necessary API credentials and permissions configured.

## Contributing

Contributions are welcome! Please read the [CONTRIBUTING.md](CONTRIBUTING.md) for details on the code of conduct and the process for submitting pull requests.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
