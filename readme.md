# gclone

A simple and convenient CLI tool to streamline the GitHub repository cloning process. `gclone` allows you to clone repositories with shorter commands and provides additional features like automatically opening projects in VSCode and managing a temporary project directory.

## Features

- 🚀 Quick clone with shortened commands
- 📂 Default username configuration
- 💻 Optional VSCode integration
- 📁 Temporary directory management
- ⚙️ User-configurable settings

## Installation

1. Download the latest release from the [releases page](https://github.com/setsudan/gclone/releases).

2. Add the executable to your system's PATH:
   - Windows: Move `gclone.exe` to a directory in your PATH (e.g., `C:\Users\YourUsername\bin\`)
   - Linux/MacOS: Move `gclone` to `/usr/local/bin/` or another directory in your PATH

3. Run the configuration:

```bash
gclone -config
```

## Usage

### Basic Usage

```bash
# Clone your own repository (uses configured username)
gclone myrepo

# Clone from specific user/organization
gclone kubernetes/kubernetes
```

### Advanced Features

```bash
# Clone and open in VSCode
gclone -c myrepo

# Clone to temporary directory
gclone -tmp myrepo

# Combine flags
gclone -c -tmp myrepo
```

### Configuration

```bash
# Run initial setup or update configuration
gclone -config
```

The configuration will prompt you for:

- Your default GitHub username
- Custom temporary directory path (optional)

Configuration is stored in `~/.gclone/config.json`:

```json
{
  "default_username": "yourusername",
  "tmp_directory": "C:\\Dev\\tmp"
}
```

## Requirements

- Git installed and accessible in PATH
- VSCode installed and `code` command available in PATH (only for `-c` flag)
- Write permissions for temporary directory creation

## Examples

```bash
# Clone your repository 'awesome-project'
gclone awesome-project
# Equivalent to: git clone https://github.com/yourusername/awesome-project

# Clone specific repository
gclone microsoft/vscode
# Equivalent to: git clone https://github.com/microsoft/vscode

# Clone and open in VSCode
gclone -c myproject

# Clone to temporary directory
gclone -tmp experimental-feature
```

## Command Line Arguments

| Flag      | Description                            |
|-----------|----------------------------------------|
| `-c`      | Open cloned repository in VSCode       |
| `-tmp`    | Clone into configured temporary directory |
| `-config` | Configure gclone settings              |

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## Building from Source

```bash
# Clone the repository
git clone https://github.com/setsudan/gclone

# Navigate to the directory
cd gclone

# Build the executable
go build -o gclone.exe
```

## License

MIT License - see the [LICENSE](LICENSE) file for details

## Acknowledgments

- Inspired by the need for a more streamlined git clone workflow
- Built with Go for cross-platform compatibility

## Support

If you encounter any issues or have suggestions, please file an issue on the [GitHub repository](https://github.com/setsudan/gclone/issues).
