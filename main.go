package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

type Config struct {
	NotesDir string `json:"notes_dir"`
	Editor   string `json:"editor"`
}

func main() {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	// Load config or create default if it doesn't exist
	var config Config
	configPath := filepath.Join(home, ".nowt.json")
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		defaultConfig := Config{
			NotesDir: filepath.Join(home, "notes"),
			Editor:   "code",
		}
		jsonConfig, err := json.MarshalIndent(defaultConfig, "", "    ")
		if err != nil {
			log.Fatal(err)
		}
		err = os.WriteFile(configPath, jsonConfig, 0644)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		configFile, err := os.ReadFile(configPath)
		if err != nil {
			log.Fatal(err)
		}
		err = json.Unmarshal(configFile, &config)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Parse command line arguments
	name := os.Args[0]
	switch len(os.Args) {
	case 1:
		fmt.Println("nowt - organize your daily notes")
		fmt.Println("")
		fmt.Println("Usage:", name, "[OPTIONS] COMMAND")
		fmt.Println("")
		fmt.Println("Commands:")
		fmt.Println("  config    Configure nowt (not implemented yet)")
		fmt.Println("  write     Write a note")
		fmt.Println("  list      List notes (not implemented yet)")
		fmt.Println("")
		fmt.Println("Options:")
		fmt.Println("  -h, --help    Show this help message and exit")
	case 2:
		switch os.Args[1] {
		case "write":
			dirname, err := os.UserHomeDir()
			if err != nil {
				log.Fatal(err)
			}
			now := time.Now()
			dirPath := filepath.Join(
				dirname, "notes",
				now.Format("2006"),
				now.Format("01"),
				now.Format("02"),
			)
			err = os.MkdirAll(dirPath, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
			filePath := filepath.Join(dirPath, now.Format(time.DateOnly)+".md")
			os.Create(filePath)
			exec.Command("code", filePath).Run()
		}
	}
}
