package main

import (
	"bufio"
	"fmt"
	"html/template"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const configFileName = ".gocreateconfig"

func getRepoPrefix() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	configPath := filepath.Join(homeDir, configFileName)

	if _, err := os.Stat(configPath); err == nil {
		content, err := os.ReadFile(configPath)
		if err != nil {
			return "", nil
		}

		repoPrefix := strings.TrimSpace(string(content))
		if repoPrefix != "" {
			return repoPrefix, nil
		}
	}

	fmt.Println("Enter your default Git repository (e.g., github.com/yourname):")
	reader := bufio.NewReader(os.Stdin)
	repoPrefix, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	repoPrefix = strings.TrimSpace(repoPrefix)
	if err := os.WriteFile(configPath, []byte(repoPrefix+"\n"), 0644); err != nil {
		return "", err
	}

	return repoPrefix, nil
}

func createProject(projectName, repoPrefix string) error {
	modulePath := fmt.Sprintf("%s/%s", repoPrefix, projectName)

	if err := runCommand("go", "mod", "init", modulePath); err != nil {
		return err
	}

	currentDir, err := os.Getwd()
	if err != nil {
		return err
	}

	mainGoPath := filepath.Join(currentDir, "main.go")
	mainContent := `package main

import "fmt"

func main() {
	fmt.Println("Hello, {{.}}!")
}
`
	file, err := os.Create(mainGoPath)
	if err != nil {
		return err
	}
	defer file.Close()

	tmpl, err := template.New("main").Parse(mainContent)
	if err != nil {
		return err
	}
	return tmpl.Execute(file, projectName)
}

func runCommand(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: gocreate <project-name>")
		os.Exit(1)
	}

	projectName := os.Args[1]
	repoPrefix, err := getRepoPrefix()
	if err != nil {
		fmt.Printf("Failed to get default repository: %v\n", err)
		os.Exit(1)
	}

	if err := createProject(projectName, repoPrefix); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Project '%s' successfully created using repository '%s'.\n", projectName, repoPrefix)
}
