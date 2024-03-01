package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

func main() {
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
