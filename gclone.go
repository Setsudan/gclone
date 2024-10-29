package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// Config holds the user configuration
type Config struct {
	DefaultUsername string `json:"default_username"`
	TmpDirectory    string `json:"tmp_directory"`
}

func getConfigPath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Error getting home directory: %v\n", err)
		os.Exit(1)
	}
	return filepath.Join(homeDir, ".gclone", "config.json")
}

func loadConfig() Config {
	configPath := getConfigPath()
	config := Config{
		DefaultUsername: "",
		TmpDirectory:    filepath.Join("C:", "Dev", "tmp"),
	}

	// Try to read existing config
	data, err := os.ReadFile(configPath)
	if err == nil {
		err = json.Unmarshal(data, &config)
		if err != nil {
			fmt.Printf("Error parsing config file: %v\n", err)
			os.Exit(1)
		}
	}

	return config
}

func saveConfig(config Config) error {
	configPath := getConfigPath()
	configDir := filepath.Dir(configPath)

	// Create config directory if it doesn't exist
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return fmt.Errorf("error creating config directory: %v", err)
	}

	// Save config
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling config: %v", err)
	}

	if err := os.WriteFile(configPath, data, 0644); err != nil {
		return fmt.Errorf("error writing config file: %v", err)
	}

	return nil
}

func setupConfig() {
	config := loadConfig()

	fmt.Print("Enter your default GitHub username: ")
	fmt.Scanln(&config.DefaultUsername)

	fmt.Printf("Enter your temporary directory path (press Enter for default '%s'): ", config.TmpDirectory)
	var tmpDir string
	fmt.Scanln(&tmpDir)
	if tmpDir != "" {
		config.TmpDirectory = tmpDir
	}

	if err := saveConfig(config); err != nil {
		fmt.Printf("Error saving config: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Configuration saved successfully!")
}

func main() {
	// Define flags
	openInVSCode := flag.Bool("c", false, "Open the cloned repository in VSCode")
	useTmpDir := flag.Bool("tmp", false, "Clone into temporary directory")
	configureFlag := flag.Bool("config", false, "Configure gclone settings")
	flag.Parse()

	// Handle configuration
	if *configureFlag {
		setupConfig()
		return
	}

	// Load configuration
	config := loadConfig()

	// Check if username is configured
	if config.DefaultUsername == "" {
		fmt.Println("gclone is not configured. Please run 'gclone -config' first.")
		os.Exit(1)
	}

	// Check if repository name is provided
	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Usage: gclone [-c] [-tmp] repository-name")
		fmt.Println("       gclone -config (to configure settings)")
		os.Exit(1)
	}

	repoName := args[0]

	// If only one word is provided (no slash), prepend default username
	if !strings.Contains(repoName, "/") {
		repoName = config.DefaultUsername + "/" + repoName
	}

	// Construct GitHub URL
	githubURL := fmt.Sprintf("https://github.com/%s", repoName)

	// Determine clone directory
	cloneDir := "."
	if *useTmpDir {
		cloneDir = config.TmpDirectory
		// Create tmp directory if it doesn't exist
		if err := os.MkdirAll(cloneDir, 0755); err != nil {
			fmt.Printf("Error creating tmp directory: %v\n", err)
			os.Exit(1)
		}
	}

	// Extract the final directory name from the repository
	finalDirName := filepath.Join(cloneDir, strings.Split(repoName, "/")[len(strings.Split(repoName, "/"))-1])

	// Execute git clone
	cmd := exec.Command("git", "clone", githubURL)
	cmd.Dir = cloneDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Printf("Cloning %s...\n", githubURL)
	if err := cmd.Run(); err != nil {
		fmt.Printf("Error cloning repository: %v\n", err)
		os.Exit(1)
	}

	// Open in VSCode if requested
	if *openInVSCode {
		codeCmd := exec.Command("code", finalDirName)
		if err := codeCmd.Run(); err != nil {
			fmt.Printf("Error opening VSCode: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Opened in VSCode")
	}

	fmt.Println("Clone completed successfully!")
}
